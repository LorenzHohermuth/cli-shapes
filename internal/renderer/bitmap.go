package renderer

type Bitmap struct {
	Matrix [][]float64 //number between 0 and 1
}

func (b Bitmap) Render() (mat [][]string ) {
	for _, v := range b.Matrix {
		var arr []string
		for _, x := range v {
			arr = append(arr, GetChar(x))
		}
		mat = append(mat, arr)
	}
	return mat
}
