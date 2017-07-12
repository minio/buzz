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
	"regexp"
	"strconv"
	"strings"
)

type ETAStruct struct {
	ETA, Error string
}

func getETAFromComment(org string, repo string, id int) ETAStruct {

	var comments []string
	var eta ETAStruct
	iComments, _, err := buzzClient.Issues.ListComments(ctx, org, repo, id, nil)
	if err != nil {
		fmt.Printf("Unable to get comment. Error: \"%s\"\n", err)
		if strings.Contains(err.Error(), "404 Not Found") {
			eta = ETAStruct{"", "Issue Not Found"}
		}
	} else {
		for _, comment := range iComments {
			comments = append(comments, comment.GetBody())
		}
		if len(comments) == 0 {
			eta = ETAStruct{"", "No Comments Found"}
		} else {
			re := regexp.MustCompile(`ETA: (20(1[7-9]|[2-9][0-9])-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) (0[1-9]|1[0-9]|2[0-3]):[0-5][0-9])`)
			for i := len(comments) - 1; i >= 0; i-- {
				match := re.FindStringSubmatch(comments[i])
				if len(match) > 0 {
					eta = ETAStruct{match[1], ""}
				}
			}
			if eta.ETA == "" {
				eta = ETAStruct{"", "No ETA Specified"}
			}
		}
	}
	return eta
}

func getETA(w http.ResponseWriter, req *http.Request) {
	org := req.URL.Query().Get("org")
	repo := req.URL.Query().Get("repo")
	id_str := req.URL.Query().Get("id")
	id_int, err := strconv.Atoi(id_str)

	var eta ETAStruct

	if err != nil || id_int == 0 || org == "" || repo == "" {
		fmt.Println("Missing/Wrong argument(s) in your URL!")
		fmt.Println("Expected: ../getETA?org=<org>&repo=<repo>&id=<id>")
		fmt.Println("Received: ../getETA?org=\"" + org + "\"&repo=\"" + repo + "\"&id=\"" + id_str + "\".")
		eta = ETAStruct{"", "Wrong ARGs"}
	} else {
		eta = getETAFromComment(org, repo, id_int)
	}
	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(eta)
	if err != nil {
		fmt.Print(err)
		return
	}
	w.Write(js)
}
