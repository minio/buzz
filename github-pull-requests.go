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
	"net/http"
	"strings"

	"github.com/google/go-github/github"
)

// GitPR - a single pull request
type GitPR struct {
	Number      int            `json:"number"`
	ID          int64          `json:"id"`
	Title       string         `json:"title"`
	Labels      []github.Label `json:"name"`
	Sender      string         `json:"sender"`
	Assignees   string         `json:"login"`
	State       string         `json:"state"`
	UpdatedAt   int64          `json:"updated_at"`
	Repo        string         `json:"repo_name"`
	Link        string         `json:"html_url"`
	ReviewState []ReviewState
}

// Retrieve the relevant pull request data from an owner/repo.
func getPullRequests(owner string, repo string) (pullRequests []GitPR, err error) {
	prs, _, err := buzzClient.PullRequests.List(ctx, owner, repo, nil)

	if err != nil {
		return nil, err
	}
	pullRequests = []GitPR{}

	for _, elem := range prs {

		pullRequest := GitPR{}
		pullRequest.Number = *elem.Number
		pullRequest.Title = *elem.Title
		pullRequest.Sender = *elem.Head.User.Login
		for _, assignee := range elem.Assignees {
			pullRequest.Assignees += *assignee.Login + " "
		}
		pullRequest.UpdatedAt = elem.UpdatedAt.Unix()
		pullRequest.Repo = repo
		pullRequest.ID = *elem.ID
		pullRequest.Link = *elem.HTMLURL

		for _, label := range elem.Labels {
			pullRequest.Labels = append(pullRequest.Labels, *label)
		}

		err, states := getReviewStatesForPR(owner, pullRequest.Repo, pullRequest.Number)
		if err != nil {
			return nil, err
		}
		pullRequest.ReviewState = states

		pullRequests = append(pullRequests, pullRequest)
	}

	return pullRequests, nil
}

func getPRs(w http.ResponseWriter, req *http.Request) {
	pullRequests := []GitPR{}

	for _, locator := range config.RepoNames {
		// It's in the format owner/repo.
		parts := strings.SplitN(locator, "/", 2)
		prs, err := getPullRequests(parts[0], parts[1])
		if err != nil {
			w.WriteHeader(501)
			fmt.Fprint(w, err)
			return
		}

		pullRequests = append(pullRequests, prs...)
	}

	js, err := json.Marshal(pullRequests)
	if err != nil {
		w.WriteHeader(501)
		fmt.Fprint(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
