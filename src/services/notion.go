package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
)



type NotionClient struct{
	baseUrl string
}

func (n *NotionClient) Init(){
	n.baseUrl = "https://api.notion.com/v1"
}

func (n *NotionClient) setupAuthorizationHeader(req *http.Request){
	req.Header.Set("Notion-Version", "2022-06-28")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SECRET_TOKEN"))
}

func (n *NotionClient) setupJSONHeader(req *http.Request){
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
}

func (n *NotionClient) Get(endpoint string)  *http.Response{
	req, err := http.NewRequest("GET", n.baseUrl+endpoint, nil)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	n.setupAuthorizationHeader(req)

	response, _ := http.DefaultClient.Do(req)

	return response
}


func (n *NotionClient) Post(endpoint string, body io.Reader)  *http.Response{
	req, err := http.NewRequest("POST", n.baseUrl+endpoint, body)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	n.setupAuthorizationHeader(req)
	n.setupJSONHeader(req)

	response, _ := http.DefaultClient.Do(req)

	return response
}
