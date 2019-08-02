package handler

import (
	"github.com/gin-gonic/gin"
)

func GetAllenergy(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "All data energy",
	})
}
