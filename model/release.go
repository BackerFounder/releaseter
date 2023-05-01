package model

import "time"

type GithubGetRelease struct {
	TagName   string    `json:"tag_name"`
	Draft     bool      `json:"draft"`
	Id        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type GithubPostRelease struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	Body    string `json:"body"`
	Draft   bool   `json:"draft"`
}

type Release struct {
	Sha       string
	TagName   string    `json:"tag_name"`
	CreatedAt time.Time `json:"created_at"`
}
