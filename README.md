# Backend Takehome

Available endpoints are:

* `/echo` -- Returns the matrix as a string.
* `/invert` -- Returns the matrix as a string where the columns and rows are inverted.
* `/flatten` -- Returns the matrix as a 1 line string.
* `/sum` -- Returns the sum of the integers in the matrix.
* `/multiply` -- Returns the product of the integers in the matrix.

The input file to these functions is a matrix, of any dimension where the number of rows is equal to the number of columns (square).
Each value is an integer, and there is no header row.

A valid input can be like:
```
1,2,3
4,5,6
7,8,9
```

The input file to these functions is a matrix, of any dimension where the number of rows is equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is an example valid input. 

Run webserver
```
go run . &
```
You can also, build, test and run the code with our provided Makefile.
Build with:
```
make build
```
Test with:
```
make test
```
Run with
```
make run
```
To generate a Test Coverage Report in HTML run
```
make test-html
```

Send request
```
curl -F 'file=@matrix.csv' "localhost:8080/echo"
```

## Design Summary
This code has been developed using Test Driven Development, aiming at the highest coverage level possible (100%).

It is separated into two distinct packages:
* `handlers` -- for all Http routing handlers.
* `operation` -- this one takes care of all the operations of our service (echo, invert, flatten, sum, multiply). In this package the `check.go` file takes care of the error handling. We are guarding against, non-squared and empty matrix, overflow in sum and multiplication.