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

	init, exit struct {
		c chan struct{}
		sync.Once
	}
	closed bool

	config *Config
	router *router.Router
}

var (
	errClosedProxy = errors.New("use of closed proxy")
)

func New() *Proxy {
	return NewWithConfig(NewDefaultConfig())
}

func NewWithConfig(config *Config) *Proxy {
	s := &Proxy{
		token:  rpc.NewToken(),
		config: config,
		router: router.NewWithAuth(config.ProductAuth),
	}
	s.init.c = make(chan struct{})
	s.exit.c = make(chan struct{})

	go s.keepAlive()

	return s
}

func (s *Proxy) keepAlive() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for tick := 0; ; {
		select {
		case <-s.exit.c:
			return
		case <-ticker.C:
			if maxTick := s.config.BackendPingPeriod; maxTick != 0 {
				if tick++; tick >= maxTick {
					tick = 0
					s.router.KeepAlive()
				}
			}
		}
	}
}

func (s *Proxy) GetSlots() []*models.SlotInfo {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return nil
	}
	return s.router.GetSlots()
}

func (s *Proxy) GetToken() string {
	return s.token
}

func (s *Proxy) GetConfig() *Config {
	return s.config
}

func (s *Proxy) FillSlot(slots ...*models.SlotInfo) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return errClosedProxy
	}
	for _, slot := range slots {
		if err := s.router.FillSlot(slot.Id, slot.BackendAddr, slot.MigrateFrom, slot.Locked); err != nil {
			return err
		}
	}
	return nil
}

func (s *Proxy) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return errClosedProxy
	}
	s.init.Do(func() {
		close(s.init.c)
	})
	return nil
}

func (s *Proxy) Shutdown() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.exit.Do(func() {
		s.router.Close()
		s.closed = true
		close(s.exit.c)
	})
	return nil
}

func (s *Proxy) ServeHTTP(l net.Listener) error {
	defer l.Close()

	eh := make(chan error, 1)

	go func() {
		h := http.NewServeMux()
		h.Handle("/", newApiServer(s))
		hs := &http.Server{Handler: h}
		eh <- hs.Serve(l)
	}()

	select {
	case <-s.exit.c:
		return errClosedProxy
	case err := <-eh:
		return err
	}
}

func (s *Proxy) Serve(l net.Listener) error {
	defer l.Close()

	select {
	case <-s.exit.c:
		return errClosedProxy
	case <-s.init.c:
	}

	ch := make(chan net.Conn, 4096)
	go func() {
		for c := range ch {
			x := router.NewSessionSize(c, s.config.ProductAuth,
				s.config.SessionMaxBufSize, s.config.SessionMaxTimeout)

			x.SetKeepAlivePeriod(s.config.SessionKeepAlivePeriod)

			go x.Serve(s.router, s.config.SessionMaxPipeline)
		}
	}()

	eh := make(chan error, 1)
	go func() {
		defer close(ch)
		for {
			c, err := l.Accept()
			if err != nil {
				eh <- err
				return
			} else {
				ch <- c
			}
		}
	}()

	select {
	case <-s.exit.c:
		return errClosedProxy
	case err := <-eh:
		return err
	}
}
