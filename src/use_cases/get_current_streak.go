package usecases

import (
	"fmt"
	notionapi "go_notion_api/src/notion_api"
	"go_notion_api/src/notion_api/models"
	"time"
)

func GetCurrentStreak(habit string) int {

	sortBy := make([]models.Sort, 0)
	sortBy = append(sortBy, models.Sort{
		Property: "Date",
		Direction: "descending",
	})

	var has_more bool = true
	next_cursor := ""

	current_streak := 0
	var should_continue bool

	for has_more{
		pagesQuery := notionapi.GetPages(models.Filter{
			Created_time: models.DateFilter{
				Before: time.Now().Format("2006-01-02"),
			}},
			next_cursor,
			sortBy...)
		should_continue = getStreak(pagesQuery.Pages, habit, &current_streak)

		if !should_continue{
			break
		}
		has_more = pagesQuery.Has_More
		next_cursor = pagesQuery.Next_Cursor
	}

	return current_streak
}

func getStreak(pages []models.Page, habit string, counter *int) bool{

	var should_continue = true
	for _, page := range pages{
		if page.Properties[habit].Checkbox {
			*counter+=1
		}else{
			should_continue = false
			break
		}
	}
	fmt.Print(should_continue)
	return should_continue

}
