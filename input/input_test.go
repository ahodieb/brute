package input

import (
	"fmt"
	"strings"
)

func ExampleInput_Text() {
	var tstInput = `
Line1
Line2
Line3
`
	i := FromReader(strings.NewReader(tstInput))
	for i.Scan() {
		fmt.Println(i.Text())
	}

	// Output:
	// Line1
	// Line2
	// Line3
}

func ExampleInput_ReadText() {
	var tstInput = `Line1
Line2
Line3
`
	i := FromReader(strings.NewReader(tstInput))
	fmt.Println(i.ReadText())
	fmt.Println(i.ReadText())
	fmt.Println(i.ReadText())

	// Output:
	// Line1
	// Line2
	// Line3
}

func ExampleInput_Int() {
	var tstInput = `
1
2
3
`
	i := FromReader(strings.NewReader(tstInput))
	for i.Scan() {
		fmt.Println(i.Int())
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleInput_Ints() {
	var tstInput = `1 2 3 4 5`
	i := FromReader(strings.NewReader(tstInput))
	for i.Scan() {
		fmt.Println(i.Ints())
	}

	// Output:
	// [1 2 3 4 5]
}
