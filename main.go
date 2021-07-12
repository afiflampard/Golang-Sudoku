package main

import (
	"log"
	"testidn/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	routes.Routes(r)
	log.Fatal(r.Run(":8000"))
}
