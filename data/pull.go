package data

import (
	"test/base"
	"test/dao"
	"test/model"
)

var (
	labelPulls model.LablePulls
)

func GetNewPullWithLables() model.LablePulls {

	if labelPulls == nil {
		tags := dao.GetTags(userInfo)
		latestRelease := dao.GetLatestRelease(userInfo, tags)
		pulls := dao.GetNewPulls(userInfo, latestRelease)

		config := base.GetConfig()
		categories := config.Categories
		labelPulls = make(model.LablePulls, len(categories))

		for _, category := range categories {

			labelPulls[category.Title] = make(model.Pulls, 0, 20)

			for _, pull := range pulls {

				for _, label := range pull.Labels {

					if label.Name == category.Label {
						labelPulls[category.Title] = append(labelPulls[category.Title], pull)
						pull.Count++
						continue
					}

					for _, cnfLabel := range category.Labels {
						if label.Name == cnfLabel {
							labelPulls[category.Title] = append(labelPulls[category.Title], pull)
							pull.Count++
							break
						}
					}
				}
			}

			if len(labelPulls[category.Title]) == 0 {
				delete(labelPulls, category.Title)
			}
		}

		if config.CategoryOther.Show {
			otherTitle := config.CategoryOther.Title
			if otherTitle == "" {
				otherTitle = base.GetDefaultConfig().CategoryOther.Title
			}

			for _, pull := range pulls {
				if pull.Count == 0 {
					labelPulls[otherTitle] = append(labelPulls[otherTitle], pull)
				}
			}
		}
	}

	return labelPulls
}
