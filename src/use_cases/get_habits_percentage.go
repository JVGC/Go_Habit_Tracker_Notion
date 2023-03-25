package usecases

import (
	notionapi "go_notion_api/src/notion_api"
	"go_notion_api/src/notion_api/models"
)

func GetHabitsPercentage(startDate string) map[string]float64{
	var has_more bool = true
	next_cursor := ""

	pages := []models.Page{}
	for has_more{
		pagesResponse := notionapi.GetPages(
			models.Filter{
				Created_time: models.DateFilter{On_or_after: startDate},
			},
			next_cursor,
		)
		pages = append(pages, pagesResponse.Pages...)
		has_more = pagesResponse.Has_More
		next_cursor = pagesResponse.Next_Cursor
	}

	habits := notionapi.GetHabits()
	sumObj := getHabitsSum(pages, habits)

	return sumObj

}

func getHabitsSum(pages []models.Page, habits []string) map[string]float64{
	sumObj := make(map[string]float64)
	for _, habit := range habits {
		sumObj[habit] = 0
		counter := 0.0
		for _, page := range pages{
			counter+=1
			if page.Properties[habit].Checkbox{
				sumObj[habit] +=1
			}
		}
		sumObj[habit] = sumObj[habit] / counter
	}
	return sumObj
}