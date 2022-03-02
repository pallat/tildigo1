package foobar

import "strconv"

func Say(n int) string {
	if n == 5 {
		return "Bar"
	}

	if n%3 == 0 {
		return "Foo"
	}

	return strconv.Itoa(n)
}
