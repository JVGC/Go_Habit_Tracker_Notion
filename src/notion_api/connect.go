package notionapi_models

import (
	"encoding/json"
	"fmt"
	database_models "go_notion_api/src/notion_api/database"
	"io"
	"log"
	"net/http"
	"os"
)

func GetDatabaseSettings() database_models.NotionDatabase {

	req, err := http.NewRequest("GET", "https://api.notion.com/v1/databases/"+
															os.Getenv("DATABASE_ID"),nil)

	req.Header.Set("Authorization", "Bearer "+os.Getenv("SECRET_TOKEN"))
	req.Header.Set("Notion-Version", "2022-06-28")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	response, _ := http.DefaultClient.Do(req)

	responseData, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
		}
		// fmt.Println(string(responseData))

		var data database_models.NotionDatabase
		json.Unmarshal(responseData, &data)
		return data
}

func GetHabits() []string {
	database := GetDatabaseSettings()
	habits := make([]string, 0, len(database.Properties))
	for k := range database.Properties {
		if database.Properties[k].Type == "checkbox"{
			habits = append(habits, k)
		}
	}
	return habits
}