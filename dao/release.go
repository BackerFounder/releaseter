package dao

import (
	"test/api"
	"test/model"
)

var (
	releases      []*model.GithubGetRelease
	latestRelease *model.Release
)

func GetRelease() []*model.GithubGetRelease {
	userInfo := GetUserInfo()
	if releases == nil {
		releases = api.GetRelease(userInfo.Repo, userInfo.Token)
	}
	return releases
}

func GetLatestRelease() *model.Release {
	userInfo := GetUserInfo()
	if latestRelease == nil {
		originRelease := api.GetLatestRelease(userInfo.Repo, userInfo.Token)
		originTags := api.GetTags(userInfo.Repo, userInfo.Token)

		for _, tag := range originTags {
			if tag.Name == originRelease.TagName {
				latestRelease = &model.Release{
					Sha:       tag.Sha,
					TagName:   originRelease.TagName,
					CreatedAt: originRelease.CreatedAt,
				}
				return latestRelease
			}
		}
		panic("找不到 tags")
	}

	return latestRelease
}

func GetAllReleaseDraftIds() []uint64 {
	var ids = make([]uint64, 0, 5)
	for _, release := range GetRelease() {
		if release.Draft {
			ids = append(ids, release.Id)
		}
	}
	return ids
}
