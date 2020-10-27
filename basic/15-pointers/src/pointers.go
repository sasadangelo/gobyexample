// Go supports <em><a href="http://en.wikipedia.org/wiki/Pointer_(computer_programming)">pointers</a></em>,
// allowing you to pass references to values and records
// within your program.
package main

import "fmt"

// We'll show how pointers work in contrast to values with
// 2 functions: `zeroval` and `zeroptr`. `zeroval` has an
// `int` parameter, so arguments will be passed to it by
// value. `zeroval` will get a copy of `ival` distinct
// from the one in the calling function.
func zeroval(ival int) {
	ival = 0
}

// `zeroptr` in contrast has an `*int` parameter, meaning
// that it takes an `int` pointer. The `*iptr` code in the
// function body then _dereferences_ the pointer from its
// memory address to the current value at that address.
// Assigning a value to a dereferenced pointer changes the
// value at the referenced address.
func zeroptr(iptr *int) {
	*iptr = 0
}

type person struct {
        name string
        age int
}

func main() {
        // Use of pointers with primitive types
        var j int = 1
        var k *int = &j

        l:= &j
        fmt.Println("j:", j)
        fmt.Println("k->", *k)
        fmt.Println("l->", *l)

        // Use of pointers with composite types.
        // Note the dot operator with pointers de-reference
        // it automatically
        var john person=person{"John", 50}
        var pjohn=&john

        fmt.Println("John age:", john.age)
        fmt.Println("John age:", pjohn.age)

        // Use of pointers as function parameters
	zeroval(j)
	fmt.Println("zeroval:", j)

	// The `&j` syntax gives the memory address of `j`,
	// i.e. a pointer to `j`.
	zeroptr(k)
	fmt.Println("j:", j)

	// Pointers can be printed too.
	fmt.Println("pointer:", &j)
}
