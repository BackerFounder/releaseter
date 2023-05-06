package dao

import (
	"test/api"
	"test/base"
	"test/model"
)

var (
	releases      []*model.GithubGetRelease
	latestRelease *model.Release
)

func GetRelease(userInfo base.UserInfo) []*model.GithubGetRelease {
	if releases == nil {
		releases = api.GetRelease(userInfo.Repo, userInfo.Token)
	}
	return releases
}

func GetLatestRelease(userInfo base.UserInfo, tags []*model.Tag) *model.Release {

	if latestRelease == nil {
		originRelease := api.GetLatestRelease(userInfo.Repo, userInfo.Token)

		for _, tag := range tags {
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

func GetAllReleaseDraftIds(userInfo base.UserInfo) []uint64 {
	var ids = make([]uint64, 0, 5)
	for _, release := range GetRelease(userInfo) {
		if release.Draft {
			ids = append(ids, release.Id)
		}
	}
	return ids
}
