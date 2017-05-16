/*
 * Buzz, (C) 2017 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GitPRs - holds only relevant PRs
type GitPRs struct {
	Number    int    `json:"number"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Labels    string `json:"name"`
	Sender    string `json:"sender"`
	Assignees string `json:"login"`
	State     string `json:"state"`
	UpdatedAt string `json:"updated_at"`
	Repo      string `json:"repo_name"`
	Link      string `json:"html_url"`
	Hours     int64  `json:"hours"`
	Reviewers []ReviewStatus
}

// PRIssues - holds all PRs
type PRIssues struct {
	ID      int    `json:"id"`
	HTMLURL string `json:"html_url"`
	Number  int    `json:"number"`
	State   string `json:"state"`
	Title   string `json:"title"`
	User    struct {
		Login string `json:"login"`
	} `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Assignee  struct {
		Login string `json:"login"`
	} `json:"assignee"`
	Assignees []struct {
		Login string `json:"login"`
	} `json:"assignees"`
	Head struct {
		User struct {
			Login string `json:"login"`
		} `json:"user"`
		Repo struct {
			Name string `json:"name"`
		} `json:"repo"`
	} `json:"head"`
}

var pRequests []GitPRs

func populatePRs(rName string, url string) {
	pullURL := url + token
	fmt.Println("Fetching from....", pullURL)
	resp, err := http.Get(pullURL)
	exitOnErr(err)

	defer resp.Body.Close()
	htmlData, err := ioutil.ReadAll(resp.Body)
	exitOnErr(err)

	pIssues := []PRIssues{}
	json.Unmarshal(htmlData, &pIssues)

	for _, elem := range pIssues {
		eachPRIssue := GitPRs{}
		eachPRIssue.Number = elem.Number
		eachPRIssue.Title = elem.Title
		eachPRIssue.Sender = elem.Head.User.Login
		for _, assignee := range elem.Assignees {
			eachPRIssue.Assignees += assignee.Login + " "
		}
		eachPRIssue.UpdatedAt = elem.UpdatedAt.Format(buzzTimeLayout)
		delta := elem.UpdatedAt.Sub(elem.CreatedAt)
		eachPRIssue.Hours = int64(delta.Hours())
		eachPRIssue.Repo = elem.Head.Repo.Name
		eachPRIssue.ID = elem.ID
		eachPRIssue.Link = elem.HTMLURL
		eachPRIssue.Reviewers = getReviewers(eachPRIssue.Repo, eachPRIssue.Number)
		pRequests = append(pRequests, eachPRIssue)
	} // end of for
}

func getPRs(w http.ResponseWriter, req *http.Request) {
	pRequests = nil
	// One Minio.
	for _, rName := range config.RepoNames {
		populatePRs(rName, `https://api.github.com/repos/`+rName+`/pulls?state=open&per_page=100&access_token=`)
	}
	js, err := json.Marshal(pRequests)
	if err != nil {
		fmt.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
