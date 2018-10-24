package main

import "fmt"

/* global variable declaration */
var g int

func main() {

	local()
	global()
}

func local() {

	/* local variable declaration */
	var a, b, c int

	/* actual initialization */
	a = 10
	b = 20
	c = a + b
	fmt.Printf("value of a = %d, b = %d and c = %d\n", a, b, c)
}

func global(){

	fmt.Printf("value of g = %d\n",  g);

}
