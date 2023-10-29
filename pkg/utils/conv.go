package utils

import "strconv"

func ToInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return i, err
}
