package model

type Publisher struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Book struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Price     float64   `json:"price"`
	Publisher Publisher `json:"publisher"`
}
