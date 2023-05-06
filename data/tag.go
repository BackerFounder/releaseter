package data

import (
	"github.com/JZGoopi/releaseter/base"
	"github.com/JZGoopi/releaseter/consts"
)

var (
	tag       string
	isTagInit = false
)

func GetTag() string {
	if !isTagInit {
		tagTemplate := base.GetConfig().TagTemplate
		version := GetVersion()

		if tagTemplate == consts.EMPTY {
			tag = tagTemplate
			isTagInit = true
			return ""
		}

		tagTemplate = replaceVersionPlaceholder(tagTemplate, version)
		tagTemplate = replaceTimePlaceholder(tagTemplate)

		tag = tagTemplate
		isTagInit = true
	}

	return tag
}
