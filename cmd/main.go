package main

import (
	"github.com/gin-gonic/gin"
	"internal/app/handlers"
)

func main() {
	r := gin.Default()
	accountHandler := handlers.NewAccountHandler()
	accountHandler.RegisterRoutes(r)
	r.Run(":8080")
}