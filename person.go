package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func getPerson() Person {
	// Using new() to allocate memory for Person struct
	/* Use:
	- new() when dealing with value types like structs to allocate memory for a new zeroed value.
	  This is suitable for scenarios where you want a pointer to an initialized structure.
	  new() returns a pointer
	- make() to create an initialized instance for slices, maps, and channels, where initialization involves setting up data structures
	  and internal pointers
	  make() returns a non-zeroed value
	*/
	p := new(Person)

	p.Name = "Jon Doe"
	p.Age = 39
	p.Gender = "male"

	fmt.Println(p)

	// Using make() for initializing slices, maps, and channels – data structures that require runtime initialization
	s := make([]int, 10, 15)
	for i := 0; i < 10; i++ {
		s[i] = i + 1
	}
	fmt.Println(s)
	return *p
}
