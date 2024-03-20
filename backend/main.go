package main

import (
	"fmt"
	"gin-example/db"
	"gin-example/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const PORT string = "8080"

func init() {
	db.Init()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.Use(cors.Default())
	routes.RegisterRoutes(server)

	fmt.Println("Server running on port : ", PORT)
	log.Fatal(server.Run(":" + PORT))
}
