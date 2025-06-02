package main

import (
	"example.com/RestAPI/db"
	"example.com/RestAPI/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
	//err :=
	//if err != nil {
	//	panic(err)
	//}
}
