package base

import "os"

type UserInfo struct {
	Token string `yaml:"token"`
	Repo  string `yaml:"repo"`
}

var (
	userInfo UserInfo
)

func GetRepo() string {
	return userInfo.Repo
}

func GetToken() string {
	return userInfo.Token
}

func GetUserInfo() UserInfo {
	return userInfo
}

func initUserInfo() {
	userInfo = UserInfo{
		Repo:  os.Getenv("GITHUB_REPOSITORY"),
		Token: os.Getenv("GITHUB_TOKEN"),
	}
}
