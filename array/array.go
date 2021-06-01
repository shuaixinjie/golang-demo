package array

import "fmt"

func demo1()  {
	var a [3]int             // array of 3 integers
	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}


	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"

}
