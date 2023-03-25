package notionapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_notion_api/src/notion_api/models"
	"go_notion_api/src/services"
	"io"
	"net/http"
	"os"
)

func getPagesStruct(responseBody io.ReadCloser) models.GetPagesResponse{
	responseData, err := io.ReadAll(responseBody)
	if err != nil{
		fmt.Print(err.Error())
		os.Exit(1)
	}

	var data models.GetPagesResponse
	json.Unmarshal(responseData, &data)
	return data
}

func doPagesRequest(jsonString string) *http.Response{

	fmt.Print(jsonString)
	jsonBody := []byte(jsonString)
	bodyReader := bytes.NewReader(jsonBody)

	notionClient := services.NotionClient{}
	notionClient.Init()

	response := notionClient.Post("/databases/"+os.Getenv("DATABASE_ID")+"/query", bodyReader)

	return response
}

func GetPages(f models.Filter, start_cursor string, s  ...models.Sort) models.GetPagesResponse{

  requestQuery := models.GetPagesRequest{
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
	data := []models.Page{}

	jsonFilter, _ := json.Marshal(requestQuery)
	response := doPagesRequest(string(jsonFilter))

	pagesQuery := getPagesStruct(response.Body)
	data = append(data, pagesQuery.Pages...)

	return models.GetPagesResponse{
		Pages: data,
		Next_Cursor: pagesQuery.Next_Cursor,
		Has_More: pagesQuery.Has_More,
	}

}