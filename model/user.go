package model

type User struct {
	ID    int    `json:"ID"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
