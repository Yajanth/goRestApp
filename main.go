package main

import (
	"github.com/Yajanth/goRestApp/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8095")
}
