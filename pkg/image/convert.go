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
