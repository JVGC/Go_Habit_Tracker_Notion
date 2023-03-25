package usecases

import (
	notionapi "go_notion_api/src/notion_api"
	"go_notion_api/src/notion_api/models"
	"time"
)


func GetLongestStreak(habit string) int{

	sortBy := make([]models.Sort, 0)
	sortBy = append(sortBy, models.Sort{
		Property: "Date",
		Direction: "descending",
	})

	var has_more bool = true
	next_cursor := ""

	current_streak := 0
	longest_streak :=0

	pages := []models.Page{}

	for has_more{
		pagesResponse := notionapi.GetPages(models.Filter{
			Created_time: models.DateFilter{
				Before: time.Now().Format("2006-01-02"),
			}},
			next_cursor,
			sortBy...)
		pages = append(pages, pagesResponse.Pages...)
		has_more = pagesResponse.Has_More
		next_cursor = pagesResponse.Next_Cursor
	}

	for _, page := range pages{
		if page.Properties[habit].Checkbox {
			current_streak+=1
		}else{
			if current_streak > longest_streak{
				longest_streak = current_streak
				current_streak = 0
			}
		}
	}

	return longest_streak

}