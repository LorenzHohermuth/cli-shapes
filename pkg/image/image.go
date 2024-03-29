package image

import (
	"image"
	"image/png"
	"io"
	"os"
)

func ImageToPixles(path string) [][]Pixel {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	file, err := os.Open(path)
	check(err)
	defer file.Close()
	mat := getPixles(file)
	return squashMatrix(mat)
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

func GetBrightness(img [][]Pixel) float64 {
	amountOfPixels := 0
	avgR, avgG, avgB, avgA := 0, 0, 0, 0
	for _, y := range img {
		for _, x := range y {
			avgR += x.R
			avgG += x.G
			avgB += x.B
			avgA += x.A
			amountOfPixels++
		}
	}

	return Pixel{avgR / amountOfPixels,
		avgG / amountOfPixels,
		avgB / amountOfPixels,
		avgA / amountOfPixels}.Brightness()
}
