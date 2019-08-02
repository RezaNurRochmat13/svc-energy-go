package energy

import (
	"svc-energy-go/modules/energy/delivery/handler"

	"github.com/gin-gonic/gin"
)

func EnergyRoutes(userRoutes *gin.Engine) {

	energyRouteConfiguration := userRoutes.Group("/public/api/v1")
	{
		energyRouteConfiguration.GET("energy", handler.GetAllenergy)
	}
}
