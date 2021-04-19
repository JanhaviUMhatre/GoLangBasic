package main

import (
	"fmt"
)

func main() {
	var a float32 = 10
	var b float32 = 3

	var add float32 = a + b
	var sub float32 = a - b
	var div float32 = a / b
	var mul float32 = a*b
	var d int = 10
	var e int = 3
	mod := d % e
	
	fmt.Printf(
		"addition: %v \nsubtraction: %v \ndivision: %.2f \nmodulo: %v \nmultiplication: %v",
		add, sub, div, mod, mul)
}