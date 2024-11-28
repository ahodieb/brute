package grid

import (
	"fmt"
)

func ExampleAnsiRenderer_Render() {
	g := Grid[string]{Values: [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}}
	r := NewAnsiRenderer().SetDefaultFmt(square)
	fmt.Println(r.Render(&g))

	// Output:
	// [ 1 ][ 2 ][ 3 ]
	// [ 4 ][ 5 ][ 6 ]
	// [ 7 ][ 8 ][ 9 ]
}

func ExampleAnsiRenderer_FmtPoints() {
	g := Grid[string]{Values: [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}}
	r := NewAnsiRenderer().SetDefaultFmt(square)
	r.FmtPoints(triangle, p(0, 0), p(1, 1), p(2, 2))
	fmt.Println(r.Render(&g))

	// Output:
	// < 1 >[ 2 ][ 3 ]
	// [ 4 ]< 5 >[ 6 ]
	// [ 7 ][ 8 ]< 9 >
}

func ExampleAnsiRenderer_FmtFn() {
	g := Grid[string]{Values: [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}}
	r := NewAnsiRenderer().SetDefaultFmt(square)
	r.FmtFn(triangle, func(c Cell[string]) bool {
		return (c.Point.Row*g.Size().Columns+c.Point.Column+1)%2 == 0
	})
	fmt.Println(r.Render(&g))

	// Output:
	// [ 1 ]< 2 >[ 3 ]
	// < 4 >[ 5 ]< 6 >
	// [ 7 ]< 8 >[ 9 ]
}

func square(s string) string {
	return fmt.Sprintf("[%s]", s)
}

func triangle(s string) string {
	return fmt.Sprintf("<%s>", s)
}

func p(x, y int) Point {
	return Point{
		Row: x, Column: y,
	}
}
