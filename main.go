package main

import (
	"github.com/gabrielg2020/blog/logger"
	"github.com/gabrielg2020/blog/routes"
)

func main()  {
	router := routes.SetupRouter()

	if err := router.Run(":8080"); err != nil {
		logger.Fatal("Failed to start server: ", err)
	}
}