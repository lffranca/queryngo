package util

func InterfaceToArrayString(arg interface{}) (result []string) {
	if values, ok := arg.([]string); ok {
		return values
	}

	values, ok := arg.([]interface{})
	if !ok {
		return
	}

	for _, value1 := range values {
		if value2, ok := value1.(string); ok {
			result = append(result, value2)
		}
	}

	return
}
