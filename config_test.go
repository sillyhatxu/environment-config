package envconfig

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"reflect"
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

//string_value=test src;int_value=111;int32_value=32;int64_value=64;bool_value=true;float32_value=3.14;float64_value=3.14159265372
func TestParseEnvironmentConfig(t *testing.T) {
	type EnvironmentConfig struct {
		StringValue  string `env:"string_value"`
		NoTag        string
		IntValue     int     `env:"int_value"`
		Int32Value   int32   `env:"int32_value"`
		Int64Value   int64   `env:"int64_value"`
		BoolValue    bool    `env:"bool_value"`
		Float32Value float32 `env:"float32_value"`
		Float64Value float64 `env:"float64_value"`
	}
	var config EnvironmentConfig
	err := ParseEnvironmentConfig(&config)
	logrus.Infof("env : %#v",config)
	assert.Nil(t, err)
}

func TestParseEnvironmentConfig1(t *testing.T) {
	type MyStruct struct {
		N int
	}
	n := MyStruct{1}
	// get
	immutable := reflect.ValueOf(n)
	val := immutable.FieldByName("N").Int()
	fmt.Printf("N=%d\n", val) // prints 1
	// set
	mutable := reflect.ValueOf(&n).Elem()
	mutable.FieldByName("N").SetInt(7)
	fmt.Printf("N=%d\n", n.N) // prints 7
}
