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



type Sort struct{
	Property string `json:"property,omitempty"`
	Direction string `json:"direction,omitempty"`
}

type DateFilter struct{
	On_or_after string `json:"on_or_after,omitempty"`
	Before string `json:"before,omitempty"`
}

type Filter struct {
	Timestamp string `json:"timestamp,omitempty"`
	Created_time DateFilter `json:"created_time,omitempty"`
}

type PagesRequestQuery struct{
	Filter *Filter `json:"filter,omitempty"`
	Start_cursor string `json:"start_cursor,omitempty"`
	Page_Size int `json:"page_size,omitempty"`
	Sorts []Sort `json:"sorts,omitempty"`
}