package api

import (
	"encoding/json"
	"fmt"

	"github.com/JZGoopi/releaseter/model"
)

func GetTags(repo, token string) []*model.Tag {
	url := fmt.Sprintf("https://api.github.com/repos/%s/tags", repo)
	headers := genBaseHeaders(token)

	data := getData(url, headers, nil)

	var apiTags = make([]model.GithubGetTagWithCommit, 0, 100)
	err := json.Unmarshal(data, &apiTags)
	if err != nil {
		panic(err)
	}

	var tags = make([]*model.Tag, 0, len(apiTags))
	for _, tag := range apiTags {
		tags = append(tags, &model.Tag{
			Name: tag.Name,
			Sha:  tag.Commit.Sha,
		})
	}

	return tags
}

func GetTagByTagName(repo, token, tagName string) *model.Tag {
	url := fmt.Sprintf("https://api.github.com/repos/%s/git/ref/tags/%s", repo, tagName)
	headers := genBaseHeaders(token)
	data := getData(url, headers, nil)

	var apiTag = model.GithubGetTagWithObject{}
	err := json.Unmarshal(data, &apiTag)
	if err != nil {
		panic(err)
	}
	return &model.Tag{
		Name: tagName,
		Sha:  apiTag.Object.Sha,
	}
}
