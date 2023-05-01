package dao

import (
	"test/config"
	"test/model"
)

var (
	userInfo *model.UserInfo
)

func GetUserInfo() *model.UserInfo {
	if userInfo == nil {
		userInfo = &model.UserInfo{
			Token: config.GetToken(),
			Repo:  config.GetRepo(),
		}
	}
	return userInfo

}
