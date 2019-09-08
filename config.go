package envconfig

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func ParseConfig(configFile string, unmarshalfunc func([]byte)) {
	if fileInfo, err := os.Stat(configFile); err != nil {
		if os.IsNotExist(err) {
			panic(fmt.Sprintf("configuration file [%s] does not exist.", configFile))
		} else {
			panic(fmt.Sprintf("configuration file [%s] can not be stated. %v", configFile, err))
		}
	} else {
		if fileInfo.IsDir() {
			panic(fmt.Sprintf("%s is a directory name", configFile))
		}
	}
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Sprintf("read configuration file error. %v", err))
	}
	content = bytes.TrimSpace(content)
	unmarshalfunc(content)
}
