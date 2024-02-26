package image

var heigthCell int = 2
var widthCell int = 1

func reFormatVertical(matrix [][]Pixel ) [][]Pixel {
	heightMatrix := len(matrix)
	widthMatrix := len(matrix[0])
	newMatrix := makePixleMatrix(widthMatrix, heightMatrix)
	for y := range matrix {
		for x := range matrix[y] {
			newMatrix[x][y] = matrix[y][x]
		}
	}
	return newMatrix
}

func splitArray(arr []Pixel, size int) [][]Pixel {
	var out [][]Pixel
	var j int 
	for i := 0; i < len(arr); i += size {
		j += size
		if j > len(arr) {
			j = len(arr)
		}
		out = append(out, arr[i:j])
	}
	return out
}

func squashMatrix(pix [][]Pixel) [][]Pixel{
	pix = reFormatVertical(pix)
	var tensor [][][]Pixel
	for _, arr := range pix {
		tensor = append(tensor, splitArray(arr, heigthCell))
	}

	out := makePixleMatrix(len(tensor[0]), len(tensor))
	
	for x := range tensor {
		for y := range tensor[x] {
			p := AveragePixel(tensor[x][y])
			out[y][x] = p
		}
	}

	return out
}


type cord struct {
	x int
	y int
}

func makePixleMatrix(height int, width int) [][]Pixel {
	arr := make([][]Pixel, height)
	for i := range arr {
		arr[i] = make([]Pixel, width)
	}
	return arr
}
