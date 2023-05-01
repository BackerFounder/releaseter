package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	config []byte
)

func GetConfig() []byte {
	if config == nil {
		filePath := filepath.Join(os.Getenv("GITHUB_WORKSPACE"), ".github", os.Getenv("JZGOOPI_REPEASETER_CONFIG_PATH"))
		cnf, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("您沒有提供配置文件，即將使用默認配置")
			filePath = filepath.Join(os.Getenv("JZGOOPI_REPEASETER_WS_PATH"), "config/default.yml")
			cnf, err = ioutil.ReadFile(filePath)
			if err != nil {
				panic(err)
			}
		}

		config = cnf
	}

	return config
}
