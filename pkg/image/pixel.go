package image

//func getBighnessMap(map [][]Pixel) [][]float64{}

func (pix Pixel) Brightness() float64 {
	R := float64(pix.R)
	G := float64(pix.G)
	B := float64(pix.B)
	A := float64(pix.A)
	brightness := (0.2126*R + 0.7152*G + 0.0722*B) /255 
	return brightness * (A / 255)
}

type Pixel struct {
	R int
	G int
	B int
	A int
}

func averagePixel(arr []Pixel) Pixel {
	l := len(arr)
	var sumR, sumG, sumB, sumA int
	for _, v := range arr {
		sumR += v.R
		sumG += v.G
		sumB += v.B
		sumA += v.A
	}

	return Pixel{sumR / l, sumG / l, sumB / l, sumA / l}
}
