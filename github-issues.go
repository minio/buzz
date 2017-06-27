/*
 * Buzz, (C) 2016,2017 Minio, Inc.
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

// Label - holds issue labels
type Label struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

// GitIssues - holds only relevant issues
type GitIssues struct {
	Number    int    `json:"number"`
	Title     string `json:"title"`
	Link      string `json:"html_url"`
	Labels    []Label
	Assignees string `json:"login"`
	Milestone string `json:"milestone"`
	State     string `json:"state"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Repo      string `json:"repository_url"`
	Hours     int64  `json:"hours"`
}

// RepoIssues - holds all issues per repo.
type RepoIssues struct {
	URL           string `json:"url"`
	RepositoryURL string `json:"repository_url"`
	HTMLURL       string `json:"html_url"`
	ID            int    `json:"id"`
	Number        int    `json:"number"`
	Title         string `json:"title"`
	User          struct {
		Login string `json:"login"`
	} `json:"user"`
	Labels []struct {
		ID      int    `json:"id"`
		URL     string `json:"url"`
		Name    string `json:"name"`
		Color   string `json:"color"`
		Default bool   `json:"default"`
	} `json:"labels"`
	Assignee struct {
		Login string `json:"login"`
	} `json:"assignee"`
	Assignees []struct {
		Login string `json:"login"`
	} `json:"assignees"`
	Milestone struct {
		Title string `json:"title"`
		State string `json:"state"`
	} `json:"milestone"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	ClosedAt    interface{} `json:"closed_at"`
	PullRequest struct {
		URL string `json:"url"`
	} `json:"pull_request"`
}

var gIssues []GitIssues

func populateIssues(url string) {
	pullURL := url + token
	fmt.Println("Fetching from....", pullURL)
	resp, err := http.Get(pullURL)
	exitOnErr(err)

	defer resp.Body.Close()
	htmlData, err := ioutil.ReadAll(resp.Body)
	exitOnErr(err)

	mIssues := []RepoIssues{}
	json.Unmarshal(htmlData, &mIssues)

	// iterate through each issue and scrape what buzz needs.
	for _, elem := range mIssues {
		// var declaration automatically initialize flag var to 0
		var flag int
		eachGitIssue := GitIssues{}

		if elem.PullRequest.URL != "" {
			flag = 1
		} else {
			eachGitIssue.Number = elem.Number
			eachGitIssue.Title = elem.Title
			eachGitIssue.Repo = elem.RepositoryURL
			eachGitIssue.Link = elem.HTMLURL
			eachGitIssue.CreatedAt = elem.CreatedAt.Format(buzzTimeLayout)
			eachGitIssue.UpdatedAt = elem.UpdatedAt.Format(buzzTimeLayout)
			delta := elem.UpdatedAt.Sub(elem.CreatedAt)
			eachGitIssue.Hours = int64(delta.Hours())

			// iterate to get all labels and colors.
			for _, labe := range elem.Labels {
				eachGitIssue.Labels = append(eachGitIssue.Labels, Label{
					Name:  labe.Name,
					Color: labe.Color,
				})
			}
			for _, assignee := range elem.Assignees {
				eachGitIssue.Assignees += assignee.Login + " "
			}
			eachGitIssue.Milestone = elem.Milestone.Title
			eachGitIssue.State = elem.Milestone.State
		}
		if flag != 1 {
			gIssues = append(gIssues, eachGitIssue)
		}
	} // end of for
}

func getIssues(w http.ResponseWriter, req *http.Request) {
	gIssues = nil
	// One Minio.
	for _, rName := range config.RepoNames {
		populateIssues(`https://api.github.com/repos/` + rName + `/issues?state=open&per_page=100&access_token=`)
	}
	js, err := json.Marshal(gIssues)
	if err != nil {
		fmt.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
