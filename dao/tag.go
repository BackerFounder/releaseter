package dao

import (
	"test/api"
	"test/base"
	"test/model"
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
