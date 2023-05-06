package data

import (
	"regexp"
	"strings"
	"test/base"
	"test/consts"
	"test/dao"
	"test/model"
	"test/utils"
)

const (
	RE_STR_BASE = `(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?`
)

var (
	version *model.TagVersion
)

func GetVersion() *model.TagVersion {
	if version == nil {

		tagTemplate := base.GetConfig().TagTemplate
		tags := dao.GetTags(userInfo)
		placeholders := consts.GetVersionPlaceholders()
		version = &model.TagVersion{}

		for _, placeholder := range placeholders {
			if strings.Contains(tagTemplate, placeholder) {
				reStr := "^" + strings.Replace(tagTemplate, placeholder, RE_STR_BASE, 1) + "$"
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
