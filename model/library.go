package model

type Library struct {
	Name     string            `json:"name"`
	Location string            `json:"location"`
	Books    []Book            `json:"books"`
	Metadata map[string]string `json:"metadata"`
}
