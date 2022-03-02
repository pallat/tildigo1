package foobar

import "strconv"

func Say(n int) string {
	if n%15 == 0 {
		return "FooBar"
	}
	if n%5 == 0 {
		return "Bar"
	}

	if n%3 == 0 {
		return "Foo"
	}

	return strconv.Itoa(n)
}
