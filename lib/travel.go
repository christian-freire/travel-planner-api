package lib

import (
	"travel-planner-api/services"

	"github.com/gin-gonic/gin"
)

type TravelInfo struct {
	StarDate        string `json:"start_date"`
	EndDate         string `json:"end_date"`
	CityOrigin      string `json:"city_origin"`
	CityDestination string `json:"city_destination"`
}

func Travel(c *gin.Context, info TravelInfo) {
	Prompt(info)
}

func Prompt(trip TravelInfo) {
	Question := "Vou fazer uma viagem de férias partindo de " + trip.CityOrigin + " para " + trip.CityDestination +
		" do dia " + trip.StarDate + " até o dia " + trip.EndDate + ", crie para mim um roteiro de viajem para cada um desses dias que irei passar em " + trip.CityDestination

	services.OpenAIService(Question)

}
