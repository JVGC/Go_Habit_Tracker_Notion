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
