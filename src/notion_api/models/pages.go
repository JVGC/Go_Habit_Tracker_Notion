package models


type Page struct {
	Id string
	Properties  map[string]Property `json:"properties"`
}

type PagesQuery struct{
	Pages []Page `json:"results"`
}
