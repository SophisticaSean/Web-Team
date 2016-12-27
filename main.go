package main

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
	"github.com/levigross/grequests"
)

type issue struct {
	ID              string          `json:"ID"`
	Name            string          `json:"name"`
	ObjCode         string          `json:"objCode"`
	ParameterValues parameterValues `json:"parameterValues"`
}

type issues struct {
	Issues []issue `json:"data"`
}

type parameterValues struct {
	WebTeamField1 string `json:"DE:Web Team Field"`
	WebTeamField2 string `json:"DE:Second Web Team Field"`
	WebTeamField3 string `json:"DE:Last Field for Web Team"`
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
		spew.Dump(issue)
		//ro := &grequests.RequestOptions{}

		//url := `https://chrisvirostko.attask-ondemand.com/attask/api-internal/optask/` + issue.ID + `/convertToTask?apiKey=apiKey=lqmy3lotx574xgujt5dkosqosld8fgzh&updates={"options":["preservePrimaryContact"],"task":{"name":"` + URL.QueryEscape(issue.Name) + `"},"categoryID":"5850b48d0004df896e7bd11765f94020","DE:Web Team Field":"` + WebTeamField1 + `"}&method=PUT`
		//fmt.Println(url)
		//resp, err := grequests.Put(url, ro)
		//if err != nil {
		//panic(err)
		//}

		//spew.Dump(resp.Bytes())
		//spew.Dump(resp.StatusCode)
	}
}