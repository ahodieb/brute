// Package ansi provides some helpers to use ANSI coloring
// https://en.wikipedia.org/wiki/ANSI_escape_code#Colors:~:text=Graphic%20Rendition)%20parameters-,Colors,-3%2Dbit%20and

package ansi

import "fmt"

func Red(s string) string {
	return fmt.Sprintf("\u001b[31m%s\u001b[0m", s)
}

func RedBG(s string) string {
	return fmt.Sprintf("\u001b[41m%s\u001b[0m", s)
}

func Green(s string) string {
	return fmt.Sprintf("\u001b[32m%s\u001b[0m", s)
}

func GreenBG(s string) string {
	return fmt.Sprintf("\u001b[42m%s\u001b[0m", s)
}

func YellowBG(s string) string {
	return fmt.Sprintf("\u001b[43m%s\u001b[0m", s)
}

func BlueBG(s string) string {
	return fmt.Sprintf("\u001b[44m%s\u001b[0m", s)
}

func PurpleBG(s string) string {
	return fmt.Sprintf("\u001b[45m%s\u001b[0m", s)
}

func WhiteBG(s string) string {
	return fmt.Sprintf("\u001b[47m%s\u001b[0m", s)
}

func Underline(s string) string {
	return fmt.Sprintf("\u001b[4m%s\u001b[0m", s)
}
