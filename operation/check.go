package operation

import (
	"errors"
	"fmt"
	"strconv"
)

type errNotSquare struct {
	rowTotal int
	colTotal int
}

type errNotAllNumbers struct {
	elementValue string
	position     int
	rowNumber    int
}

var errEmptyMatrix = errors.New("This is an empty matrix")

var errMultiplicationOverflow = errors.New("Multiplication overflow, please consider smaller numbers")

var errAdditionOverflow = errors.New("Addition overflow, please consider smaller numbers")

func (e errNotSquare) Error() string {
	return fmt.Sprintf("Matrix is not square. It has %d rows but %d columns.", e.rowTotal, e.colTotal)
}

func (e errNotAllNumbers) Error() string {
	return fmt.Sprintf("Element '%s' at position %d in row #%d is not a number, please use only numbers.", e.elementValue, e.position, e.rowNumber)
}

// In Golang, multiplication overflow is possible.
func checkMultiplyOverflow(a, b int) (int, error) {
	result := a * b
	if a != 0 && result/a != b {
		return 0, errMultiplicationOverflow
	}
	return result, nil
}

// In Golang, integer overflow is possible. We need to check for that.
func checkAdditionOverflow(a, b int) (int, error) {
	result := a + b
	if (result > a) == (b > 0) {
		return result, nil // ok
	}
	return 0, errAdditionOverflow

}

// checkMatrix function verifies if the matrix is a square and has only integers.
func checkMatrix(matrix [][]string) error {
	// a nil slice is functionally equivalent to a zero-length slice.
	if matrix == nil {
		return errEmptyMatrix
	}

	noRows := len(matrix)
	noCols := len(matrix[0])

	if noRows != noCols {
		return errNotSquare{rowTotal: noRows, colTotal: noCols}
	}

	for i := range matrix {
		for j, ele := range matrix[i] {
			_, converr := strconv.Atoi(ele)
			if converr != nil {
				return errNotAllNumbers{elementValue: ele, position: j, rowNumber: i}
			}
		}
	}

	return nil
}
