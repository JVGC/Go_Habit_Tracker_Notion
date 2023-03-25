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

func GetPages(f models.Filter, start_cursor string, s  ...models.Sort) models.PagesQuery{

  requestQuery := models.PagesRequestQuery{
		Page_Size: 10,
	}
	dateFilter := models.DateFilter{
		On_or_after: f.Created_time.On_or_after,
		Before: f.Created_time.Before,
	}
	filter := models.Filter{
		Timestamp: "created_time",
		Created_time: dateFilter,
	}

	if(f.Created_time.On_or_after != ""){
		requestQuery.Filter = &filter
	}
	if(f.Created_time.Before != ""){
		requestQuery.Filter = &filter
	}
	if(len(s) > 0){
		requestQuery.Sorts = s
	}
	if start_cursor != ""{
		requestQuery.Start_cursor = start_cursor
	}
	data := []pages_models.Page{}

	jsonFilter, _ := json.Marshal(requestQuery)
	response := doPagesRequest(string(jsonFilter))

	pagesQuery := getPagesStruct(response.Body)
	data = append(data, pagesQuery.Pages...)

	return models.PagesQuery{
		Pages: data,
		Next_Cursor: pagesQuery.Next_Cursor,
		Has_More: pagesQuery.Has_More,
	}

}