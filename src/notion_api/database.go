package notionapi

import (
	"encoding/json"
	database_models "go_notion_api/src/notion_api/models"
	"go_notion_api/src/services"
	"io"
	"log"
	"os"
)

func GetDatabaseSettings() database_models.NotionDatabase {

	notionClient := services.NotionClient{}
	notionClient.Init()

	response := notionClient.Get("/databases/"+ os.Getenv("DATABASE_ID"))

	responseData, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
		}
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