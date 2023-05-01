package model

type GithubCommit struct {
	Message string `json:"message"`
}

type GithubGetCommit struct {
	Commit GithubCommit `json:"commit"`
	Sha    string       `json:"sha"`
}
