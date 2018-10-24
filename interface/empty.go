package main

import "fmt"


func main() {
	s := "Hello World"
	describe(s)
	i := 55
	describe(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(strt)


	// Use v, ok := i.(T)
	var s1 interface{} = 56
	assert(s1)
	var i1 interface{} = "Steven Paul"
	assert(i1)


	findType("Naveen")
	findType(77)
	findType(89.98)
}


/**
	An interface which has zero methods is called empty interface. It is represented as interface{}.
	Since the empty interface has zero methods, all types implement the empty interface.
 */

func describe(i interface{}) {

	// Type assertion is used extract the underlying value of the interface.
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

/**
	Type can be identified using v, ok := i.(T)
 */

func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

/**
	Use same as above but in switch block.
 */
func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}