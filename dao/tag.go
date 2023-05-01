package dao

import (
	"regexp"
	"strings"
	"test/api"
	"test/consts"
	"test/model"
	"test/utils"
)

var (
	tags    []*model.Tag
	version *model.TagVersion
)

func GetTags() []*model.Tag {
	userInfo := GetUserInfo()

	if tags == nil {
		tags = api.GetTags(userInfo.Repo, userInfo.Token)
	}
	return tags
}

func GetVersion() *model.TagVersion {
	if version == nil {

		tags := GetTags()
		tagTemplate := GetOriginConfig().TagTemplate
		placeholders := consts.GetVersionPlaceholders()
		version = &model.TagVersion{}

		for _, placeholder := range placeholders {
			if strings.Contains(tagTemplate, placeholder) {
				reStr := `(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?`
				reStr = "^" + strings.Replace(tagTemplate, placeholder, reStr, 1) + "$"
				re := regexp.MustCompile(reStr)

				for _, tag := range tags {
					matches := re.FindStringSubmatch(tag.Name)
					if len(matches) != 0 {
						version.Major = utils.StrToUint64(matches[1])
						version.Minor = utils.StrToUint64(matches[2])
						version.Patch = utils.StrToUint64(matches[3])
						version.PreRelease = matches[4]
						version.BuildMetaData = matches[5]
						return version
					}
				}
				return version
			}
		}

	}
	return version
}
