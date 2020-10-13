// In Go, _variables_ are explicitly declared and used by
// the compiler to e.g. check type-correctness of function
// calls.
package main

import "fmt"

type MyType struct {
    field1 int
    field2 string
}

func main() {
        // default value is 0
	var myVariable1 int
        fmt.Println(myVariable1)

	var myVariable2 int = 10
        fmt.Println(myVariable2)

	var myVariable3 float64 = 5.0
        fmt.Println(myVariable3)

        // default value is 0.0
	var myVariable4 float32
        fmt.Println(myVariable4)

        // default value is false
	var myVariable5 bool
        fmt.Println(myVariable5)

	var myVariable6 bool = true
        fmt.Println(myVariable6)

	var myVariable7 string = "MyString"
        fmt.Println(myVariable7)

        // default value is ""
	var myVariable8 string
        fmt.Println(myVariable8)

        // MyType is a composite type.
        // Default value is the struct with default value
        // for each field
	var myVariable9 MyType
        fmt.Println(myVariable9)

	// You can declare multiple variables at once.
	var myVariable10, myVariable11 int = 1, 2
	fmt.Println(myVariable10, myVariable11)

	// Go will infer the type of initialized variables.
	var myVariable12 = true
	fmt.Println(myVariable12)

	// The `:=` syntax is shorthand for declaring and
	// initializing a variable, e.g. for
	// `var f string = "apple"` in this case.
	myVariable13 := "apple"
	fmt.Println(myVariable13)

        myVariable14 := 5
        fmt.Println(myVariable14)

        myVariable15 := false
        fmt.Println(myVariable15)
}
