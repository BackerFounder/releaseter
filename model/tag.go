package model

import (
	"fmt"
	"test/consts"
)

type GithubGetTagWithCommit struct {
	Name   string `json:"name"`
	Commit struct {
		Sha string `json:"sha"`
	} `json:"commit"`
}

type GithubGetTagWithObject struct {
	Object struct {
		Sha string `json:"sha"`
	} `json:"object"`
}

type Tag struct {
	Name string
	Sha  string
}

type TagVersion struct {
	Major         uint64
	Minor         uint64
	Patch         uint64
	PreRelease    string
	BuildMetaData string
}

// next: major | minor | patch
func (tv *TagVersion) Join(next, preRelease, build string) string {

	versions := [3]uint64{tv.Major, tv.Minor, tv.Patch}
	for idx, placeholder := range consts.GetVersionPlaceholders() {
		if next == placeholder {
			versions[idx]++
		}
	}

	versionString := fmt.Sprintf("%d.%d.%d", versions[0], versions[1], versions[2])
	if preRelease != "" {
		versionString = fmt.Sprintf("%s-%s", versionString, preRelease)
	}

	if build != "" {
		versionString = fmt.Sprintf("%s+%s", versionString, build)
	}

	return versionString
}
