package projutil

import (
	"errors"
	"strconv"
	"strings"
)

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

func ProjNameSplit(name string) (string, int, error) {
	index := strings.LastIndex(name, "_")
	if index < 0 {
		return name, -1, errors.New("no index exists")
	}
	id, err := strconv.Atoi(reverse(name[index+1:]))
	return name[:index], id, err
}

func ProjectNoteName(title string, id int) string {
	return ProjName(title+"_note", id)
}
