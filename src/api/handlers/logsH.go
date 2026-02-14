package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"logAnalyzer/models"
	"logAnalyzer/services"
)

type LogsHandler struct {
}

func NewLogsHandler() *LogsHandler {
	return &LogsHandler{}
}

func (h *LogsHandler) GetLogs(c *gin.Context) {
	lg := models.Log{}
	err := c.ShouldBindHeader(&lg)
	if err != nil {
		log.Printf("Error binding header: %v", err)
	}
	services.RegisterLog(&lg)
}
