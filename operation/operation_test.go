package operation

import (
	"testing"
)

func TestEcho(t *testing.T) {
	tt := []struct {
		name           string
		matrix         [][]string
		expectedResult string
		expectedError  error
	}{
		{"odd matrix", [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}, "1,2,3\n4,5,6\n7,8,9\n", nil},
		{"one element", [][]string{{"1"}}, "1\n", nil},
		{"all letters", [][]string{{"a", "b"}, {"c", "d"}}, "", errNotAllNumbers{elementValue: "a", position: 0, rowNumber: 0}},
		{"err not all numbers", [][]string{{"1", "2"}, {"c", "4"}}, "", errNotAllNumbers{elementValue: "c", position: 0, rowNumber: 1}},
		{"even matrix", [][]string{{"1", "2", "3", "4"}, {"5", "6", "7", "8"}, {"9", "10", "11", "12"}, {"13", "14", "15", "16"}}, "1,2,3,4\n5,6,7,8\n9,10,11,12\n13,14,15,16\n", nil},
		{"not square", [][]string{{"1", "2", "3", "4"}, {"5", "6", "7", "8"}, {"9", "10", "11", "12"}}, "", errNotSquare{3, 4}},
		{"zero element", nil, "", errEmptyMatrix},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			echo, err := Echo(tc.matrix)
			var errString string
			if err != nil {
				errString = err.Error()
			}
			if err != tc.expectedError {
				t.Fatalf("echo ran into problem %s", errString)
			}
			if echo != tc.expectedResult {
				t.Fatalf("echo should return %s but got %s", tc.expectedResult, echo)
			}
		})
	}
}

func TestInvert(t *testing.T) {
	tt := []struct {
		name           string
		matrix         [][]string
		expectedResult string
		expectedError  error
	}{
		{"odd matrix", [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}, "1,4,7\n2,5,8\n3,6,9\n", nil},
		{"one element", [][]string{{"1"}}, "1\n", nil},
		{"all letters", [][]string{{"a", "b"}, {"c", "d"}}, "", errNotAllNumbers{elementValue: "a", position: 0, rowNumber: 0}},
		{"err not all numbers", [][]string{{"1", "2"}, {"c", "4"}}, "", errNotAllNumbers{elementValue: "c", position: 0, rowNumber: 1}},
		{"even matrix", [][]string{{"1", "2", "3", "4"}, {"5", "6", "7", "8"}, {"9", "10", "11", "12"}, {"13", "14", "15", "16"}}, "1,5,9,13\n2,6,10,14\n3,7,11,15\n4,8,12,16\n", nil},
		{"not square", [][]string{{"1", "2", "3", "4"}, {"5", "6", "7", "8"}, {"9", "10", "11", "12"}}, "", errNotSquare{3, 4}},
		{"zero element", nil, "", errEmptyMatrix},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			invert, err := Invert(tc.matrix)
			if err != tc.expectedError {
				t.Fatalf("invert ran into problem %s", err.Error())
			}
			if invert != tc.expectedResult {
				t.Fatalf("invert should return %s but got %s", tc.expectedResult, invert)
			}
		})
	}
}

func TestFlatten(t *testing.T) {
	tt := []struct {
		name           string
		matrix         [][]string
		expectedResult string
		expectedError  error
	}{
		{"odd matrix", [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}, "1,2,3,4,5,6,7,8,9", nil},
		{"one element", [][]string{{"1"}}, "1", nil},
		{"all letters", [][]string{{"a", "b"}, {"c", "d"}}, "", errNotAllNumbers{elementValue: "a", position: 0, rowNumber: 0}},
		{"err not all numbers", [][]string{{"1", "2"}, {"c", "4"}}, "", errNotAllNumbers{elementValue: "c", position: 0, rowNumber: 1}},
		{"even matrix", [][]string{{"1", "2", "3", "4"}, {"5", "6", "7", "8"}, {"9", "10", "11", "12"}, {"13", "14", "15", "16"}}, "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16", nil},
		{"not square", [][]string{{"1", "2", "3", "4"}, {"5", "6", "7", "8"}, {"9", "10", "11", "12"}}, "", errNotSquare{3, 4}},
		{"zero element", nil, "", errEmptyMatrix},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			flatten, err := Flatten(tc.matrix)
			if err != tc.expectedError {
				t.Fatalf("flatten ran into problem %s", err.Error())
			}
			if flatten != tc.expectedResult {
				t.Fatalf("flatten should return %s but got %s", tc.expectedResult, flatten)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tt := []struct {
		name           string
		matrix         [][]string
		expectedResult string
		expectedError  error
	}{
		{"odd matrix", [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}, "45", nil},
		{"one element", [][]string{{"1"}}, "1", nil},
		{"all letters", [][]string{{"a", "b"}, {"c", "d"}}, "", errNotAllNumbers{elementValue: "a", position: 0, rowNumber: 0}},
		{"err not all numbers", [][]string{{"1", "2"}, {"c", "4"}}, "", errNotAllNumbers{elementValue: "c", position: 0, rowNumber: 1}},
		{"even matrix", [][]string{{"1", "2", "3", "4"}, {"5", "6", "7", "8"}, {"9", "10", "11", "12"}, {"13", "14", "15", "16"}}, "136", nil},
		{"not square", [][]string{{"1", "2", "3", "4"}, {"5", "6", "7", "8"}, {"9", "10", "11", "12"}}, "", errNotSquare{3, 4}},
		{"zero element", nil, "", errEmptyMatrix},
		{"addition overflow", [][]string{{"9223372036854775802", "2"}, {"2", "3"}}, "", errAdditionOverflow},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sum, err := Sum(tc.matrix)
			if err != tc.expectedError {
				t.Fatalf("sum ran into problem %s", err.Error())
			}
			if sum != tc.expectedResult {
				t.Fatalf("sum should return %s but got %s", tc.expectedResult, sum)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tt := []struct {
		name           string
		matrix         [][]string
		expectedResult string
		expectedError  error
	}{
		{"odd matrix", [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}, "362880", nil},
		{"one element", [][]string{{"1"}}, "1", nil},
		{"all letters", [][]string{{"a", "b"}, {"c", "d"}}, "", errNotAllNumbers{elementValue: "a", position: 0, rowNumber: 0}},
		{"err not all numbers", [][]string{{"1", "2"}, {"c", "4"}}, "", errNotAllNumbers{elementValue: "c", position: 0, rowNumber: 1}},
		{"even matrix", [][]string{{"1", "2", "3", "4"}, {"5", "6", "7", "8"}, {"9", "10", "11", "12"}, {"13", "14", "15", "16"}}, "20922789888000", nil},
		{"large numbers matrix", [][]string{{"123", "2221", "353", "4335"}, {"535", "67465", "745674567", "84765"}, {"9456", "105", "11667", "12567"}, {"15673", "15674", "154567", "164560"}}, "", errMultiplicationOverflow},
		{"not square", [][]string{{"1", "2", "3", "4"}, {"5", "6", "7", "8"}, {"9", "10", "11", "12"}}, "", errNotSquare{3, 4}},
		{"zero element", nil, "", errEmptyMatrix},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			multiply, err := Multiply(tc.matrix)
			if err != tc.expectedError {
				t.Fatalf("multiply ran into problem: %s", err.Error())
			}
			if multiply != tc.expectedResult {
				t.Fatalf("multiply should return %s but got %s", tc.expectedResult, multiply)
			}
		})
	}
}
