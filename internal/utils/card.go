package utils

import "strconv"

func LuhnCheck(number string) bool {
	var sum int
	var skip bool

	for i:= len(number) - 1; i > -1; i-- {
		val, _ := strconv.Atoi(string(number[i]))
		if skip {
			val *= 2
			if val > 9 {
				val = (val % 10) + 1
			}
		}

		skip = !skip
		sum += val
	}
	return sum % 10 == 0
}