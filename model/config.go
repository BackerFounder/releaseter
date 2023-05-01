package model

type ConfigCategorie struct {
	Title  string   `yaml:"title"`
	Labels []string `yaml:"labels"`
	Label  string   `yaml:"label"`
}

type Config struct {
	NameTemplate string             `yaml:"name-template"`
	TagTemplate  string             `yaml:"tag-template"`
	Categories   []*ConfigCategorie `yaml:"categories"`
	TimeFormat   string             `yaml:"time-format"`
	TimeLocation string             `yaml:"time-location"`
}

type UserInfo struct {
	Token string `yaml:"token"`
	Repo  string `yaml:"repo"`
}
