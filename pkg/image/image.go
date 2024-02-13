package image

import (
	"image"
	"image/png"
	"io"
	"os"
)

type Pixel struct {
	R int
	G int
	B int
	A int
}
func ImageToPixles(path string) [][]Pixel {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	file, err := os.Open(path)
	check(err)
	defer file.Close()
	return  getPixles(file)
}

func getPixles(f io.Reader) [][]Pixel {
	img, _, err := image.Decode(f)
	check(err)
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x,y).RGBA()))
		}
		pixels = append(pixels, row)
	}
	return pixels
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
    return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
