package api

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/JZGoopi/releaseter/model"
)

func GetRelease(repo, token string) []*model.GithubGetRelease {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases", repo)
	headers := genBaseHeaders(token)
	query := Querys{}
	query["per_page"] = "100"

	data := getData(url, headers, query)
	var release []*model.GithubGetRelease
	json.Unmarshal(data, &release)
	return release
}

func GetLatestRelease(repo, token string) *model.GithubGetRelease {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)
	headers := genBaseHeaders(token)
	data := getData(url, headers, nil)
	var release model.GithubGetRelease
	json.Unmarshal(data, &release)
	return &release
}

func PostReleases(repo, token string, releaseData model.GithubPostRelease) {

	url := fmt.Sprintf("https://api.github.com/repos/%s/releases", repo)
	headers := genBaseHeaders(token)
	jsonStr, err := json.Marshal(releaseData)
	if err != nil {
		panic(err)
	}
	postData(url, headers, nil, bytes.NewBuffer(jsonStr))
}

func DelReleases(repo, token string, release_ids []uint64) {
	headers := genBaseHeaders(token)

	for _, id := range release_ids {
		url := fmt.Sprintf("https://api.github.com/repos/%s/releases/%d", repo, id)
		delData(url, headers, nil)
	}
}
