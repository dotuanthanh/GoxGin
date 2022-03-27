package pck

import (
	"fmt"
	"strings"
	"time"
)

func GetTimeStamp() int64 {
	return time.Now().Unix()
}
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

const (
	CrCode      = "\r"
	LfCode      = "\n"
	TabCode     = "\t"
	EmptyString = ""
)

func GetString(input interface{}, count int) string {
	temp := fmt.Sprintf("        %v", input)

	return temp[len(temp)-count:]
}

func RemoveEndLine(input string) string {
	temp := strings.ReplaceAll(input, CrCode, EmptyString)
	temp = strings.ReplaceAll(temp, LfCode, EmptyString)

	return temp
}

func FormatString(input string, trimSpace bool) string {
	if trimSpace {
		return strings.TrimSpace(input)
	}

	return input
}
