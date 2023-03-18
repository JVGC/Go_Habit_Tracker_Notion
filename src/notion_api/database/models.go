package models

type Property struct {
	Id string
	Type string
}

type Properties struct {
	ReviewAnki Property `json:"Revisar Anki"`
	Meditation Property `json:"Meditação"`
	Reading Property `json:"Leitura"`
	Date Property
}

type ExternalObj struct {
	Url string
}

type NotionDatabase struct{
	Object string
	Description []string
	Properties Properties
}
