package grid

import (
	"fmt"
	"github.com/ahodieb/brute/ansi"
	"strings"
)

// AnsiRenderer is  text based renderer for Grid
// Ansi codes are used to color and format the cells
// It technically can render any grid size, but it is more sutuble
// small grids that can fit in a typical terminal window
type AnsiRenderer struct {
	defaultFmt formatter
	posFmt     map[Point]formatter
	condFmt    []conditionalFormatter
}

// NewAnsiRenderer creates a new ansi renderer
func NewAnsiRenderer() *AnsiRenderer {
	return &AnsiRenderer{
		defaultFmt: AnsiFormatter(ansi.WhiteBG).toFmt(),
		posFmt:     make(map[Point]formatter),
	}
}

// SetDefaultFmt changes the default formatter
func (r *AnsiRenderer) SetDefaultFmt(f AnsiFormatter) *AnsiRenderer {
	r.defaultFmt = f.toFmt()
	return r
}

// FmtPoints sets a different format for the specified points
// useful for marking points on a grid with different colors
func (r *AnsiRenderer) FmtPoints(f AnsiFormatter, p ...Point) *AnsiRenderer {
	for _, pp := range p {
		r.posFmt[pp] = f.toFmt()
	}

	return r
}

// FmtFn sets a conditional formatter for a cell
// All cells within the grid that fit the specified condition would be formated with the specified formatter
func (r *AnsiRenderer) FmtFn(f AnsiFormatter, fn func(Cell[string]) bool) *AnsiRenderer {
	r.condFmt = append(r.condFmt, conditionalFormatter{Condition: fn, Formatter: f.toFmt()})
	return r
}

// Render returns the full ansi string render of the Grid
func (r *AnsiRenderer) Render(g *Grid[string]) string {
	var sb strings.Builder
	for cell := range g.Cells() {
		if cell.Point.FirstInRow() {
			_, _ = fmt.Fprintln(&sb)
		}

		formatter := r.defaultFmt

		if f, found := r.posFmt[cell.Point]; found {
			formatter = f
		}

		for _, cf := range r.condFmt {
			if cf.Condition(cell) {
				formatter = cf.Formatter
			}
		}

		_, _ = fmt.Fprint(&sb, formatter(cell))
	}

	return sb.String()
}

// AnsiFormatter is a function applying ansi code formats ( e.g. ansi.WhitBG )
type AnsiFormatter func(string) string

func (f AnsiFormatter) toFmt() formatter {
	return func(cell Cell[string]) string {
		return f(fmt.Sprintf(" %s ", cell.Value))
	}
}

type formatter func(cell Cell[string]) string

type conditionalFormatter struct {
	Condition func(Cell[string]) bool
	Formatter formatter
}
