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
		config := base.GetConfig()
		tags := dao.GetTags(userInfo)
		latestRelease := dao.GetLatestRelease(userInfo, tags, config)
		pulls := dao.GetNewPulls(userInfo, latestRelease)

		categories := config.Categories
		labelPulls = make(model.LablePulls, 0, len(categories))

		for _, category := range categories {

			categoryPulls := make(model.Pulls, 0, 20)

		foreachPulls:
			for _, pull := range pulls {

				for _, label := range pull.Labels {
					for _, exceptLable := range config.CategoryExceptLables {
						if label.Name == exceptLable {
							pull.NoRelease = true
							continue foreachPulls
						}
					}
				}

				for _, label := range pull.Labels {

					if label.Name == category.Label {
						categoryPulls = append(categoryPulls, pull)
						pull.Count++
						continue foreachPulls
					}

					for _, cnfLabel := range category.Labels {
						if label.Name == cnfLabel {
							categoryPulls = append(categoryPulls, pull)
							pull.Count++
							continue foreachPulls
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
				if pull.Count == 0 && !pull.NoRelease {
					otherPulls = append(otherPulls, pull)
				}
			}

			if len(otherPulls) != 0 {
				labelPulls = append(labelPulls, model.LablePull{
					Pulls: otherPulls,
					Title: otherTitle,
				})
			}
		}
	}

	return labelPulls
}
