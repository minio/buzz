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

import "fmt"

// ReviewStatus - holds Review info
type ReviewStatus struct {
	User struct {
		Login string `json:"login"`
	} `json:"user"`
}

// ReviewState - holds each PRs Reviewer's state
type ReviewState struct {
	User struct {
		Login string `json:"login"`
	} `json:"user"`
	State string `json:"state"`
}

var rStatus []ReviewStatus
var rReviews []ReviewState

func getReviewers(repo string, number int) []ReviewStatus {
	rStatus = nil
	// List all reviewers to an assigned PR # in a repo.
	rList, _, err := buzzClient.PullRequests.ListReviewers(ctx, "minio", repo, number, nil)
	if err != nil {
		// We should not exit on error in this case.
		fmt.Println("Unable to get Reviewers for %s", number)
	}
	for _, elem := range rList {
		eachReviewer := ReviewStatus{}
		eachReviewer.User.Login = elem.GetLogin()
		rStatus = append(rStatus, eachReviewer)
	}
	return rStatus
}

func getReviewStatesForPR(repo string, number int) []ReviewState {
	rReviews = nil
	// List all reviewers to an assigned PR # in a repo.
	rList, _, err := buzzClient.PullRequests.ListReviews(ctx, "minio", repo, number, nil)
	if err != nil {
		// We should not exit on error in this case.
		fmt.Println("Unable to get Reviewers for %s", number)
	}
	for _, elem := range rList {
		eachReview := ReviewState{}
		eachReview.User.Login = elem.User.GetLogin()
		eachReview.State = elem.GetState()
		rReviews = append(rReviews, eachReview)
	}
	return rReviews
}
