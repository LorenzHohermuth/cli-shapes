package color

import (
	"fmt"
	"math"
)

var colorForeground = map[[3]int]string {
	[3]int{0, 0, 0}:"\033[30m", //black
	[3]int{204, 0, 0}:"\033[31m", //dark red
	[3]int{25, 203, 0}:"\033[32m", //dark green
	[3]int{206, 203, 0}:"\033[33m", //dark yellow
	[3]int{13, 115, 204}:"\033[34m", //dark blue
	[3]int{203, 20, 209}:"\033[35m", //dark pink
	[3]int{13, 205, 205}:"\033[36m", //dark cyan
	[3]int{221, 221, 221}:"\033[37m", //dark white
	[3]int{118, 118, 118}:"\033[90m", // grey
	[3]int{242, 32, 31}:"\033[91m", //red
	[3]int{35, 253, 0}:"\033[92m", //green
	[3]int{255, 253, 0}:"\033[93m", //yellow
	[3]int{26, 143, 255}:"\033[94m", //blue
	[3]int{253, 40, 255}:"\033[95m", //pink
	[3]int{20, 255, 255}:"\033[96m", // cyan
	[3]int{255, 255, 255}:"\033[97m", //white
}

func PrintColorPallet() {
	non := "\033[0m"
	arr := []string{"\033[40m", "\033[41m", "\033[42m", "\033[43m", "\033[44m", "\033[45m", "\033[46m", "\033[47m", "\033[100m", "\033[101m", "\033[102m", "\033[103m", "\033[104m", "\033[105m", "\033[106m", "\033[107m"}
	for i, v := range arr {
		fmt.Println(i, v, "           ", non)
	}
}

func GetDefaultColor(p [3]int) (string, string) {
	var color string = "\033[39m"
	var delta float64 = 255
	r2  := float64(p[0])
	g2  := float64(p[1])
	b2  := float64(p[2])
	for k, v := range colorForeground {
		r1  := float64(k[0])
		g1  := float64(k[1])
		b1  := float64(k[2])
		otherDelta := math.Sqrt(math.Pow((r2-r1)*0.6, 2) + math.Pow((g2-g1)*1.5, 2) + math.Pow((b2-b1)*0.5,2))
		if (otherDelta < delta) {
			delta = otherDelta
			color = v
		}
	}
	return color, "\033[0m"
}
