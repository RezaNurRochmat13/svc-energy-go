package energy

import (
	"svc-energy-go/modules/energy/delivery/handler"

	"github.com/gin-gonic/gin"
)

func EnergyRoutes(userRoutes *gin.Engine) {

	energyRouteConfiguration := userRoutes.Group("/public/api/v1")
	{
		energyRouteConfiguration.GET("meter-reads", handler.GetAllMeterReadHandler)
		energyRouteConfiguration.GET("meter-reads/:meterReadId", handler.GetDetailMeterReadHandler)
		energyRouteConfiguration.POST("meter-reads", handler.CreateNewMeterReadHandler)
		energyRouteConfiguration.PUT("meter-reads/:meterReadId", handler.UpdateMeterReadHandler)
		energyRouteConfiguration.DELETE("meter-reads/:meterReadId", handler.DeleteMeterReadHandler)
	}
}
