package lib

import (
	"runtime"
	"sync"
)

func ComputeMandelbrotGrid(
	grid ComplexGrid,
) (ComplexGrid, error) {
	poolsize := runtime.NumCPU()
	row_channel := make(chan ComplexGridRow, poolsize)

	var wg sync.WaitGroup
	for core := 0; core < poolsize; core++ {
		wg.Add(1)
		go computeMandelbrotRow(row_channel, &wg)
	}

	for n := 0; n < grid.rows; n++ {
		offset := n * grid.row_bytes
		row_channel <- ComplexGridRow{
			data:    grid.data[offset : offset+grid.row_bytes],
			delta:   grid.delta,
			c0_real: grid.c0_real,
			c0_imag: grid.c0_imag + float64(n)*grid.delta,
			cols:    grid.cols,
		}
	}
	close(row_channel)

	wg.Wait()

	return grid, nil
}

func computeMandelbrotRow(row_channel chan ComplexGridRow, wg *sync.WaitGroup) {
	defer wg.Done()
	for row := range row_channel {
		for m := 0; m < row.cols; m++ {
			c0_real := row.c0_real + float64(m)*row.delta

			bytepos := m >> 3
			bitpos := 7 - m%8

			bit := escapes(c0_real, row.c0_imag)

			row.data[bytepos] |= (bit << bitpos)
		}
	}
}

// escapes returns 1 if a point c with coordinates cr + i * ci is part of the
// Mandelbrot set, or 0 when it escapes the set.
func escapes(cr, ci float64) uint8 {
	zr := cr
	zi := ci
	for i := 0; i < 50; i++ {
		tr := zr*zr - zi*zi + cr
		ti := 2*zr*zi + ci
		if tr*tr+ti*ti > 4.0 {
			return 0
		}
		zr = tr
		zi = ti
	}
	return 1
}
