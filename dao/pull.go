package dao

import (
	"test/api"
	"test/base"

	"test/model"
)

var (
	newPulls model.Pulls
)

func GetNewPulls(userInfo base.UserInfo, latestRelease *model.Release) model.Pulls {

	if newPulls == nil {
		var page uint64 = 1
		newPulls = make(model.Pulls, 0, 50)

	getpulls:
		for {
			pagePulls := api.GetPulls(userInfo.Repo, userInfo.Token, page)

			if len(pagePulls) == 0 {
				break
			}

			for _, pull := range pagePulls {
				if pull.MergeAt.IsZero() {
					continue
				}

				if pull.MergeAt.After(latestRelease.CreatedAt) {
					newPulls = append(newPulls, pull)
				} else {
					break getpulls
				}
			}

			page++
		}
	}

	return newPulls
}
