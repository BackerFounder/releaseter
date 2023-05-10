package main

import (
	"fmt"

	"github.com/JZGoopi/releaseter/base"
	"github.com/JZGoopi/releaseter/data"
	"github.com/JZGoopi/releaseter/view"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	base.Init()
	data.Init()

	lablePulls := data.GetNewPullWithLables()
	template := view.GenCategoriesTemplate(lablePulls)

	// userInfo := base.GetUserInfo()
	// config := base.GetConfig()

	// if config.ClearHistoryDraft {
	// 	api.DelReleases(userInfo.Repo, userInfo.Token, dao.GetAllReleaseDraftIds(userInfo))
	// }
	// api.PostReleases(userInfo.Repo, userInfo.Token, model.GithubPostRelease{
	// 	TagName: data.GetTag(),
	// 	Name:    data.GetName(),
	// 	Body:    template,
	// 	Draft:   true,
	// })
	fmt.Println(template)
}
