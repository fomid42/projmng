package util

import "strconv"

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func ProjName(title string, id int) string {
	return title + "_" + reverse(strconv.Itoa(id))
}
