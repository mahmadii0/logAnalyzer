package api

import (
	"log"
	"logAnalyzer/api/routers"
	"logAnalyzer/api/validations"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer() {
	r := gin.Default()
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// registering validation for nontoneof
		val.RegisterValidation("service", validations.ServiceValidator)
		val.RegisterValidation("level", validations.LevelValidator)
	}
	v1 := r.Group("/api/v1")
	{
		logs := v1.Group("/logs")
		routers.HealthRouter(logs)
	}
	log.Println("[Log Analyzer Started!]")
	r.Run(":8000")
}
