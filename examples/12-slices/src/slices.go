// _Slices_ are a key data type in Go, giving a more
// powerful interface to sequences than arrays.

package main

import "fmt"

func main() {
        // Unlink arrays, slices are a variable sequence of elements.
        // They are allocated on heap. We can consider them as
        // pointers to arrays.
        var slice []string;

	// You can allocate a slice using the make command,
        // specifying the type and the length. The slice's
        // elements are initialized with the default value 
        // of the element type. In this case, the empty string.
	slice = make([]string, 3)
	fmt.Println("slices:", slice)

	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
        fmt.Println("slices:", slice)
	fmt.Println("get:", slice[2])

	// `len` returns the length of the slice as expected.
	fmt.Println("len:", len(slice))

	// In addition to these basic operations, slices
	// support several more that make them richer than
	// arrays. One is the builtin `append`, which
	// returns a slice containing one or more new values.
	// Note that we need to accept a return value from
	// `append` as we may get a new slice value.
	slice = append(slice, "d")
	slice = append(slice, "e", "f")
	fmt.Println("slice:", slice)

	// Slices can also be `copy`'d. Here we create an
	// empty slice `c` of the same length as `s` and copy
	// into `c` from `s`.
	slice_copy := make([]string, len(slice))
	copy(slice_copy, slice)
	fmt.Println("slice copy:", slice_copy)

	// Slices support a "slice" operator with the syntax
	// `slice[low:high]`. For example, this gets a slice
	// of the elements `s[2]`, `s[3]`, and `s[4]`.
	l := s[2:5]
	fmt.Println("sl1:", l)

	// This slices up to (but excluding) `s[5]`.
	l = s[:5]
	fmt.Println("sl2:", l)

	// And this slices up from (and including) `s[2]`.
	l = s[2:]
	fmt.Println("sl3:", l)

	// We can declare and initialize a variable for slice
	// in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slices can be composed into multi-dimensional data
	// structures. The length of the inner slices can
	// vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
