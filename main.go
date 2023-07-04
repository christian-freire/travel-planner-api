package main

import (
	"travel-planner-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routes.AddRequest(r)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080

}
