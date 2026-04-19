package validations

import (
	"log"
	"logAnalyzer/pkg"

	"github.com/go-playground/validator/v10"
)

func ServiceValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}
	services := pkg.GetSlice("services")
	if services == nil {
		log.Println("Error while getting services list")
		return false
	}
	for _, service := range services {
		if service == value {
			return true
		}
	}
	return false
}

func LevelValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}
	levels := pkg.GetSlice("levels")
	if levels == nil {
		log.Println("Error while getting levels list")
		return false
	}
	for _, level := range levels {
		if level == value {
			return true
		}
	}
	return false
}
