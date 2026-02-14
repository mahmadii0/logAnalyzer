package routers

import (
	"logAnalyzer/api/handlers"

	"github.com/gin-gonic/gin"
)

func HealthRouter(r *gin.RouterGroup) {
	handler := handlers.NewLogsHandler()

	r.POST("", handler.GetLogs)
}
