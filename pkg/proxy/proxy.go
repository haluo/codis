package proxy

import (
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/wandoulabs/codis/pkg/models"
	"github.com/wandoulabs/codis/pkg/proxy/router"
	"github.com/wandoulabs/codis/pkg/utils/errors"
	"github.com/wandoulabs/codis/pkg/utils/log"
	"github.com/wandoulabs/codis/pkg/utils/rpc"
)

type Proxy struct {
	mu sync.Mutex

	token string

	router *router.Router
}

func New() *Proxy {
	hostname, _ := os.Hostname()
	token := rpc.NewToken(hostname)

	return &Proxy{
		token: token,

		router: router.New(),
	}
}

func (s *Proxy) GetSlots() []*models.SlotInfo {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.router.GetSlots()
}

func (s *Proxy) GetToken() string {
	return s.token
}

func (s *Proxy) LockSlot(i int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.router.LockSlot(i)
}

func (s *Proxy) FillSlot(i int, addr, from string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.router.FillSlot(i, addr, from)
}

func (s *Proxy) Shutdown() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.router.Close()
}

func (s *Proxy) Serve(l net.Listener) error {
	defer l.Close()

	h := http.NewServeMux()
	h.Handle("/", newApiServer(s))

	log.Infof("server[%p] now serving on %s", s, l.Addr().String())

	hs := &http.Server{Handler: h}
	return errors.Trace(hs.Serve(l))
}
