package base

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ConfigCategorie struct {
	Title  string   `yaml:"title"`
	Labels []string `yaml:"labels"`
	Label  string   `yaml:"label"`
}

type CategoryOther struct {
	Show  bool   `yaml:"show"`
	Title string `yaml:"title"`
}

type ConfigExceptRelease struct {
	Tag  string `yaml:"tag"`
	Name string `yaml:"name"`
}

type Config struct {
	NameTemplate string `yaml:"name-template"`

	TagTemplate   string `yaml:"tag-template"`
	TagPreRelease string `yaml:"tag-preRelease"`
	TagBuild      string `yaml:"tag-build"`

	Categories    []*ConfigCategorie `yaml:"categories"`
	CategoryOther *CategoryOther     `yaml:"category-other"`
	TimeFormat    string             `yaml:"time-format"`
	TimeLocation  string             `yaml:"time-location"`

	ClearHistoryDraft bool `yaml:"clear-history-draft"`

	ExceptReleases []*ConfigExceptRelease `yaml:"except_releases"`
}

var (
	userConfig    Config
	defaultConfig Config
)

func initDefaultConfig() {
	filePath := filepath.Join(os.Getenv("JZGOOPI_REPEASETER_WS_PATH"), "base/default.yml")
	cnf, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	defaultConfig = Config{}
	err = yaml.Unmarshal(cnf, &defaultConfig)
	if err != nil {
		panic(err)
	}
}

func GetDefaultConfig() Config {
	return defaultConfig
}

func initConfig() {
	filePath := filepath.Join(os.Getenv("GITHUB_WORKSPACE"), ".github", os.Getenv("JZGOOPI_REPEASETER_CONFIG_PATH"))
	cnf, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("您沒有提供配置文件，即將使用默認配置")
		userConfig = defaultConfig
		return
	}
	userConfig = Config{}
	err = yaml.Unmarshal(cnf, &userConfig)
	if err != nil {
		panic(err)
	}

	if userConfig.TimeLocation == "" {
		userConfig.TimeLocation = defaultConfig.TimeLocation
	}

	if userConfig.TimeFormat == "" {
		userConfig.TimeFormat = defaultConfig.TimeFormat
	}

	if userConfig.NameTemplate == "" {
		userConfig.NameTemplate = defaultConfig.NameTemplate
	}

	if userConfig.TagTemplate == "" {
		userConfig.TagTemplate = defaultConfig.TagTemplate
	}

	if userConfig.CategoryOther == nil {
		userConfig.CategoryOther = &CategoryOther{
			Show:  defaultConfig.CategoryOther.Show,
			Title: defaultConfig.CategoryOther.Title,
		}
	}
}

func GetConfig() Config {
	return userConfig
}

func Init() {
	initDefaultConfig()
	initConfig()
	initUserInfo()
}
