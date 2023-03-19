package usecases

import (
	notionapi "go_notion_api/src/notion_api"
)

func HabitsPercentage(date string ) map[string]float64{
	habits := notionapi.GetHabits()
	pages := notionapi.GetPages(date)

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
		sumObj[habit] = (sumObj[habit] / counter)*100
	}

	return sumObj
}