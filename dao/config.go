package dao

// var (
// 	defaultConfig *model.Config
// 	userConfig    *model.Config
// )

// func initUserConfig() {
// 	userConfigByteData := config.GetConfig()
// 	userConfig = &model.Config{}
// 	err := yaml.Unmarshal(userConfigByteData, userConfig)
// 	if err != nil {
// 		panic(err)
// 	}
// 	initTimeSetting()

// 	version := GetVersion(userConfig.TagTemplate)
// 	initTagTemplate(version)
// 	initNameTemplate(version)
// }

// func initTimeSetting() {
// 	if userConfig.TimeLocation == "" {
// 		userConfig.TimeLocation = defaultConfig.TimeLocation
// 	}

// 	if userConfig.TimeFormat == "" {
// 		userConfig.TimeFormat = defaultConfig.TimeFormat
// 	}
// }

// func initTagTemplate(version *model.TagVersion) {
// 	tagTemplate := userConfig.TagTemplate

// 	if tagTemplate == consts.GetEmptyPlaceholder() {
// 		userConfig.TagTemplate = ""
// 		return
// 	}

// 	if tagTemplate == "" {
// 		tagTemplate = defaultConfig.TagTemplate
// 	}

// 	tagTemplate = replaceVersionPlaceholder(tagTemplate, version)
// 	tagTemplate = replaceTimePlaceholder(tagTemplate)

// 	userConfig.TagTemplate = tagTemplate
// }

// func initNameTemplate(version *model.TagVersion) {

// }

// func GetConfig() *model.Config {
// 	return userConfig
// }

// func init() {
// 	initDefaultConfig()
// 	initUserConfig()
// }
