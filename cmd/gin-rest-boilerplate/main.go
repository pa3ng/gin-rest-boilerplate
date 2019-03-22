package main

import (
	r "github.com/pa3ng/gin-rest-boilerplate/pkg/router"
)

func main() {
	router := r.NewRouter()

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
