package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	notionapi "go_notion_api/src/notion_api"
	"go_notion_api/src/notion_api/models"
	usecases "go_notion_api/src/use_cases"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func loadEnvFile() {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }
}

func NotionRoute(w http.ResponseWriter, r *http.Request) {
	res := notionapi.GetDatabaseSettings()
	json.NewEncoder(w).Encode(res)
}

func HabitsRoute(w http.ResponseWriter, r *http.Request){
	res := notionapi.GetHabits()
	json.NewEncoder(w).Encode(res)
}

func PagesRoute(w http.ResponseWriter, r *http.Request){
	res := notionapi.GetPages(models.Filter{Created_time: models.DateFilter{On_or_after: ""}})
	json.NewEncoder(w).Encode(res)
}

func SumRoute(w http.ResponseWriter, r *http.Request){
	startAt := r.URL.Query().Get("startAt")
	res := usecases.HabitsPercentage(startAt)

	returnObj := make(map[string]string)

	for habit, value := range res{
		returnObj[habit] = fmt.Sprintf("%.2f%%",value*100)
	}
	json.NewEncoder(w).Encode(returnObj)
}


func main(){

	loadEnvFile()
	r := mux.NewRouter()
	r.HandleFunc("/database", NotionRoute)
	r.HandleFunc("/habits", HabitsRoute)
	r.HandleFunc("/habits/sum", SumRoute)
	r.HandleFunc("/pages", PagesRoute)
	log.Fatal(http.ListenAndServe(":3000", r))
}