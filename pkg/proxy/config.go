package proxy

import (
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"

	"github.com/wandoulabs/codis/pkg/utils/errors"
)

type Config struct {
	BindType string `toml:"bind_type" json:"bind_type"`
	BindAddr string `toml:"bind_addr" json:"bind_addr"`
	HttpAddr string `toml:"http_addr" json:"http_addr"`

	ProductName string `toml:"product_name" json:"product_name"`
	ProductAuth string `toml:"product_auth" json:"-"`

	BackendPingPeriod      int `toml:"backend_ping_period" json:"backend_ping_period"`
	SessionMaxTimeout      int `toml:"session_max_timeout" json:"session_max_timeout"`
	SessionMaxBufSize      int `toml:"session_max_bufsize" json:"session_max_bufsize"`
	SessionMaxPipeline     int `toml:"session_max_pipeline" json:"session_max_pipeline"`
	SessionKeepAlivePeriod int `toml:"session_keepalive_period" json:"session_keepalive_period"`
}

func NewDefaultConfig() *Config {
	return &Config{
		BindType: "tcp",
		BindAddr: "0.0.0.0:19000",
		HttpAddr: "0.0.0.0:17950",

		ProductName: "Demo2",
		ProductAuth: "",

		BackendPingPeriod:  5,
		SessionMaxTimeout:  60 * 30,
		SessionMaxBufSize:  1024 * 128,
		SessionMaxPipeline: 1024,

		SessionKeepAlivePeriod: 60,
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

func (c *Config) JsonString() string {
	b, _ := json.MarshalIndent(c, "", "    ")
	return string(b)
}
