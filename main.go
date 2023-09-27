package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	handlers "shipment-service/handlers"
)

func main() {
	fmt.Print("Welcome to Shipment Service!")
	r := gin.Default()

	r.POST("/package/calculate", handlers.Calculate)
	r.POST("/package/add", handlers.AddPackageSize)
	r.POST("/package/remove", handlers.RemovePackageSize)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
