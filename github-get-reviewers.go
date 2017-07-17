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
	"time"

	"github.com/google/go-github/github"
)

// ReviewState - holds each PRs Reviewer's state
type ReviewState struct {
	User struct {
		Login string `json:"login"`
	} `json:"user"`
	State string `json:"state"`
}

func getReviewStatesForPR(repo string, number int) (error, []ReviewState) {
	reviews := make(map[string]*github.PullRequestReview)

	// Get the username of the PR author.
	pr, _, err := buzzClient.PullRequests.Get(ctx, "minio", repo, number)
	if err != nil {
		return err, nil
	}

	author := *pr.User.Login

	// List all reviewers to an assigned PR # in a repo.
	reviewers, _, err := buzzClient.PullRequests.ListReviewers(ctx, "minio", repo, number, nil)
	if err != nil {
		return err, nil
	}

	for _, elem := range reviewers {
		review := &github.PullRequestReview{}
		review.User = elem

		// Any reviews are later than this assignment, so put it farthest back.
		assignDate := time.Unix(0, 0)
		review.SubmittedAt = &assignDate

		// Assume their review is pending until further notice.
		status := "PENDING"
		review.State = &status

		reviews[elem.GetLogin()] = review
	}

	// List all who voluntarily gave a review.
	voluntaryReviews, _, err := buzzClient.PullRequests.ListReviews(ctx, "minio", repo, number, nil)
	if err != nil {
		return err, nil
	}

	for _, elem := range voluntaryReviews {
		username := elem.User.GetLogin()

		if _, ok := reviews[username]; ok {
			// Comments never take priority.
			if elem.GetState() == "COMMENTED" {
				continue
			}

			// If they've been assigned previously and are
			// continuing with their review, compare the time of this
			// review to their previous review, replacing it if it's newer.
			if reviews[username].SubmittedAt.Before(*elem.SubmittedAt) {
				reviews[username] = elem
			}
		} else {
			// If a user was not assigned but has given a review anyways,
			// add their review.
			reviews[username] = elem
		}
	}

	// Convert the map to an array.
	rReviews := make([]ReviewState, 0, len(reviews))
	for username, latest := range reviews {
		// Do not include the author in the list.
		if username == author {
			continue
		}

		rReview := ReviewState{}
		rReview.User.Login = username
		rReview.State = latest.GetState()

		rReviews = append(rReviews, rReview)
	}

	return nil, rReviews
}
