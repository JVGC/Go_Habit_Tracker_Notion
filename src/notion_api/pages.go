package notionapi

import (
	"encoding/json"
	"fmt"
	pages_models "go_notion_api/src/notion_api/models"
	"io"
	"net/http"
	"os"
)

func GetPages() pages_models.PagesQuery{
	req, err := http.NewRequest("POST", "https://api.notion.com/v1/databases/"+
															os.Getenv("DATABASE_ID")+"/query", nil)

	if err != nil{
		fmt.Print(err.Error())
		os.Exit(1)
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("SECRET_TOKEN"))
	req.Header.Set("Notion-Version", "2022-06-28")

	response, err := http.DefaultClient.Do(req)

	if err != nil{
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil{
		fmt.Print(err.Error())
		os.Exit(1)
	}

	var data pages_models.PagesQuery
	json.Unmarshal(responseData, &data)

	return data

}

func HabitsPercentage() map[string]float64{
	habits := GetHabits()
	pages := GetPages()

	sumObj := make(map[string]float64)

	for _, habit := range habits {
		sumObj[habit] = 0
		counter := 0.0
		for _, page := range pages.Pages{
			counter+=1
			if page.Properties[habit].Checkbox{
				sumObj[habit] +=1
			}
		}
		fmt.Printf("Sum %f %f", (sumObj[habit]), (counter))
		sumObj[habit] = (sumObj[habit] / counter)*100
	}

	return sumObj
}