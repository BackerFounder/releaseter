package dao

import (
	"github.com/JZGoopi/releaseter/base"

	"github.com/JZGoopi/releaseter/api"
	"github.com/JZGoopi/releaseter/model"
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

			// 如果空了表示沒了（之後可以改判斷 http header 的 link，這樣可以少一次請求）
			if len(pagePulls) == 0 {
				break
			}

			for _, pull := range pagePulls {
				// 去除沒有被合併的
				if pull.MergeAt.IsZero() {
					continue
				}
				// 直到最新的 release 之後
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
