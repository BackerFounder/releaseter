package base

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
		Repo:  "JZGoopi/heybaybay-release-draft",          // os.Getenv("GITHUB_REPOSITORY")
		Token: "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", // os.Getenv("GITHUB_TOKEN")
	}
}
