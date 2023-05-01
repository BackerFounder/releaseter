package model

import "time"

type GithubGetLabel struct {
	Name string `json:"name"`
}

type GithubGetPull struct {
	Number         uint64            `json:"number"`
	State          string            `json:"state"`
	Title          string            `json:"title"`
	MergeAt        time.Time         `json:"merged_at"`
	MergeCommitSHA string            `json:"merge_commit_sha"`
	Labels         []*GithubGetLabel `json:"labels"`
}

type Pulls []*GithubGetPull
type LablePulls map[string]Pulls
