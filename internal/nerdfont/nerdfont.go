package nerdfont

import (
	"os"
	"strings"
)

func ParseFile(pathToFile string) []string {
	dat, err := os.ReadFile(pathToFile)
	check(err)
	text := strings.Replace(string(dat), "\n", "", -1)
	lines := strings.Split(text, ",")
	return lines
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
