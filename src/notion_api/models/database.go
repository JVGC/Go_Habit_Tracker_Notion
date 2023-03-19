package models


type Property struct {
	Id string
	Name string `json:",omitempty"`
	Type string
	Checkbox bool
}
type ExternalObj struct {
	Url string
}

type NotionDatabase struct{
	Object string
	Description []string
	Properties map[string]Property `json:"properties"`
}
