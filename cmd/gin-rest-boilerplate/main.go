package main

import (
	"github.com/pa3ng/gin-boilerplate/internal/app"
)

func main() {
	router := app.NewRouter()

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
