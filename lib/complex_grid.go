package lib

import (
	"fmt"
)

type ComplexGrid struct {
	data      []byte
	rows      int
	cols      int
	row_bytes int
	c0_imag   float64
	c0_real   float64
	delta     float64
}

type ComplexGridRow struct {
	data    []byte
	cols    int
	c0_imag float64
	c0_real float64
	delta   float64
}

func NewComplexGrid(
	res int,
) (ComplexGrid, error) {
	// define complex area
	c0_real := -1.5
	c0_imag := -1.0
	imag_range := 2
	real_range := 2

	delta := 1.0 / float64(res)
	cols := res * real_range
	rows := res * imag_range

	if cols%8 != 0 {
		fmt.Println("resolution needs to be a multiple of 4")
	}
	row_bytes := cols / 8

	var grid = ComplexGrid{
		data:      make([]byte, rows*row_bytes),
		cols:      cols,
		rows:      rows,
		row_bytes: row_bytes,
		delta:     delta,
		c0_real:   c0_real,
		c0_imag:   c0_imag,
	}

	return grid, nil
}
