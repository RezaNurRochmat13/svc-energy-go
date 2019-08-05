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
		ctx.JSON(500, gin.H{"error": "Internal server error"})
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
		ctx.JSON(500, gin.H{"error": "Internal server error"})
	} else if meterReadService == nil {
		ctx.JSON(200, gin.H{
			"message": "Data kosong",
			"data":    []int{},
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
		ctx.JSON(500, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(200, gin.H{
		"message": "Meter data created",
		"created": createNewMeterReadPayload,
	})
}

func UpdateMeterRead(ctx *gin.Context) {
	var (
		updateMeterReadPayload model.UpdateMeterRead
		meterReadId            = ctx.Param("meterReadId")
	)

	meterReadService, errorHandlerService := service.FetchDetailMeterRead(meterReadId)

	if errorHandlerService != nil {
		log.Printf("Error when get service : %s", errorHandlerService)
		ctx.JSON(500, gin.H{"error": "Internal server error"})
	}

	if meterReadService == nil {
		ctx.JSON(200, gin.H{
			"message": "Data kosong",
			"data":    []int{},
		})
	} else {
		ctx.ShouldBindJSON(&updateMeterReadPayload)

		UpdateMeterReadProcess(updateMeterReadPayload, meterReadId, ctx)
	}
}

func UpdateMeterReadProcess(meterReadUpdatePayload model.UpdateMeterRead, meterReadId string, ctx *gin.Context) {

	errorHandlerUpdateMeterReadService := service.UpdateMeterRead(meterReadUpdatePayload, meterReadId)

	if errorHandlerUpdateMeterReadService != nil {
		log.Printf("Error when get service : %s", errorHandlerUpdateMeterReadService)
		ctx.JSON(500, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(200, gin.H{
		"message": "Meter data updated",
		"updated": meterReadUpdatePayload,
	})
}
