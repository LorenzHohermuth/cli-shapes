package color

import "fmt"

//printf "\x1b[38;2;255;100;0mTRUECOLOR\x1b[0m\n"

func GetTrueColor(p [3]int) (string, string) {
	color := fmt.Sprintf("\x1b[38;2;%d;%d;%dm", p[0], p[1], p[2])
	return color, "\x1b[0m"
}

