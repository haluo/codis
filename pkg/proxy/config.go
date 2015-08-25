package proxy

import (
	"bytes"

	"github.com/BurntSushi/toml"

	"github.com/wandoulabs/codis/pkg/utils/errors"
)

type Config struct {
	ServAddr string `toml:"serv_addr"`
	HttpAddr string `toml:"http_addr"`

	ProductName   string `toml:"product_name"`
	ProductAuth   string `toml:"product_auth"`
	DashboardAddr string `toml:"dashboard_addr"`

	BackendPingPeriod      int  `toml:"backend_ping_period"`
	SessionMaxTimeout      int  `toml:"session_max_timeout"`
	SessionMaxBufSize      int  `toml:"session_max_bufsize"`
	SessionMaxPipeline     int  `toml:"session_max_pipeline"`
	SessionKeepAlive       bool `toml:"session_keepalive"`
	SessionKeepAlivePeriod int  `toml:"session_keepalive_period"`
}

func NewDefaultConfig() *Config {
	return &Config{
		ServAddr: "0.0.0.0:19000",
		HttpAddr: "0.0.0.0:17950",

		ProductName:   "Demo2.0",
		ProductAuth:   "",
		DashboardAddr: "0.0.0.0:18950",

		BackendPingPeriod:  5,
		SessionMaxTimeout:  60 * 30,
		SessionMaxBufSize:  1024 * 128,
		SessionMaxPipeline: 1024,

		SessionKeepAlive:       true,
		SessionKeepAlivePeriod: 60 * 15,
	}
}

func (c *Config) LoadFromFile(path string) error {
	_, err := toml.DecodeFile(path, c)
	return errors.Trace(err)
}

func (c *Config) String() string {
	var b bytes.Buffer
	e := toml.NewEncoder(&b)
	e.Indent = "    "
	e.Encode(c)
	return b.String()
}
