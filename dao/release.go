package dao

import (
	"strings"

	"github.com/JZGoopi/releaseter/base"

	"github.com/JZGoopi/releaseter/api"
	"github.com/JZGoopi/releaseter/model"
)

var (
	releases      []*model.GithubGetRelease
	latestRelease *model.Release
)

func GetReleases(userInfo base.UserInfo) []*model.GithubGetRelease {
	if releases == nil {
		releases = api.GetReleases(userInfo.Repo, userInfo.Token)
	}
	return releases
}

func GetLatestRelease(userInfo base.UserInfo, tags []*model.Tag, cfg base.Config) *model.Release {

	if latestRelease == nil {

		// 1. 找到最新的 Release
		// 1.1 獲取所有 Release
		releases := GetReleases(userInfo)
		// 1.2 先初始化一個 Release
		githubRelease := &model.GithubGetRelease{}

		// 1.3 找出最新的 Release
	find_latest_releaser:
		for _, release := range releases {
			// 1.3.1 如果是草稿，就剔除
			if release.Draft {
				continue
			}

			// 1.3.2 如果沒有需要過濾，就取第一個
			if len(cfg.ExceptReleases) == 0 && cfg.ExceptKeyword == "" {
				githubRelease = release
				break
			}

			// 1.3.3 找到是否符合過濾規則的
			for _, configExceptRelease := range cfg.ExceptReleases {
				if release.TagName == configExceptRelease.Tag {
					if configExceptRelease.Name == "" || configExceptRelease.Name == release.Name {
						continue find_latest_releaser
					}
				}
			}

			if cfg.ExceptKeyword != "" && strings.Contains(release.Name, cfg.ExceptKeyword) {
				continue find_latest_releaser
			}

			// 1.3.4 如果不符合就表示找到了
			githubRelease = release
			break
		}

		// 2. 創建一個業務邏輯需要用的 Release
		latestRelease = &model.Release{
			Sha:       "",
			TagName:   githubRelease.TagName,
			CreatedAt: githubRelease.CreatedAt,
		}

		if githubRelease.CreatedAt.IsZero() {
			return latestRelease
		}

		for _, tag := range tags {
			if tag.Name == githubRelease.TagName {
				latestRelease.Sha = tag.Sha
				return latestRelease
			}
		}
	}

	return latestRelease
}

func GetAllReleaseDraftIds(userInfo base.UserInfo) []uint64 {
	var ids = make([]uint64, 0, 5)
	for _, release := range GetReleases(userInfo) {
		if release.Draft {
			ids = append(ids, release.Id)
		}
	}
	return ids
}
