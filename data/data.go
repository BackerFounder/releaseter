package data

import (
	"strings"

	"github.com/JZGoopi/releaseter/model"

	"time"

	"github.com/JZGoopi/releaseter/base"
	"github.com/JZGoopi/releaseter/consts"
)

var (
	userInfo base.UserInfo
)

func InitUserInfo() {
	userInfo = base.GetUserInfo()
}

func replaceVersionPlaceholder(str string, version *model.TagVersion) string {
	config := base.GetConfig()
	tagPreRelease := config.TagPreRelease
	tagBuild := config.TagBuild
	for _, placeholder := range consts.GetVersionPlaceholders() {
		str = strings.ReplaceAll(str, placeholder, version.Join(placeholder, tagPreRelease, tagBuild))
	}
	return str
}

func replaceTimePlaceholder(str string) string {
	config := base.GetConfig()
	loc, err := time.LoadLocation(config.TimeLocation)
	workflowTime := base.GetWorkflowTime()
	if err != nil {
		loc, _ = time.LoadLocation(base.GetDefaultConfig().TagTemplate)
	}

	str = strings.ReplaceAll(str, consts.TIME_WORKTIME, workflowTime.In(loc).Format(config.TimeFormat))

	return str
}

func Init() {
	InitUserInfo()
}
