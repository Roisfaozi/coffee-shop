package config

type Result struct {
	Data    interface{}
	Meta    interface{}
	Message interface{}
}

type Metas struct {
	Next  interface{} `json:"next"`
	Prev  interface{} `json:"prev"`
	Total int         `json:"total"`
}
