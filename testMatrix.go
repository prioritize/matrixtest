package main

import (
	"matrix"
)

func main() {
	var dims matrix.Dims
	dims.SetDims(3, 3)
	mat := matrix.New(dims)
	newRow := []int{1, 1, 2}
	mat.SetRow(0, newRow)
	mat.PrintSpecific(2, 2)
	mat.PrettyPrint()
}
