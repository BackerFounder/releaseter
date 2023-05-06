package api

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/JZGoopi/releaseter/model"
)

func GetPulls(repo, token string, page uint64) []*model.GithubGetPull {

	url := fmt.Sprintf("https://api.github.com/repos/%s/pulls", repo)
	headers := genBaseHeaders(token)
	query := Querys{}
	query["state"] = "closed"
	query["sort"] = "updated"
	query["direction"] = "desc"
	query["per_page"] = "100"
	query["page"] = strconv.Itoa(int(page))
	data := getData(url, headers, query)
	var pull []*model.GithubGetPull
	json.Unmarshal(data, &pull)
	return pull
}
