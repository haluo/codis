package proxy

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"

	"github.com/wandoulabs/codis/pkg/models"
	"github.com/wandoulabs/codis/pkg/proxy/router"
	"github.com/wandoulabs/codis/pkg/utils"
	"github.com/wandoulabs/codis/pkg/utils/errors"
	"github.com/wandoulabs/codis/pkg/utils/log"
	"github.com/wandoulabs/codis/pkg/utils/rpc"
)

type Info struct {
	Version string `json:"version"`
	Compile string `json:"compile"`

	UnixTime int64  `json:"unixtime"`
	BootTime string `json:"boottime"`

	Pid    int     `json:"pid"`
	Pwd    string  `json:"pwd"`
	Config *Config `json:"config"`
	Online bool    `json:"online"`
	Closed bool    `json:"closed"`

	Token string `json:"token"`

	Stats *Stats             `json:"stats,omitempty"`
	Slots []*models.SlotInfo `json:"slots,omitempty"`
}

type Stats struct {
	Ops struct {
		Total int64             `json:"total"`
		Cmds  []*router.OpStats `json:"cmds,omitempty"`
	} `json:"ops"`

	Sessions struct {
		Total   int64 `json:"total"`
		Actived int64 `json:"actived"`
	} `json:"sessions"`
}

type apiServer struct {
	proxy *Proxy
}

func newApiServer(p *Proxy) http.Handler {
	m := martini.New()
	m.Use(martini.Recovery())
	m.Use(func(w http.ResponseWriter, req *http.Request, c martini.Context) {
		addr := req.Header.Get("X-Real-IP")
		if addr == "" {
			addr = req.Header.Get("X-Forwarded-For")
			if addr == "" {
				addr = req.RemoteAddr
			}
		}
		path := req.URL.Path
		if strings.HasPrefix(path, "/api") {
			log.Infof("[%p] API from %s call %s", p, addr, path)
		}
		c.Next()
	})

	api := &apiServer{p}

	r := martini.NewRouter()
	r.Get("/", api.Info)
	r.Get("/api/stats/:token/:xauth", api.Stats)
	r.Put("/api/online/:token/:xauth", api.Online)
	r.Put("/api/shutdown/:token/:xauth", api.Shutdown)
	r.Put("/api/fillslot/:token/:xauth", binding.Json([]*models.SlotInfo{}), api.FillSlot)

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
	xauth := params["xauth"]
	if xauth == "" {
		return errors.New("Missing XAuth")
	}
	if xauth != s.proxy.GetXAuth() {
		return errors.New("Unmatched XAuth")
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
	info.Config = s.proxy.GetConfig()
	info.Online = s.proxy.IsOnline()
	info.Closed = s.proxy.IsClosed()

	info.Token = s.proxy.GetToken()
	info.Slots = s.proxy.GetSlots()
	info.Stats = s.GetStats()
	return rpc.ApiResponseJson(info)
}

func (s *apiServer) GetStats() *Stats {
	stats := &Stats{}
	stats.Ops.Total = router.OpsTotal()
	stats.Ops.Cmds = router.GetAllOpStats()
	stats.Sessions.Total = router.SessionsTotal()
	stats.Sessions.Actived = router.SessionsActived()
	return stats
}

func (s *apiServer) Stats(params martini.Params) (int, string) {
	if err := s.verifyToken(params); err != nil {
		return rpc.ApiResponseError(err)
	} else {
		return rpc.ApiResponseJson(s.GetStats())
	}
}

func (s *apiServer) Online(params martini.Params) (int, string) {
	if err := s.verifyToken(params); err != nil {
		return rpc.ApiResponseError(err)
	}
	if err := s.proxy.Start(); err != nil {
		return rpc.ApiResponseError(err)
	} else {
		return rpc.ApiResponseJson("OK")
	}
}

func (s *apiServer) Shutdown(params martini.Params) (int, string) {
	if err := s.verifyToken(params); err != nil {
		return rpc.ApiResponseError(err)
	}
	if err := s.proxy.Close(); err != nil {
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

type ApiClient struct {
	host  string
	token string
	xauth string
}

func NewApiClient(host string) *ApiClient {
	return &ApiClient{host: host}
}

func (c *ApiClient) SetToken(token string, auth string) {
	c.token = token
	c.xauth = rpc.EncryptAuth(auth, token)
}

func (c *ApiClient) encodeURL(format string, args ...interface{}) string {
	return rpc.EncodeURL(c.host, format, args...)
}

func (c *ApiClient) GetInfo() (*Info, error) {
	url := c.encodeURL("/")
	info := &Info{}
	if err := rpc.ApiGetJson(url, info); err != nil {
		return nil, err
	}
	return info, nil
}

func (c *ApiClient) GetStats() (*Stats, error) {
	url := c.encodeURL("/api/stats/%s/%s", c.token, c.xauth)
	stats := &Stats{}
	if err := rpc.ApiGetJson(url, stats); err != nil {
		return nil, err
	}
	return stats, nil
}

func (c *ApiClient) Online() error {
	url := c.encodeURL("/api/online/%s/%s", c.token, c.xauth)
	return rpc.ApiPutJson(url, nil, nil)
}

func (c *ApiClient) Shutdown() error {
	url := c.encodeURL("/api/shutdown/%s/%s", c.token, c.xauth)
	return rpc.ApiPutJson(url, nil, nil)
}

func (c *ApiClient) FillSlot(slots ...*models.SlotInfo) error {
	url := c.encodeURL("/api/fillslot/%s/%s", c.token, c.xauth)
	return rpc.ApiPutJson(url, slots, nil)
}
