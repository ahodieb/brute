// Package input
// Provides an input handler for common use cases for online problems.
package input

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Input handles reading and parsing the inputs for problems
// internally it uses a bufio.Scanner
type Input struct {
	scanner *bufio.Scanner
	closer  func()
}

// FromPath creates an input handler from a file
func FromPath(path string) *Input {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return &Input{
		scanner: bufio.NewScanner(f),
		closer:  func() { _ = f.Close() },
	}
}

// FromReader creates an input handler from a reader
func FromReader(r io.Reader) *Input {
	return &Input{
		scanner: bufio.NewScanner(r),
		closer:  func() {},
	}
}

// Scan reads and buffers one line from input
// The buffered value can be accessed by Input.Text or Input.Int or similar
func (i *Input) Scan() bool { return i.scanner.Scan() }

// Slurp reads all lines till the end of the input
func (i *Input) Slurp() []string {
	var lines []string
	for i.Scan() {
		lines = append(lines, i.Text())
	}

	return lines
}

// Text returns the buffered line of text
func (i *Input) Text() string {
	return i.scanner.Text()
}

// ReadText scans for the next line, and if one is found returns it, or it returns an empty string
// Internally it calls both Input.Scan and Input.Text
func (i *Input) ReadText() string {
	i.Scan()
	return i.Text()
}

// Int returns a parsed int value
// Parsing errors are ignored, if `BRUTE_DEBUG` is set it panics.
func (i *Input) Int() int {
	n, err := strconv.Atoi(i.scanner.Text())
	if err != nil && os.Getenv("BRUTE_DEBUG") != "" {
		panic(fmt.Errorf("failed to parse int: <%s>, %w", i.scanner.Text(), err))
	}

	return n
}

// Ints returns slice of parsed int values
// the numbers are expected to be space seperated
func (i *Input) Ints() []int {
	return parseInts(i.Text())
}

func parseInts(s string) []int {
	var values []int
	tokens := strings.Fields(s)
	for _, token := range tokens {
		n, err := strconv.Atoi(token)
		if err == nil {
			values = append(values, n)
		} else if os.Getenv("BRUTE_DEBUG") != "" {
			panic(fmt.Errorf("failed to parse int: <%s>, %w", token, err))
		}
	}

	return values
}

// Close the inner buffer
// a noop if a reader or stdin was passed in
func (i *Input) Close() { i.closer() }
