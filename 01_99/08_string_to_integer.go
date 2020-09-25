package _01_99

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.Trim(str, " ")
	ret := 0
	factor := 1
	if len(str) == 0 {
		return 0
	}
	if str[0] == '-' {
		factor = -factor
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}
	for _, v := range str[0:] {
		if !(v-'0' >= 0 && v-'0' <= 9) {
			break
		}
		ret = ret*10 + int(v-'0')
	}
	if factor*ret >= math.MaxInt32 {
		return math.MaxInt32
	}
	if factor*ret <= math.MinInt32 {
		return math.MinInt32
	}
	return factor * ret
}
