package dao

import (
	"github.com/JZGoopi/releaseter/base"

	"github.com/JZGoopi/releaseter/api"
	"github.com/JZGoopi/releaseter/model"
)

var (
	tags []*model.Tag
)

func GetTags(userInfo base.UserInfo) []*model.Tag {

	if tags == nil {
		tags = api.GetTags(userInfo.Repo, userInfo.Token)
	}
	return tags
}
