package generated

import (
	"math"
	"math/rand/v2"
	"os"
)

func CreatFile(name string) *os.File { 
	fileName := "internal/generated/gen-" +  name + ".go"
	_ = os.Remove(fileName)
	f, err := os.Create(fileName)
	check(err)
	_, err = f.WriteString("package generated\n\nvar " + name + " = map[string]float64{\n")
	check(err)
	return f
}

func AddToFile(line string, file *os.File) {
	_, err := file.WriteString(line + "\n")
	check(err)
}

func CloseFile(f *os.File) {
	_, err := f.WriteString("}")
	check(err)
	f.Close()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func GetChar(brightness float64) string {
	chars := []string{" "}
	var delta float64 = 1
	for k, v := range nerdFont {
		otherDelta := math.Abs(v - brightness)
		if (otherDelta < delta) {
			delta = otherDelta
			chars = []string{k}
		}
		if (otherDelta == delta) {
			chars = append(chars, k)
		}
	}
	return chars[rand.IntN(len(chars))]
}
