package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"takehome/operation"
)

// Header string on default endpoint and server documentation.
const Header = "Available endpoints are:\n" +
	"curl -F 'file=@matrix.csv' \"localhost:8080/echo\" -- Returns the matrix as a string.\n" +
	"curl -F 'file=@matrix.csv' \"localhost:8080/invert\" -- Returns the matrix as a string where the columns and rows are inverted.\n" +
	"curl -F 'file=@matrix.csv' \"localhost:8080/flatten\" -- Returns the matrix as a 1 line string.\n" +
	"curl -F 'file=@matrix.csv' \"localhost:8080/sum\" -- Returns the sum of the integers in the matrix.\n" +
	"curl -F 'file=@matrix.csv' \"localhost:8080/multiply\" -- Returns the product of the integers in the matrix.\n\n" +
	"The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square).\n" +
	"Each value is an integer, and there is no header row.\n" +
	"\n" +
	"A valid input can be like:\n" +
	"1,2,3\n" +
	"4,5,6\n" +
	"7,8,9\n"

// Handler method
func Handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/echo", echoHandler)
	r.HandleFunc("/invert", invertHandler)
	r.HandleFunc("/flatten", flattenHandler)
	r.HandleFunc("/sum", sumHandler)
	r.HandleFunc("/multiply", multiplyHandler)
	return r
}

// This is a generic operation handler which takes an operation as a function parameter.
func operationHandler(w http.ResponseWriter, r *http.Request, fn operation.Operation) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Please provide a multi-part file.\n" +
			"For example, with curl command: \n" +
			"curl -F 'file=@matrix.csv' \"localhost:8080/echo\"")))
		return
	}
	defer file.Close()

	// We don't need to check if all our file is a legit CSV where all rows have the same number of columns.
	// This is done thank to the encoding/csv library.
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Sorry I cannot read the input file %s\nPlease try again.", err.Error())))
		return
	}

	operationResult, err := fn(records)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s\nTry again.", err.Error())))
		return
	}

	fmt.Fprint(w, operationResult)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf(Header)))
}

// Send request with:
//
//	curl -F 'file=@matrix.csv' "localhost:8080/echo"
func echoHandler(w http.ResponseWriter, r *http.Request) {
	operationHandler(w, r, operation.Echo)
}

// Send request with:
//
//	curl -F 'file=@matrix.csv' "localhost:8080/invert"
func invertHandler(w http.ResponseWriter, r *http.Request) {
	operationHandler(w, r, operation.Invert)
}

// Send request with:
//
//	curl -F 'file=@matrix.csv' "localhost:8080/flatten"
func flattenHandler(w http.ResponseWriter, r *http.Request) {
	operationHandler(w, r, operation.Flatten)
}

// Send request with:
//
//	curl -F 'file=@matrix.csv' "localhost:8080/sum"
func sumHandler(w http.ResponseWriter, r *http.Request) {
	operationHandler(w, r, operation.Sum)
}

// Send request with:
//
//	curl -F 'file=@matrix.csv' "localhost:8080/multiply"
func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	operationHandler(w, r, operation.Multiply)
}
