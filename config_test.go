package envconfig

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Config struct {
	Title   string
	Owner   ownerInfo
	DB      database `toml:"database"`
	Servers map[string]server
	Clients clients
}

type ownerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

type database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}

const path = `/Users/shikuanxu/go/src/github.com/sillyhatxu/environment-config`

func TestParseConfig(t *testing.T) {
	var conf Config
	ParseConfig(fmt.Sprintf("%s/%s", path, "config.conf"), func(content []byte) {
		err := toml.Unmarshal(content, &conf)
		if err != nil {
			panic(fmt.Sprintf("unmarshal toml object error. %v", err))
		}
	})
	assert.NotNil(t, conf)
	assert.EqualValues(t, conf.Title, "TOML Example")
	logrus.Infof("%#v", conf)
}
