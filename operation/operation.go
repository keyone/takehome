package operation

import (
	"fmt"
	"strconv"
	"strings"
)

// Operation type takes a two dimensions slice of strings and returns a string with an error.
type Operation func([][]string) (string, error)

// Echo returns the matrix as a string in matrix format.
func Echo(matrix [][]string) (string, error) {
	if checkErr := checkMatrix(matrix); checkErr != nil {
		return "", checkErr
	}
	return printMatrix(matrix), nil
}

// Invert returns the matrix as a string where the columns and rows are inverted.
func Invert(matrix [][]string) (string, error) {
	if checkErr := checkMatrix(matrix); checkErr != nil {
		return "", checkErr
	}
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == matrix[j][i] {
				break
			} else {
				temp := matrix[i][j]
				matrix[i][j] = matrix[j][i]
				matrix[j][i] = temp
			}
		}
	}
	return printMatrix(matrix), nil
}

// Flatten returns the matrix as a 1 line string.
func Flatten(matrix [][]string) (string, error) {
	if checkErr := checkMatrix(matrix); checkErr != nil {
		return "", checkErr
	}
	slice := make([]string, 0)
	for _, row := range matrix {
		for _, val := range row {
			slice = append(slice, val)
		}
	}
	return strings.Join(slice, ","), nil
}

// Sum returns the sum of the integers in the matrix.
func Sum(matrix [][]string) (string, error) {
	if checkErr := checkMatrix(matrix); checkErr != nil {
		return "", checkErr
	}
	sum := 0
	for _, row := range matrix {
		for _, val := range row {
			i, _ := strconv.Atoi(val)
			additionResult, errAddition := checkAdditionOverflow(sum, i)
			if errAddition != nil {
				return "", errAdditionOverflow
			}
			sum = additionResult
		}
	}
	return strconv.Itoa(sum), nil
}

// Multiply returns the product of the integers in the matrix.
func Multiply(matrix [][]string) (string, error) {
	if checkErr := checkMatrix(matrix); checkErr != nil {
		return "", checkErr
	}
	multiply := 1
	for _, row := range matrix {
		for _, val := range row {
			i, _ := strconv.Atoi(val)
			multiplyResult, errMultiply := checkMultiplyOverflow(multiply, i)
			if errMultiply != nil {
				return "", errMultiply
			}
			multiply = multiplyResult
		}
	}
	return strconv.Itoa(multiply), nil
}

// printMatrix returns a string representation of the matrix.
func printMatrix(records [][]string) string {
	var matrix string
	for _, row := range records {
		matrix = fmt.Sprintf("%s%s\n", matrix, strings.Join(row, ","))
	}
	return matrix
}
