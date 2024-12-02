package histogram

import (
	"fmt"
)

func ExampleNew() {
	h := New[string]()
	h.Add("b")
	h.Add("a")
	h.Add("c")
	h.Add("b")

	fmt.Println(h)
	// Output:
	// {"a":1,"b":2,"c":1}
}

func ExampleFromSlice() {
	h := FromSlice([]int{1, 2, 2, 3, 5, 5, 5})

	fmt.Println(h)
	// Output:
	// {"1":1,"2":2,"3":1,"5":3}
}

func ExampleFromItems() {
	h := FromItems(1, 2, 2, 3, 5, 5, 5)

	fmt.Println(h)
	// Output:
	// {"1":1,"2":2,"3":1,"5":3}
}

func ExampleHistogram_Add() {
	h := New[string]()
	h.Add("b", "a", "c")
	h.Add("c", "b")
	h.Add("b", "f")

	fmt.Println(h)
	// Output:
	// {"a":1,"b":3,"c":2,"f":1}
}

func ExampleHistogram_Count() {
	h := FromSlice([]rune("aaabcdd"))

	fmt.Println(h.Count('a'))
	fmt.Println(h.Count('b'))
	fmt.Println(h.Count('c'))
	fmt.Println(h.Count('d'))

	// Output:
	// 3
	// 1
	// 1
	// 2
}
