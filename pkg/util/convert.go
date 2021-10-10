package util

import "encoding/json"

func InterfaceToFloat(val interface{}) float64 {
	var final float64

	switch t := val.(type) {
	case int:
		final = float64(t)
	case int8:
		final = float64(t)
	case int16:
		final = float64(t)
	case int32:
		final = float64(t)
	case int64:
		final = float64(t)
	// case bool:
	case float32:
		final = float64(t)
	case float64:
		final = t
	case uint8:
		final = float64(t)
	case uint16:
		final = float64(t)
	case uint32:
		final = float64(t)
	case uint64:
		final = float64(t)
		// case string:
		// default:
	}

	return final
}

func InterfaceToInt(val interface{}) int {
	var final int

	switch t := val.(type) {
	case int:
		final = t
	case int8:
		final = int(t)
	case int16:
		final = int(t)
	case int32:
		final = int(t)
	case int64:
		final = int(t)
	// case bool:
	case float32:
		final = int(t)
	case float64:
		final = int(t)
	case uint8:
		final = int(t)
	case uint16:
		final = int(t)
	case uint32:
		final = int(t)
	case uint64:
		final = int(t)
		// case string:
		// default:
	}

	return final
}

func ToJSON(data interface{}) string {
	jsonResult, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(jsonResult)
}
