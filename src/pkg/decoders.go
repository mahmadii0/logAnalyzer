package pkg

import (
	"strings"
)

func DecodeGroupKey(key string) map[string]string {
	data := strings.Split(key, "|")
	if len(data) < 3 {
		return nil
	}
	groupKey := map[string]string{
		"service": data[0],
		"level":   data[1],
		"message": data[2],
	}
	return groupKey

}
