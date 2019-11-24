package util

import (
	"strconv"
	"strings"
)

//ax,ay,id
func split(str string) (int, int, int, bool) {
	ss := strings.Split(str, ",")
	if len(ss) != 3 {
		return 0, 0, 0, false
	}
	id, err := strconv.Atoi(ss[2])
	if err != nil {
		return 0, 0, 0, false
	}
	x, y := charToInt(ss[0], ss[1])
	return id, x, y, true
}

func intToChar(i int) rune {
	return rune('a' + i)
}

func charToInt(ax, ay string) (int, int) {
	x := 100
	y := 100
	switch ax {
	case "a":
		x = 0
	case "b":
		x = 1
	case "c":
		x = 2
	}
	switch ay {
	case "A":
		y = 0
	case "B":
		y = 1
	case "C":
		y = 2
	}
	return x, y
}
