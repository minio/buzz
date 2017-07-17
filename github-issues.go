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
	"net/http"
	"strings"

	"github.com/google/go-github/github"
)

// GitIssue - holds only relevant issues
type GitIssue struct {
	Number    int    `json:"number"`
	Title     string `json:"title"`
	Link      string `json:"html_url"`
	Labels    []github.Label
	Assignees string `json:"login"`
	Milestone string `json:"milestone"`
	State     string `json:"state"`
	CreatedAt string `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	Repo      string `json:"repository_url"`
	ETA       string `json:"eta"`
}

func getIssuesForRepo(owner string, repo string) (issues []GitIssue, err error) {
	issuesGH, _, err := buzzClient.Issues.ListByRepo(ctx, owner, repo, nil)
	if err != nil {
		return nil, err
	}

	// iterate through each issue and scrape what buzz needs.
	for _, elem := range issuesGH {
		// var declaration automatically initialize flag var to 0
		issue := GitIssue{}

		if elem.PullRequestLinks != nil {
			continue
		}

		issue.Number = *elem.Number
		issue.Title = *elem.Title
		issue.Repo = owner + "/" + repo
		issue.Link = *elem.HTMLURL
		issue.CreatedAt = elem.CreatedAt.Format(buzzTimeLayout)
		issue.UpdatedAt = elem.UpdatedAt.Unix()
		issue.Labels = elem.Labels

		// Get the ETA for this issue.
		eta := getETAFromComment(owner, repo, *elem.Number)
		issue.ETA = eta.ETA
		// Ignore eta.Error, will be empty string, which is fine

		for _, assignee := range elem.Assignees {
			issue.Assignees += *assignee.Login + " "
		}

		if elem.Milestone != nil {
			issue.Milestone = *elem.Milestone.Title
			issue.State = *elem.Milestone.State
		}

		issues = append(issues, issue)
	}

	return issues, nil
}

func getIssues(w http.ResponseWriter, req *http.Request) {
	issues := []GitIssue{}

	for _, locator := range config.RepoNames {
		parts := strings.SplitN(locator, "/", 2)

		iss, err := getIssuesForRepo(parts[0], parts[1])
		if err != nil {
			w.WriteHeader(502)
			fmt.Fprint(w, err)
			return
		}

		issues = append(issues, iss...)
	}

	js, err := json.Marshal(issues)
	if err != nil {
		fmt.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
