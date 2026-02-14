package api

import (
	"logAnalyzer/api/routers"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		logs := v1.Group("/logs")
		routers.HealthRouter(logs)
	}

	r.Run(":8000")
}
