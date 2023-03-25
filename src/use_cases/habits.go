package usecases

import (
	notionapi "go_notion_api/src/notion_api"
	"go_notion_api/src/notion_api/models"
)

func HabitsPercentage(startDate string) map[string]float64{
	habits := notionapi.GetHabits()
	pages := notionapi.GetPages(models.Filter{Created_time: models.DateFilter{On_or_after: startDate}})

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
		sumObj[habit] = (sumObj[habit] / counter)

	}

	return sumObj
}