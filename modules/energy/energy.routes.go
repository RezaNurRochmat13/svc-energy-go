package energy

import (
	"svc-energy-go/modules/energy/delivery/handler"

	"github.com/gin-gonic/gin"
)

func EnergyRoutes(userRoutes *gin.Engine) {

	energyRouteConfiguration := userRoutes.Group("/public/api/v1")
	{
		energyRouteConfiguration.GET("meter-reads", handler.GetAllMeterRead)
		energyRouteConfiguration.GET("meter-reads/:meterReadId", handler.GetDetailMeterRead)
		energyRouteConfiguration.POST("meter-reads", handler.CreateNewMeterRead)
	}
}
