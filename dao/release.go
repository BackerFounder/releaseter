package dao

import (
	"github.com/JZGoopi/releaseter/base"

	"github.com/JZGoopi/releaseter/api"
	"github.com/JZGoopi/releaseter/model"
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
		latestRelease = &model.Release{
			Sha:       "",
			TagName:   originRelease.TagName,
			CreatedAt: originRelease.CreatedAt,
		}

		if originRelease.CreatedAt.IsZero() {
			return latestRelease
		}

		for _, tag := range tags {
			if tag.Name == originRelease.TagName {
				latestRelease.Sha = tag.Sha
				return latestRelease
			}
		}
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
