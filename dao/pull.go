package dao

import (
	"test/api"
	"test/model"
)

func GetNewPulls() model.Pulls {

	release := GetLatestRelease()
	userInfo := GetUserInfo()

	var page uint64 = 1
	var pulls = make(model.Pulls, 0, 50)

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

			if pull.MergeAt.After(release.CreatedAt) {
				pulls = append(pulls, pull)
			} else {
				break getpulls
			}
		}

		page++
	}
	return pulls
}

func GetNewPullWithLables() *model.LablePulls {
	pulls := GetNewPulls()
	categories := GetConfig().Categories
	lablePulls := make(model.LablePulls, len(categories))

	for _, category := range categories {

		lablePulls[category.Title] = make(model.Pulls, 0, 20)

		for _, pull := range pulls {

			for _, label := range pull.Labels {

				if label.Name == category.Label {
					lablePulls[category.Title] = append(lablePulls[category.Title], pull)
					continue
				}

				for _, cnfLabel := range category.Labels {
					if label.Name == cnfLabel {
						lablePulls[category.Title] = append(lablePulls[category.Title], pull)
					}
				}
			}
		}

		if len(lablePulls[category.Title]) == 0 {
			delete(lablePulls, category.Title)
		}

	}

	return &lablePulls
}
