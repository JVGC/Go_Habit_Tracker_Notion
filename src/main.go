package main

import (
	"encoding/json"
	"log"
	"net/http"

	notionapi "go_notion_api/src/notion_api"

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
	res := notionapi.GetPages()
	json.NewEncoder(w).Encode(res)
}

func SumRoute(w http.ResponseWriter, r *http.Request){
	res := notionapi.HabitsPercentage()
	json.NewEncoder(w).Encode(res)
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