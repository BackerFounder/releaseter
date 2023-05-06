package data

import (
	"github.com/JZGoopi/releaseter/base"
	"github.com/JZGoopi/releaseter/dao"
	"github.com/JZGoopi/releaseter/model"
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
		labelPulls = make(model.LablePulls, 0, len(categories))

		for _, category := range categories {

			categoryPulls := make(model.Pulls, 0, 20)

			for _, pull := range pulls {

			findLable:
				for _, label := range pull.Labels {

					if label.Name == category.Label {
						categoryPulls = append(categoryPulls, pull)
						pull.Count++
						continue findLable
					}

					for _, cnfLabel := range category.Labels {
						if label.Name == cnfLabel {
							categoryPulls = append(categoryPulls, pull)
							pull.Count++
							break findLable
						}
					}
				}
			}

			if len(categoryPulls) != 0 {
				labelPulls = append(labelPulls, model.LablePull{
					Pulls: categoryPulls,
					Title: category.Title,
				})
			}
		}

		if config.CategoryOther.Show {
			otherTitle := config.CategoryOther.Title
			if otherTitle == "" {
				otherTitle = base.GetDefaultConfig().CategoryOther.Title
			}
			otherPulls := make(model.Pulls, 0, 20)
			for _, pull := range pulls {
				if pull.Count == 0 {
					otherPulls = append(otherPulls, pull)
				}
			}

			labelPulls = append(labelPulls, model.LablePull{
				Pulls: otherPulls,
				Title: otherTitle,
			})
		}
	}

	return labelPulls
}
