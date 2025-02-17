package helper

import "strconv"

func Currency(value int) string {
	return "Rp " + strconv.Itoa(value)
}