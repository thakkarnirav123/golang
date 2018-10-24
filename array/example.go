package main

import "fmt"

func main() {

	intArrayWithoutValue()

	stringArrayWithoutValue()

	intArrayWithValue()

	copyArray()

	iterateArray()

	iterateUsingRange()

	iterateUsingRangeWithoutIndex()

	multiArray()
}

func intArrayWithoutValue() {

	var x [5]int
	fmt.Println(x)
}

func stringArrayWithoutValue() {

	var y [8]string
	fmt.Println(y)
}


func intArrayWithValue() {

	//Note that we do not need to specify the type of the variable a as in var a [5]int,
	//because the compiler can automatically infer the type from the expression on the right hand side.

	var a = [5]int{2, 4, 6, 8, 10}
	fmt.Print(a)

	// Short hand declaration
	b := [5]int{2, 4, 6, 8, 10}
	fmt.Print(b)


	// You don't need to initialize all the elements of the array.
	// The un-initialized elements will be assigned the zero value of the corresponding array type
	c := [5]int{2}
	fmt.Println(c)

	// Letting Go compiler infer the length of the array
	d := [...]int{3, 5, 7, 9, 11, 13, 17}
	fmt.Println(d)
}

func copyArray() {

	a1 := [5]string{"English", "Japanese", "Spanish", "French", "Hindi"}
	a2 := a1

	// A copy of the array `a1` is assigned to `a2`
	a2[1] = "German"
	fmt.Println("a1 = ", a1)

	// The array `a1` remains unchanged
	fmt.Println("a2 = ", a2)
}

/**
	Use len() function.
 */
func iterateArray() {

	names := [3]string{"Mark Zuckerberg", "Bill Gates", "Larry Page"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}


	a := [4]float64{3.5, 7.2, 4.8, 9.5}
	sum := float64(0)

	for i := 0; i < len(a); i++ {
		sum = sum + a[i]
	}

	fmt.Printf("Sum of all the elements in array  %v = %f\n", a, sum)
}



func iterateUsingRange() {

	daysOfWeek := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	for index, value := range daysOfWeek {
		fmt.Printf("Day %d of week = %s\n", index, value)
	}

	a := [4]float64{3.5, 7.2, 4.8, 9.5}
	sum := float64(0)

	for i, value := range a {
		sum = sum + value
		i = i+0
	}
	fmt.Printf("Sum of all the elements in array %v = %f", a, sum)
}

/**
	Go compiler doesn’t allow creating variables that are never used. You can fix this by using an _ (underscore) in place of index -
 */
func iterateUsingRangeWithoutIndex() {

	a := [4]float64{3.5, 7.2, 4.8, 9.5}
	sum := float64(0)
	for _, value := range a {
		sum = sum + value
	}
	fmt.Printf("Sum of all the elements in array %v = %f", a, sum)
}

func multiArray() {

	a := [2][2]int{
		{3, 5},
		{7, 9}, // This trailing comma is mandatory
	}
	fmt.Println(a)
	// Just like 1D arrays, you don't need to initialize all the elements in a multi-dimensional array.
	// Un-initialized array elements will be assigned the zero value of the array type.
	b := [3][4]float64{
		{1, 3},
		{4.5, -3, 7.4, 2},
		{6, 2, 11},
	}
	fmt.Println(b)
}



