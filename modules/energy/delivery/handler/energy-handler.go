package handler

import (
	"log"
	"svc-energy-go/modules/energy/model"
	"svc-energy-go/modules/energy/service"

	"github.com/gin-gonic/gin"
)

func GetAllMeterRead(ctx *gin.Context) {
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")

	meterReadsService, errorHandlerService := service.FetchAllMeterReads(limit, offset)

	if errorHandlerService != nil {
		log.Printf("Error when get service : %s", errorHandlerService)
	}

	ctx.JSON(200, gin.H{
		"total": len(meterReadsService),
		"count": len(meterReadsService),
		"data":  meterReadsService,
	})
}

func GetDetailMeterRead(ctx *gin.Context) {
	meterReadId := ctx.Param("meterReadId")

	meterReadService, errorHandlerService := service.FetchDetailMeterRead(meterReadId)

	if errorHandlerService != nil {
		log.Printf("Error when get service : %s", errorHandlerService)
	} else if meterReadService == nil {
		ctx.JSON(200, gin.H{
			"data": []int{},
		})
	} else {
		ctx.JSON(200, gin.H{
			"data": meterReadService,
		})
	}
}

func CreateNewMeterRead(ctx *gin.Context) {
	var createNewMeterReadPayload model.CreateNewMeterRead

	ctx.ShouldBindJSON(&createNewMeterReadPayload)

	errorHandlerCreateMeterRead := service.CreateNewMeterRead(createNewMeterReadPayload)

	if errorHandlerCreateMeterRead != nil {
		log.Printf("Error when get service : %s", errorHandlerCreateMeterRead)
	}

	ctx.JSON(200, gin.H{
		"message": "Meter data created",
		"created": createNewMeterReadPayload,
	})
}
