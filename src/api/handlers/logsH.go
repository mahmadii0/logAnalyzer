package handlers

import (
	"logAnalyzer/api/helper"
	"logAnalyzer/models"
	"logAnalyzer/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

type LogsHandler struct {
}

func NewLogsHandler() *LogsHandler {
	return &LogsHandler{}
}

func (h *LogsHandler) PostLogs(c *gin.Context) {
	lg := models.Log{}
	err := c.ShouldBindJSON(&lg)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseRwithValidationError("Error binding header",
			false,-1,err))
		return
	}
	err=services.RegisterLog(&lg)
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseRwithValidationError("Error while registering log",
			false,-1,err))
		return
	}
}
