package proxy

import (
	"net/http"
	"os"
	"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"

	"github.com/wandoulabs/codis/pkg/models"
	"github.com/wandoulabs/codis/pkg/proxy/router"
	"github.com/wandoulabs/codis/pkg/utils"
	"github.com/wandoulabs/codis/pkg/utils/errors"
	"github.com/wandoulabs/codis/pkg/utils/rpc"
)

type Info struct {
	Version string `json:"version"`
	Compile string `json:"compile"`

	UnixTime int64  `json:"unixtime"`
	BootTime string `json:"boottime"`

	Pid      int     `json:"pid"`
	Pwd      string  `json:"pwd"`
	Hostname string  `json:"hostname"`
	Config   *Config `json:"config"`

	Token string `json:"token"`

	Slots []*models.SlotInfo `json:"slots"`
	Stats Stats              `json:"stats"`
}

type Stats struct {
	Ops struct {
		Total int64             `json:"total"`
		Cmds  []*router.OpStats `json:"cmds,omitempty"`
	} `json:"ops"`
}

type apiServer struct {
	proxy *Proxy
}

func newApiServer(p *Proxy) http.Handler {
	m := martini.New()
	m.Use(martini.Recovery())

	api := &apiServer{p}

	r := martini.NewRouter()
	r.Get("/", api.Info)
	r.Get("/api/:token/stats", api.Stats)
	r.Put("/api/:token/start", api.Start)
	r.Put("/api/:token/fillslot", binding.Json([]*models.SlotInfo{}), api.FillSlot)
	r.Put("/api/:token/shutdown", api.Shutdown)

	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)
	return m
}

func (s *apiServer) verifyToken(params martini.Params) error {
	token := params["token"]
	if token == "" {
		return errors.New("Missing Token")
	}
	if token != s.proxy.GetToken() {
		return errors.New("Unmatched Token")
	}
	return nil
}

func (s *apiServer) Info() (int, string) {
	info := &Info{
		Version:  utils.Version,
		Compile:  utils.Compile,
		UnixTime: time.Now().Unix(),
		BootTime: utils.BootTime.String(),
	}
	info.Pid = os.Getpid()
	info.Pwd, _ = os.Getwd()
	info.Hostname, _ = os.Hostname()
	info.Config = s.proxy.GetConfig()

	info.Token = s.proxy.GetToken()
	info.Slots = s.proxy.GetSlots()
	info.Stats.Ops.Total = router.OpsTotal()
	info.Stats.Ops.Cmds = router.GetAllOpStats()
	return rpc.ApiResponseJson(info)
}

func (s *apiServer) Stats(params martini.Params) (int, string) {
	if err := s.verifyToken(params); err != nil {
		return rpc.ApiResponseError(err)
	} else {
		stats := &Stats{}
		stats.Ops.Total = router.OpsTotal()
		stats.Ops.Cmds = router.GetAllOpStats()
		return rpc.ApiResponseJson(stats)
	}
}

func (s *apiServer) Start(params martini.Params) (int, string) {
	if err := s.verifyToken(params); err != nil {
		return rpc.ApiResponseError(err)
	}
	if err := s.proxy.Start(); err != nil {
		return rpc.ApiResponseError(err)
	} else {
		return rpc.ApiResponseJson("OK")
	}
}

func (s *apiServer) FillSlot(slots []*models.SlotInfo, params martini.Params) (int, string) {
	if err := s.verifyToken(params); err != nil {
		return rpc.ApiResponseError(err)
	}
	if err := s.proxy.FillSlot(slots...); err != nil {
		return rpc.ApiResponseError(err)
	} else {
		return rpc.ApiResponseJson("OK")
	}
}

func (s *apiServer) Shutdown(params martini.Params) (int, string) {
	if err := s.verifyToken(params); err != nil {
		return rpc.ApiResponseError(err)
	}
	if err := s.proxy.Shutdown(); err != nil {
		return rpc.ApiResponseError(err)
	} else {
		return rpc.ApiResponseJson("OK")
	}
}

type ApiClient struct {
	Host string
}

func NewApiClient(host string) *ApiClient {
	return &ApiClient{host}
}

func (c *ApiClient) encodeURL(format string, args ...interface{}) string {
	return rpc.EncodeURL(c.Host, format, args...)
}

func (c *ApiClient) GetInfo() (*Info, error) {
	url := c.encodeURL("/")
	info := &Info{}
	if err := rpc.ApiGetJson(url, info); err != nil {
		return nil, err
	}
	return info, nil
}

func (c *ApiClient) GetStats(token string) (*Stats, error) {
	url := c.encodeURL("/api/%s/stats", token)
	stats := &Stats{}
	if err := rpc.ApiGetJson(url, stats); err != nil {
		return nil, err
	}
	return stats, nil
}

func (c *ApiClient) Start(token string) error {
	url := c.encodeURL("/api/%s/start", token)
	return rpc.ApiPutJson(url, nil, nil)
}

func (c *ApiClient) FillSlot(token string, slots ...*models.SlotInfo) error {
	url := c.encodeURL("/api/%s/fillslot", token)
	return rpc.ApiPutJson(url, slots, nil)
}

func (c *ApiClient) Shutdown(token string) error {
	url := c.encodeURL("/api/%s/shutdown", token)
	return rpc.ApiPutJson(url, nil, nil)
}
