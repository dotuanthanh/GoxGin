package pck

import (
	"strings"
)

func IsValidTypeString(str string) bool {

	if strings.Trim(str, "") == "" {
		return false
	}
	return true
}
func ArrayFiler(arr []interface{}, item interface{}) (int, interface{}) {
	if arr == nil {
		return 0, nil
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			return i, arr[i]
		}
	}
	return 0, nil
}

func FormatString(input string, trimSpace bool) string {
	if trimSpace {
		return strings.TrimSpace(input)
	}

	return input
}
