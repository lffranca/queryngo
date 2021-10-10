package util

func ConcatArrayString(arg1 ...interface{}) (result []string) {
	for _, item := range arg1 {
		result = append(result, InterfaceToArrayString(item)...)
	}

	return
}
