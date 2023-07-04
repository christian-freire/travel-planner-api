package routes

import (
	"net/http"
	"travel-planner-api/lib"

	"github.com/gin-gonic/gin"
)

func AddRequest(r *gin.Engine) {
	r.POST("/trip", func(c *gin.Context) {
		var trip lib.TravelInfo
		if err := c.ShouldBindJSON(&trip); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
			return
		}

		lib.Travel(c, trip)
	})
}
