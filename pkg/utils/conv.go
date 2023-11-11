package utils

import "strconv"

func ToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
