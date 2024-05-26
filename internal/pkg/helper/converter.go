package helper

import (
	"strings"
)

func SplitStringArrayAgg(str string) []string {
	s := strings.Trim(str, "{}")
	return strings.Split(s, ",")
}
