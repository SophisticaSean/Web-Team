package main

import (
	"encoding/json"
	"fmt"

	URL "net/url"

	"github.com/davecgh/go-spew/spew"
	"github.com/levigross/grequests"
)

type issue struct {
	ID              string                 `json:"ID,omitempty"`
	Name            string                 `json:"name"`
	ObjCode         string                 `json:"objCode,omitempty"`
	CatID           string                 `json:"categoryID"`
	ParameterValues map[string]interface{} `json:"parameterValues"`
}

type issues struct {
	Issues []issue `json:"data"`
}

func main() {

	url := "https://chrisvirostko.attask-ondemand.com/attask/api/optask/search?fields=parameterValues&apiKey=lqmy3lotx574xgujt5dkosqosld8fgzh&projectID=5850b4d80004ed55a8094721475a33bc"

	ro := &grequests.RequestOptions{}
	resp, err := grequests.Get(url, ro)
	if err != nil {
		panic(err)
	}

	spew.Printf(string(resp.Bytes()))

	var returnList issues

	if err := json.Unmarshal(resp.Bytes(), &returnList); err != nil {
		panic(err)
	}

	spew.Dump(returnList)

	for _, issue := range returnList.Issues {
		// cache our issue.ID
		id := issue.ID
		// set ID and ObjCode to blank so they're omitted from the JSON
		issue.ID = ""
		issue.ObjCode = ""
		// set our CatID
		issue.CatID = "5850b48d0004df896e7bd11765f94020"
		// Marshal our issue object into a json byte array
		jsonIssueBytes, err := json.Marshal(issue)
		if err != nil {
			panic(err)
		}

		// convert byte array to URL escaped string
		jsonIssue := URL.QueryEscape(string(jsonIssueBytes))

		ro := &grequests.RequestOptions{}

		url := `https://chrisvirostko.attask-ondemand.com/attask/api-internal/optask/` + id + `/convertToTask?apiKey=lqmy3lotx574xgujt5dkosqosld8fgzh&updates={"options":["preservePrimaryContact"],"task":` + jsonIssue + `}&method=PUT`
		fmt.Println(url)
		resp, err := grequests.Put(url, ro)
		if err != nil {
			panic(err)
		}

		spew.Dump(resp.Bytes())
		spew.Dump(resp.StatusCode)
	}
}
