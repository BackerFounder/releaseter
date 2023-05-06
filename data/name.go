package data

import (
	"test/base"
	"test/consts"
)

var (
	name       string
	isNameInit = false
)

func GetName() string {
	if !isNameInit {
		nameTemplate := base.GetConfig().NameTemplate
		version := GetVersion()

		if nameTemplate == consts.EMPTY {
			name = ""
			isNameInit = true
			return name
		}

		nameTemplate = replaceVersionPlaceholder(nameTemplate, version)
		nameTemplate = replaceTimePlaceholder(nameTemplate)

		name = nameTemplate
		isNameInit = true
	}

	return name
}
