package lib

import (
	"fmt"
	"os"
)

// WriteNetPBM dumps a matrix of bits to a file
func WriteNetPBM(
	name string, grid ComplexGrid,
) (int, error) {
	n := 0
	outfile, err := os.Create(name)
	if err != nil {
		return n, err
	}
	defer outfile.Close()

	h, err := fmt.Fprintf(
		outfile, "P4\n%d %d\n",
		grid.cols,
		grid.rows,
	)
	n += h
	if err != nil {
		return n, err
	}
	b, err := outfile.Write(grid.data)
	n += b
	if err != nil {
		return n, err
	}

	return n, nil
}
