package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

func main() {
	// r := router.SetupRouter()
	// r.Run(":8095")
	var book Book = Book{
		Id:     1,
		Title:  "Atomic Habits",
		Author: "Cal Newport",
		Price:  20.5,
	}

	json_data, err := json.Marshal(book)
	fmt.Print(err)
	fmt.Printf("----------------------------")
	fmt.Println(string(json_data))

}
