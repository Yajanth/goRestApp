package model

type Car struct {
	Model       string       `json:"model"`
	Year        int          `json:"year"`
	Manufacture Manufacturer `json:"Manufacturer"`
}

type Manufacturer struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}
