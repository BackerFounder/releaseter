package config

import "os"

var (
	repo  string
	token string
)

func GetRepo() string {
	if repo == "" {
		repo = os.Getenv("GITHUB_REPOSITORY")
	}
	return repo
}

func GetToken() string {
	if token == "" {
		token = os.Getenv("GITHUB_TOKEN")
	}
	return token
}
