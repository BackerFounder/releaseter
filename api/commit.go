package api

import (
	"encoding/json"
	"fmt"
	"test/model"
)

func GetCommits(repo, token string) []*model.GithubGetCommit {
	headers := genBaseHeaders(token)
	url := fmt.Sprintf("https://api.github.com/repos/%s/commits", repo)
	data := getData(url, headers, nil)
	commitDatas := []*model.GithubGetCommit{}
	err := json.Unmarshal(data, &commitDatas)
	if err != nil {
		panic(err)
	}

	return commitDatas
}
