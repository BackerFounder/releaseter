package dao

import (
	"bytes"
	"test/config"
	"test/consts"
	"test/model"
	"time"

	"gopkg.in/yaml.v3"
)

var (
	originConfigTxt []byte
	originConfig    *model.Config
	parsedConfig    *model.Config
)

func getOriginConfigTxt() []byte {
	if originConfigTxt == nil {
		originConfigTxt = config.GetConfig()
	}
	return append(make([]byte, 0, len(originConfigTxt)), originConfigTxt...)
}

func GetOriginConfig() *model.Config {
	if originConfig == nil {
		originCnfData := getOriginConfigTxt()
		originConfig = &model.Config{}
		err := yaml.Unmarshal(originCnfData, originConfig)
		if err != nil {
			panic(err)
		}
	}
	return originConfig
}

func replaceConfigData(data []byte) []byte {

	// 替換版本號
	version := GetVersion()
	for _, placeholder := range consts.GetVersionPlaceholders() {
		data = bytes.Replace(data, []byte(placeholder), []byte(version.Join(placeholder, "", "")), -1)
	}

	// 替換時間
	loc, err := time.LoadLocation(originConfig.TimeLocation)
	if err != nil {
		loc, _ = time.LoadLocation("Asia/Taipei")
	}
	for placeholder, time := range consts.GetTimePlaceholders() {
		data = bytes.Replace(data, []byte(placeholder), []byte(time.In(loc).Format(originConfig.TimeFormat)), -1)
	}
	return data
}

func GetConfig() *model.Config {
	if parsedConfig == nil {
		originCnfData := getOriginConfigTxt()
		originCnfData = replaceConfigData(originCnfData)
		parsedConfig = &model.Config{}
		err := yaml.Unmarshal(originCnfData, parsedConfig)
		if err != nil {
			panic(err)
		}
	}
	return parsedConfig
}
