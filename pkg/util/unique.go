package util

func UniqueString(arg interface{}) (result []string) {
	values := InterfaceToArrayString(arg)
	if values == nil {
		return
	}

	resultMap := map[string]struct{}{}
	for _, value := range values {
		resultMap[value] = struct{}{}
	}

	for key := range resultMap {
		result = append(result, key)
	}

	return
}
