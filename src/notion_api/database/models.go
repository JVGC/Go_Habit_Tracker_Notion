package models

type Property struct {
	Id string
	Name string
	Type string
}
type ExternalObj struct {
	Url string
}

type NotionDatabase struct{
	Object string
	Description []string
	Properties map[string]Property `json:"properties"`
}
