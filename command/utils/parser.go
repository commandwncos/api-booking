package utils

import "strconv"

func ParseToInteger(param string) (int64, error) {
	return strconv.ParseInt(param, 10, 64)

}
