package main

import (
	"test/api"
	"test/dao"
	"test/model"
	"test/view"
)

func main() {

	lablePulls := dao.GetNewPullWithLables()
	template := view.GenCategoriesTemplate(lablePulls)
	cnf := dao.GetConfig()

	userInfo := dao.GetUserInfo()

	api.DelReleases(userInfo.Repo, userInfo.Token, dao.GetAllReleaseDraftIds())
	api.PostReleases(userInfo.Repo, userInfo.Token, model.GithubPostRelease{
		TagName: cnf.TagTemplate,
		Name:    cnf.NameTemplate,
		Body:    template,
		Draft:   true,
	})
}
