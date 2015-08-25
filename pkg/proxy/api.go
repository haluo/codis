package proxy

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-martini/martini"

	"github.com/wandoulabs/codis/pkg/models"
	"github.com/wandoulabs/codis/pkg/proxy/router"
	"github.com/wandoulabs/codis/pkg/utils"
	"github.com/wandoulabs/codis/pkg/utils/errors"
	"github.com/wandoulabs/codis/pkg/utils/rpc"
)

type Stats struct {
	Version string `json:"version"`
	Compile string `json:"compile"`

	UnixTime int64 `json:"unixtime"`

	Token string `json:"token"`

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
	r.Get("/", api.GetStats)
	r.Get("/slots", api.GetSlots)

	r.Put("/api/:token/ping", api.Ping)
	r.Put("/api/:token/lockslot/:slotid", api.LockSlot)
	r.Put("/api/:token/fillslot/:slotid/:addr64/:from64", api.FillSlot)
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

func (s *apiServer) paramSlotId(params martini.Params) (int, error) {
	t := params["slotid"]
	if t == "" {
		return 0, errors.New("Missing SlotId")
	}
	v, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return 0, errors.New("Invalid SlotId: " + err.Error())
	}
	return int(v), nil
}

func (s *apiServer) GetStats() (int, string) {
	stats := &Stats{
		Version:  utils.Version,
		Compile:  utils.Compile,
		UnixTime: time.Now().Unix(),
	}
	stats.Token = s.proxy.GetToken()
	stats.Ops.Total = router.OpsTotal()
	stats.Ops.Cmds = router.GetAllOpStats()
	return rpc.ApiResponseJson(stats)
}

func (s *apiServer) GetSlots() (int, string) {
	return rpc.ApiResponseJson(s.proxy.GetSlots())
}

func (s *apiServer) Ping(params martini.Params) (int, string) {
	if err := s.verifyToken(params); err != nil {
		return rpc.ApiResponseError(err)
	} else {
		return rpc.ApiResponseJson("OK")
	}
}

func (s *apiServer) LockSlot(params martini.Params) (int, string) {
	if err := s.verifyToken(params); err != nil {
		return rpc.ApiResponseError(err)
	}
	i, err := s.paramSlotId(params)
	if err != nil {
		return rpc.ApiResponseError(err)
	}
	if err := s.proxy.LockSlot(i); err != nil {
		return rpc.ApiResponseError(err)
	} else {
		return rpc.ApiResponseJson("OK")
	}
}

func (s *apiServer) FillSlot(params martini.Params) (int, string) {
	if err := s.verifyToken(params); err != nil {
		return rpc.ApiResponseError(err)
	}
	i, err := s.paramSlotId(params)
	if err != nil {
		return rpc.ApiResponseError(err)
	}
	addr, err := rpc.Decode64(params["addr64"])
	if err != nil {
		return rpc.ApiResponseError(err)
	}
	from, err := rpc.Decode64(params["from64"])
	if err != nil {
		return rpc.ApiResponseError(err)
	}
	if err := s.proxy.FillSlot(i, addr, from); err != nil {
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

func (c *ApiClient) newGetRequest(reply interface{}, format string, args ...interface{}) error {
	return rpc.ApiGetAsJson(rpc.EncodeURL(c.Host, format, args...), reply)
}

func (c *ApiClient) newPutRequest(reply interface{}, format string, args ...interface{}) error {
	return rpc.ApiPutAsJson(rpc.EncodeURL(c.Host, format, args...), reply)
}

func (c *ApiClient) GetStats() (*Stats, error) {
	stats := &Stats{}
	if err := c.newGetRequest(stats, "/"); err != nil {
		return nil, err
	}
	return stats, nil
}

func (c *ApiClient) GetSlots() ([]*models.SlotInfo, error) {
	slots := []*models.SlotInfo{}
	if err := c.newGetRequest(&slots, "/slots"); err != nil {
		return nil, err
	}
	return slots, nil
}

func (c *ApiClient) Ping(token string) error {
	return c.newPutRequest(nil, "/api/%s/ping", token)
}

func (c *ApiClient) LockSlot(token string, i int) error {
	return c.newPutRequest(nil, "/api/%s/lockslot/%d", token, i)
}

func (c *ApiClient) FillSlot(token string, i int, addr, from string) error {
	addr64 := rpc.Encode64(addr)
	from64 := rpc.Encode64(from)
	return c.newPutRequest(nil, "/api/%s/fillslot/%d/%s/%s", token, i, addr64, from64)
}

func (c *ApiClient) Shutdown(token string) error {
	return c.newPutRequest(nil, "/api/%s/shutdown", token)
}
