/*
 * Buggy, (C) 2017 Minio, Inc.
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
	Assignees string `json:"login"`
	State     string `json:"state"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Repo      string `json:"repo_name"`
}

// PRIssues - holds all PRs
type PRIssues struct {
	URL      string `json:"url"`
	ID       int    `json:"id"`
	HTMLURL  string `json:"html_url"`
	DiffURL  string `json:"diff_url"`
	PatchURL string `json:"patch_url"`
	IssueURL string `json:"issue_url"`
	Number   int    `json:"number"`
	State    string `json:"state"`
	Locked   bool   `json:"locked"`
	Title    string `json:"title"`
	User     struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"user"`
	Body           string      `json:"body"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	ClosedAt       interface{} `json:"closed_at"`
	MergedAt       interface{} `json:"merged_at"`
	MergeCommitSha string      `json:"merge_commit_sha"`
	Assignee       struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"assignee"`
	Assignees []struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"assignees"`
	Milestone         interface{} `json:"milestone"`
	CommitsURL        string      `json:"commits_url"`
	ReviewCommentsURL string      `json:"review_comments_url"`
	ReviewCommentURL  string      `json:"review_comment_url"`
	CommentsURL       string      `json:"comments_url"`
	StatusesURL       string      `json:"statuses_url"`
	Head              struct {
		Label string `json:"label"`
		Ref   string `json:"ref"`
		Sha   string `json:"sha"`
		User  struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"user"`
		Repo struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			FullName string `json:"full_name"`
			Owner    struct {
				Login             string `json:"login"`
				ID                int    `json:"id"`
				AvatarURL         string `json:"avatar_url"`
				GravatarID        string `json:"gravatar_id"`
				URL               string `json:"url"`
				HTMLURL           string `json:"html_url"`
				FollowersURL      string `json:"followers_url"`
				FollowingURL      string `json:"following_url"`
				GistsURL          string `json:"gists_url"`
				StarredURL        string `json:"starred_url"`
				SubscriptionsURL  string `json:"subscriptions_url"`
				OrganizationsURL  string `json:"organizations_url"`
				ReposURL          string `json:"repos_url"`
				EventsURL         string `json:"events_url"`
				ReceivedEventsURL string `json:"received_events_url"`
				Type              string `json:"type"`
				SiteAdmin         bool   `json:"site_admin"`
			} `json:"owner"`
			Private          bool        `json:"private"`
			HTMLURL          string      `json:"html_url"`
			Description      string      `json:"description"`
			Fork             bool        `json:"fork"`
			URL              string      `json:"url"`
			ForksURL         string      `json:"forks_url"`
			KeysURL          string      `json:"keys_url"`
			CollaboratorsURL string      `json:"collaborators_url"`
			TeamsURL         string      `json:"teams_url"`
			HooksURL         string      `json:"hooks_url"`
			IssueEventsURL   string      `json:"issue_events_url"`
			EventsURL        string      `json:"events_url"`
			AssigneesURL     string      `json:"assignees_url"`
			BranchesURL      string      `json:"branches_url"`
			TagsURL          string      `json:"tags_url"`
			BlobsURL         string      `json:"blobs_url"`
			GitTagsURL       string      `json:"git_tags_url"`
			GitRefsURL       string      `json:"git_refs_url"`
			TreesURL         string      `json:"trees_url"`
			StatusesURL      string      `json:"statuses_url"`
			LanguagesURL     string      `json:"languages_url"`
			StargazersURL    string      `json:"stargazers_url"`
			ContributorsURL  string      `json:"contributors_url"`
			SubscribersURL   string      `json:"subscribers_url"`
			SubscriptionURL  string      `json:"subscription_url"`
			CommitsURL       string      `json:"commits_url"`
			GitCommitsURL    string      `json:"git_commits_url"`
			CommentsURL      string      `json:"comments_url"`
			IssueCommentURL  string      `json:"issue_comment_url"`
			ContentsURL      string      `json:"contents_url"`
			CompareURL       string      `json:"compare_url"`
			MergesURL        string      `json:"merges_url"`
			ArchiveURL       string      `json:"archive_url"`
			DownloadsURL     string      `json:"downloads_url"`
			IssuesURL        string      `json:"issues_url"`
			PullsURL         string      `json:"pulls_url"`
			MilestonesURL    string      `json:"milestones_url"`
			NotificationsURL string      `json:"notifications_url"`
			LabelsURL        string      `json:"labels_url"`
			ReleasesURL      string      `json:"releases_url"`
			DeploymentsURL   string      `json:"deployments_url"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			PushedAt         time.Time   `json:"pushed_at"`
			GitURL           string      `json:"git_url"`
			SSHURL           string      `json:"ssh_url"`
			CloneURL         string      `json:"clone_url"`
			SvnURL           string      `json:"svn_url"`
			Homepage         string      `json:"homepage"`
			Size             int         `json:"size"`
			StargazersCount  int         `json:"stargazers_count"`
			WatchersCount    int         `json:"watchers_count"`
			Language         string      `json:"language"`
			HasIssues        bool        `json:"has_issues"`
			HasProjects      bool        `json:"has_projects"`
			HasDownloads     bool        `json:"has_downloads"`
			HasWiki          bool        `json:"has_wiki"`
			HasPages         bool        `json:"has_pages"`
			ForksCount       int         `json:"forks_count"`
			MirrorURL        interface{} `json:"mirror_url"`
			OpenIssuesCount  int         `json:"open_issues_count"`
			Forks            int         `json:"forks"`
			OpenIssues       int         `json:"open_issues"`
			Watchers         int         `json:"watchers"`
			DefaultBranch    string      `json:"default_branch"`
		} `json:"repo"`
	} `json:"head"`
	Base struct {
		Label string `json:"label"`
		Ref   string `json:"ref"`
		Sha   string `json:"sha"`
		User  struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"user"`
		Repo struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			FullName string `json:"full_name"`
			Owner    struct {
				Login             string `json:"login"`
				ID                int    `json:"id"`
				AvatarURL         string `json:"avatar_url"`
				GravatarID        string `json:"gravatar_id"`
				URL               string `json:"url"`
				HTMLURL           string `json:"html_url"`
				FollowersURL      string `json:"followers_url"`
				FollowingURL      string `json:"following_url"`
				GistsURL          string `json:"gists_url"`
				StarredURL        string `json:"starred_url"`
				SubscriptionsURL  string `json:"subscriptions_url"`
				OrganizationsURL  string `json:"organizations_url"`
				ReposURL          string `json:"repos_url"`
				EventsURL         string `json:"events_url"`
				ReceivedEventsURL string `json:"received_events_url"`
				Type              string `json:"type"`
				SiteAdmin         bool   `json:"site_admin"`
			} `json:"owner"`
			Private          bool        `json:"private"`
			HTMLURL          string      `json:"html_url"`
			Description      string      `json:"description"`
			Fork             bool        `json:"fork"`
			URL              string      `json:"url"`
			ForksURL         string      `json:"forks_url"`
			KeysURL          string      `json:"keys_url"`
			CollaboratorsURL string      `json:"collaborators_url"`
			TeamsURL         string      `json:"teams_url"`
			HooksURL         string      `json:"hooks_url"`
			IssueEventsURL   string      `json:"issue_events_url"`
			EventsURL        string      `json:"events_url"`
			AssigneesURL     string      `json:"assignees_url"`
			BranchesURL      string      `json:"branches_url"`
			TagsURL          string      `json:"tags_url"`
			BlobsURL         string      `json:"blobs_url"`
			GitTagsURL       string      `json:"git_tags_url"`
			GitRefsURL       string      `json:"git_refs_url"`
			TreesURL         string      `json:"trees_url"`
			StatusesURL      string      `json:"statuses_url"`
			LanguagesURL     string      `json:"languages_url"`
			StargazersURL    string      `json:"stargazers_url"`
			ContributorsURL  string      `json:"contributors_url"`
			SubscribersURL   string      `json:"subscribers_url"`
			SubscriptionURL  string      `json:"subscription_url"`
			CommitsURL       string      `json:"commits_url"`
			GitCommitsURL    string      `json:"git_commits_url"`
			CommentsURL      string      `json:"comments_url"`
			IssueCommentURL  string      `json:"issue_comment_url"`
			ContentsURL      string      `json:"contents_url"`
			CompareURL       string      `json:"compare_url"`
			MergesURL        string      `json:"merges_url"`
			ArchiveURL       string      `json:"archive_url"`
			DownloadsURL     string      `json:"downloads_url"`
			IssuesURL        string      `json:"issues_url"`
			PullsURL         string      `json:"pulls_url"`
			MilestonesURL    string      `json:"milestones_url"`
			NotificationsURL string      `json:"notifications_url"`
			LabelsURL        string      `json:"labels_url"`
			ReleasesURL      string      `json:"releases_url"`
			DeploymentsURL   string      `json:"deployments_url"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			PushedAt         time.Time   `json:"pushed_at"`
			GitURL           string      `json:"git_url"`
			SSHURL           string      `json:"ssh_url"`
			CloneURL         string      `json:"clone_url"`
			SvnURL           string      `json:"svn_url"`
			Homepage         string      `json:"homepage"`
			Size             int         `json:"size"`
			StargazersCount  int         `json:"stargazers_count"`
			WatchersCount    int         `json:"watchers_count"`
			Language         string      `json:"language"`
			HasIssues        bool        `json:"has_issues"`
			HasProjects      bool        `json:"has_projects"`
			HasDownloads     bool        `json:"has_downloads"`
			HasWiki          bool        `json:"has_wiki"`
			HasPages         bool        `json:"has_pages"`
			ForksCount       int         `json:"forks_count"`
			MirrorURL        interface{} `json:"mirror_url"`
			OpenIssuesCount  int         `json:"open_issues_count"`
			Forks            int         `json:"forks"`
			OpenIssues       int         `json:"open_issues"`
			Watchers         int         `json:"watchers"`
			DefaultBranch    string      `json:"default_branch"`
		} `json:"repo"`
	} `json:"base"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		Issue struct {
			Href string `json:"href"`
		} `json:"issue"`
		Comments struct {
			Href string `json:"href"`
		} `json:"comments"`
		ReviewComments struct {
			Href string `json:"href"`
		} `json:"review_comments"`
		ReviewComment struct {
			Href string `json:"href"`
		} `json:"review_comment"`
		Commits struct {
			Href string `json:"href"`
		} `json:"commits"`
		Statuses struct {
			Href string `json:"href"`
		} `json:"statuses"`
	} `json:"_links"`
}

var pRequests []GitPRs

func populatePRs(url string) {
	pullURL := url + token
	fmt.Println("Fetching from....", pullURL)
	resp, err := http.Get(pullURL)
	exitOnErr(err)

	defer resp.Body.Close()
	htmlData, err := ioutil.ReadAll(resp.Body)
	exitOnErr(err)

	pIssues := []PRIssues{}
	json.Unmarshal(htmlData, &pIssues)

	//json.Unmarshal(htmlData, &gIssues)
	const layout = "Jan 2, 2006 at 3:04pm (PST)"
	for _, elem := range pIssues {
		eachPRIssue := GitPRs{}
		eachPRIssue.Number = elem.Number
		eachPRIssue.Title = elem.Title

		/*for _, labe := range elem.Labels {
		    eachGitIssue.Labels += labe.Name + " "
		}*/
		for _, assignee := range elem.Assignees {
			eachPRIssue.Assignees += assignee.Login + " "
		}

		eachPRIssue.CreatedAt = elem.CreatedAt.Format(layout)
		eachPRIssue.UpdatedAt = elem.UpdatedAt.Format(layout)
		eachPRIssue.Repo = elem.Head.Repo.Name
		eachPRIssue.ID = elem.ID
		pRequests = append(pRequests, eachPRIssue)

	} // end of for

}

func getPRs(w http.ResponseWriter, req *http.Request) {
	pRequests = nil
	// One Minio.
	for _, rName := range config.RepoNames {
		populatePRs(`https://api.github.com/repos/` + rName + `/pulls?state=open&per_page=100&access_token=`)
	}

	js, err := json.Marshal(pRequests)
	if err != nil {
		fmt.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
