package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// GitIssues - holds only relevant issues
type GitIssues struct {
	Number    int    `json:"number"`
	Title     string `json:"title"`
	Labels    string `json:"name"`
	Assignees string `json:"login"`
	Milestone string `json:"milestone"`
	State     string `json:"state"`
	Repo      string `json:"repository_url"`
}

// Repoissue - holds all issues per repo.
type Repoissue struct {
	URL           string `json:"url"`
	RepositoryURL string `json:"repository_url"`
	LabelsURL     string `json:"labels_url"`
	CommentsURL   string `json:"comments_url"`
	EventsURL     string `json:"events_url"`
	HTMLURL       string `json:"html_url"`
	ID            int    `json:"id"`
	Number        int    `json:"number"`
	Title         string `json:"title"`
	User          struct {
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
	Labels []struct {
		ID      int    `json:"id"`
		URL     string `json:"url"`
		Name    string `json:"name"`
		Color   string `json:"color"`
		Default bool   `json:"default"`
	} `json:"labels"`
	State    string `json:"state"`
	Locked   bool   `json:"locked"`
	Assignee struct {
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
	Milestone struct {
		URL         string `json:"url"`
		HTMLURL     string `json:"html_url"`
		LabelsURL   string `json:"labels_url"`
		ID          int    `json:"id"`
		Number      int    `json:"number"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Creator     struct {
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
		} `json:"creator"`
		OpenIssues   int         `json:"open_issues"`
		ClosedIssues int         `json:"closed_issues"`
		State        string      `json:"state"`
		CreatedAt    time.Time   `json:"created_at"`
		UpdatedAt    time.Time   `json:"updated_at"`
		DueOn        time.Time   `json:"due_on"`
		ClosedAt     interface{} `json:"closed_at"`
	} `json:"milestone"`
	Comments    int         `json:"comments"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	ClosedAt    interface{} `json:"closed_at"`
	Body        string      `json:"body"`
	PullRequest struct {
		URL      string `json:"url"`
		HTMLURL  string `json:"html_url"`
		DiffURL  string `json:"diff_url"`
		PatchURL string `json:"patch_url"`
	} `json:"pull_request"`
}

var gIssues []GitIssues
var pRequests []GitIssues
var token = ""

func main() {

	token = os.Getenv("GIT_TOKEN")
	fmt.Println(token)
	if token == "" {
		fmt.Println("Git Token is not set")
		os.Exit(101)

	}
	auth()
	http.HandleFunc("/getIssues", getIssues)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.ListenAndServe(":7000", nil)

}

func exitOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func auth() {
	postURL := `https://api.github.com/?access_token=` + token
	_, err := http.Get(postURL)
	exitOnErr(err)

}

func populateIssues(url string) {
	pullURL := url + token
	fmt.Println("Fetching from....", pullURL)
	resp, err := http.Get(pullURL)
	exitOnErr(err)

	defer resp.Body.Close()
	htmlData, err := ioutil.ReadAll(resp.Body)
	exitOnErr(err)

	mIssues := []Repoissue{}
	json.Unmarshal(htmlData, &mIssues)

	//json.Unmarshal(htmlData, &gIssues)
	var flag int

	for _, elem := range mIssues {
		eachGitIssue := GitIssues{}

		if elem.PullRequest.URL != "" {
			flag = 1
		} else {
			eachGitIssue.Number = elem.Number
			eachGitIssue.Title = elem.Title
			eachGitIssue.Repo = elem.RepositoryURL

			for _, labe := range elem.Labels {
				eachGitIssue.Labels += labe.Name + " "
			}
			for _, assignee := range elem.Assignees {
				eachGitIssue.Assignees += assignee.Login + " "
			}
			eachGitIssue.Milestone = elem.Milestone.Title
			eachGitIssue.State = elem.Milestone.State
		}
		if flag != 1 {
			gIssues = append(gIssues, eachGitIssue)
			flag = 0
		} else {

			pRequests = append(pRequests, eachGitIssue)
			flag = 0
		}
	} // end of for

}

func getIssues(w http.ResponseWriter, req *http.Request) {
	gIssues = nil
	pRequests = nil
	populateIssues(`https://api.github.com/repos/minio/minio/issues?state=open&per_page=100&access_token=`)
	populateIssues(`https://api.github.com/repos/minio/mc/issues?state=open&per_page=100&access_token=`)

	populateIssues(`https://api.github.com/repos/minio/minio-go/issues?state=open&per_page=100&access_token=`)
	populateIssues(`https://api.github.com/repos/minio/minio-js/issues?state=open&per_page=100&access_token=`)
	populateIssues(`https://api.github.com/repos/minio/minio-java/issues?state=open&per_page=100&access_token=`)
	populateIssues(`https://api.github.com/repos/minio/minio-py/issues?state=open&per_page=100&access_token=`)

	populateIssues(`https://api.github.com/repos/minio/doctor/issues?state=open&per_page=100&access_token=`)
	populateIssues(`https://api.github.com/repos/minio/bosh-release/issues?state=open&per_page=100&access_token=`)
	populateIssues(`https://api.github.com/repos/minio/minfs/issues?state=open&per_page=100&access_token=`)

	js, err := json.Marshal(gIssues)
	if err != nil {
		fmt.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
