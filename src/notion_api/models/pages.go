package models


type Page struct {
	Id string
	Properties  map[string]Property `json:"properties"`
}

type PagesQuery struct{
	Has_More bool
	Next_Cursor string
	Pages []Page `json:"results"`
}



type DateFilter struct{
	On_or_after string `json:"on_or_after"`
}

type Filter struct {
	Timestamp string `json:"timestamp"`
	Created_time DateFilter `json:"created_time"`
}

type PagesRequestQuery struct{
	Filter Filter `json:"filter,omitempty"`
	Start_cursor string `json:"start_cursor,omitempty"`
	Page_Size int `json:"page_size,omitempty"`
}