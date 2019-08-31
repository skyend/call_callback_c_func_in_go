package main

import "fmt"

import "C"
// import "unsafe"
//export Ftest
type Ftest func(C.int);

//export GoAdder
func GoAdder(x, y int, f Ftest) int {
	fmt.Printf("Go says: adding %v and %v\n", x, y) 
	f(10); 
	 
	return x + y
}

func main() {} // Required but ignored
