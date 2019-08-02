package main

import (
	"io"
	"log"
	"os"

	energyModulesConfiguration "svc-energy-go/modules/energy"

	"github.com/gin-gonic/gin"
)

func main() {
	SetupRootRoutesApplication()

}

func SetupRootRoutesApplication() {
	createLogFiles, errorCreate := os.Create("log/application.log")

	if errorCreate != nil {
		log.Fatal(errorCreate.Error())
	}

	gin.DefaultWriter = io.MultiWriter(createLogFiles, os.Stdout)

	rootRoutesApplicationConfiguration := gin.Default()

	energyModulesConfiguration.EnergyRoutes(rootRoutesApplicationConfiguration)

	rootRoutesApplicationConfiguration.Run(":8080")
}
