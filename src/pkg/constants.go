package pkg

import "logAnalyzer/constants"

func GetSlice(name string) []string {
	switch name {

	case "services":
		data := constants.ValidServices
		out := make([]string, len(data))
		copy(out, data)
		return out

	case "levels":
		data := constants.ValidLevels
		out := make([]string, len(data))
		copy(out, data)
		return out
	}

	return nil
}
