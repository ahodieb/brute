// Package grid contains common utilities for dealing with 2d arrays / grids
package grid

import (
	"fmt"
	"iter"
)

// Grid represents a 2d array or a grid
// internally it wrappers a two-dimensional slice of any type
// and provides common helper functions to manipulate grids
//
// The grid stats from the top left
//
//	[0,0] [0,1] [0,2]
//	[1,0] [1,1] [1,2]
//	[2,0] [2,1] [2,2]
//	[3,0] [3,1] [3,2]
type Grid[T any] struct {
	Values [][]T
}

// Point represents a point in the grid
type Point struct {
	Row    int
	Column int
}

// Cell represents a cell in the grid, it includes both the coordinates Point and the Value
type Cell[T any] struct {
	Value T
	Point Point
}

// FirstInRow is true if the point is the first column in a given row
func (p *Point) FirstInRow() bool {
	return p.Column == 0
}

// Region represents an area in a grid
type Region struct {
	Start Point
	End   Point
}

// Size represents the size of the grid
type Size struct {
	Rows    int
	Columns int
}

// Includes indicates if a Point is within the cell position in the grid
func (g *Grid[T]) Includes(p Point) bool {
	return p.Row >= 0 && p.Row < len(g.Values) && p.Column >= 0 && p.Column < len(g.Values[0])
}

// Value returns the cell value of a specific position
// will panic (index out of range) if the position is out of bounds
func (g *Grid[T]) Value(p Point) T {
	return g.Values[p.Row][p.Column]
}

// Cell returns the cell at a specific Point
// will panic (index out of range) if the position is out of bounds
func (g *Grid[T]) Cell(p Point) Cell[T] {
	return Cell[T]{
		Point: p,
		Value: g.Value(p),
	}
}

// AppendRow appends a row to the grid
func (g *Grid[T]) AppendRow(values []T) {
	if len(g.Values) != 0 {
		if size := len(g.Values[0]); size != len(values) {
			panic(fmt.Sprintf(
				"cannot add row with values %+v\nnew row must be of the same size as previous rows, got %d expected %d",
				values, len(values), size,
			))
		}
	}

	g.Values = append(g.Values, values)
}

// Size returns the size of the grid
func (g *Grid[T]) Size() Size {
	rows := len(g.Values)
	if rows == 0 {
		return Size{}
	}

	return Size{
		Rows:    rows,
		Columns: len(g.Values[0]),
	}
}

// Rows Returns the underlying two-dimensional array
// In many cases Iterator might be a better option
func (g *Grid[T]) Rows() [][]T {
	return g.Values
}

// Cells returns an iterator iter.Seq over the grid cells
// it iterates the grid left to right from top to bottom.
func (g *Grid[T]) Cells() iter.Seq[Cell[T]] {
	return func(yield func(Cell[T]) bool) {
		for r, row := range g.Values {
			for c, v := range row {
				cell := Cell[T]{
					Point: Point{Row: r, Column: c},
					Value: v,
				}
				if !yield(cell) {
					return
				}
			}
		}
	}
}

// IntToStr converts Grid[int] to a Grid[string] that can be used for rendering
func IntToStr(g *Grid[int]) Grid[string] {
	gg := Grid[string]{
		Values: make([][]string, 0, len(g.Values)),
	}
	for _, row := range g.Values {
		r := make([]string, 0, len(row))
		for _, c := range row {
			r = append(r, fmt.Sprintf("%d", c))
		}
		gg.AppendRow(r)
	}
	return gg
}
