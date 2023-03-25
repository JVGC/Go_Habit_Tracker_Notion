package notionapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_notion_api/src/notion_api/models"
	pages_models "go_notion_api/src/notion_api/models"
	"io"
	"net/http"
	"os"
)


func setRequestHeader(req *http.Request){
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SECRET_TOKEN"))
	req.Header.Set("Notion-Version", "2022-06-28")
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
}

func getPagesStruct(responseBody io.ReadCloser) pages_models.PagesQuery{
	responseData, err := io.ReadAll(responseBody)
	if err != nil{
		fmt.Print(err.Error())
		os.Exit(1)
	}

	var data pages_models.PagesQuery
	json.Unmarshal(responseData, &data)
	return data
}

func doPagesRequest(jsonString string) *http.Response{

	fmt.Print(jsonString)
	jsonBody := []byte(jsonString)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest("POST", "https://api.notion.com/v1/databases/"+
															os.Getenv("DATABASE_ID")+"/query", bodyReader)

	if err != nil{
		fmt.Print(err.Error())
		os.Exit(1)
	}

	setRequestHeader(req)

	response, err := http.DefaultClient.Do(req)

	if err != nil{
		fmt.Print(err.Error())
		os.Exit(1)
	}
	return response
}

func GetPages(f models.Filter) []pages_models.Page{

	var has_more bool = true

  requestQuery := models.PagesRequestQuery{
		Page_Size: 10,
	}

	if(f.Created_time.On_or_after != ""){
		dateFilter := models.DateFilter{
			On_or_after: f.Created_time.On_or_after,
		}
		filter := models.Filter{
			Timestamp: "created_time",
			Created_time: dateFilter,
		}
		requestQuery.Filter = &filter
	}
	data := []pages_models.Page{}

	for has_more{

		jsonFilter, _ := json.Marshal(requestQuery)
		response := doPagesRequest(string(jsonFilter))

		pagesQuery := getPagesStruct(response.Body)
		has_more = pagesQuery.Has_More
		fmt.Print(has_more, pagesQuery.Next_Cursor)
		requestQuery.Start_cursor = pagesQuery.Next_Cursor
		data = append(data, pagesQuery.Pages...)
	}

	return data

}