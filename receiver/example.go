package main

import "fmt"

/**
	Implementing interfaces using pointer receivers vs value receivers
 */
type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}

func (p Person) Describe() { //implemented using value receiver
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() { //implemented using pointer receiver
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func main() {
	var d1 Describer
	p1 := Person{"Sam", 25}
	d1 = p1
	d1.Describe()
	p2 := Person{"James", 32}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address{"Washington", "USA"}

	/* compilation error if the following line is
	   uncommented
	   cannot use a (type Address) as type Describer
	   in assignment: Address does not implement
	   Describer (Describe method has pointer
	   receiver)
	*/
	//d2 = a

	d2 = &a //This works since Describer interface
	//is implemented by Address pointer in line 22
	d2.Describe()


	/*
	it is legal to call a pointer-valued method on anything that is already a pointer or
	whose address can be taken. The concrete value stored in an interface is not addressable
	and hence it is not possible for the compiler to automatically take the address of 'a' in line no. 48 (d2 = a)
	and hence this code fails.
	 */

}