package proxy

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/wandoulabs/codis/pkg/models"
	"github.com/wandoulabs/codis/pkg/proxy/router"
	"github.com/wandoulabs/codis/pkg/utils/errors"
	"github.com/wandoulabs/codis/pkg/utils/rpc"
)

type Proxy struct {
	mu sync.Mutex

	token string

	signal chan interface{}
	closed bool

	config *Config
	router *router.Router
}

func New() *Proxy {
	return NewWithConfig(NewDefaultConfig())
}

func NewWithConfig(config *Config) *Proxy {
	s := &Proxy{
		token: rpc.NewToken(),

		signal: make(chan interface{}),

		config: config,
		router: router.NewWithAuth(config.ProductAuth),
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		var tick = 0
		for {
			select {
			case <-ticker.C:
			case <-s.signal:
				return
			}
			if maxTick := config.BackendPingPeriod; maxTick != 0 {
				if tick++; tick >= maxTick {
					tick = 0
					s.router.KeepAlive()
				}
			}
		}
	}()
	return s
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
	if s.closed {
		return nil
	}
	s.closed = true
	close(s.signal)
	s.router.Close()
	return nil
}

func (s *Proxy) ServeHTTP(l net.Listener) error {
	sig := make(chan interface{})
	defer close(sig)

	go func() {
		select {
		case <-sig:
		case <-s.signal:
		}
		l.Close()
	}()

	h := http.NewServeMux()
	h.Handle("/", newApiServer(s))

	hs := &http.Server{Handler: h}
	return errors.Trace(hs.Serve(l))
}

func (s *Proxy) Serve(l net.Listener) error {
	sig := make(chan interface{})
	defer close(sig)

	go func() {
		select {
		case <-sig:
		case <-s.signal:
		}
		l.Close()
	}()

	ch := make(chan net.Conn, 4096)
	defer close(ch)

	go func() {
		for c := range ch {
			x := router.NewSessionSize(c, s.config.ProductAuth,
				s.config.SessionMaxBufSize, s.config.SessionMaxTimeout)

			x.SetKeepAlivePeriod(s.config.SessionKeepAlivePeriod)

			go x.Serve(s.router, s.config.SessionMaxPipeline)
		}
	}()

	for {
		c, err := l.Accept()
		if err != nil {
			return nil
		} else {
			ch <- c
		}
	}
}
