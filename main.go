package main

import (
	"os"
	"test/api"
	"test/base"
	"test/dao"
	"test/data"
	"test/model"
	"test/view"
)

func main() {

	os.Setenv("JZGOOPI_REPEASETER_CONFIG_PATH", "releaster.yml")
	base.Init()
	data.Init()

	lablePulls := data.GetNewPullWithLables()
	template := view.GenCategoriesTemplate(lablePulls)

	userInfo := base.GetUserInfo()

	api.DelReleases(userInfo.Repo, userInfo.Token, dao.GetAllReleaseDraftIds(userInfo))
	api.PostReleases(userInfo.Repo, userInfo.Token, model.GithubPostRelease{
		TagName: data.GetTag(),
		Name:    data.GetName(),
		Body:    template,
		Draft:   true,
	})
}
