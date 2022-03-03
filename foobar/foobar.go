package foobar

import (
	"fmt"
	"strconv"
)

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

type Intner interface {
	Intn(n int) int
}

func RandomSay(r Intner) string {
	return fmt.Sprintf("%s-%s-%s-%s", Say(r.Intn(9)+1), Say(r.Intn(9)+1), Say(r.Intn(9)+1), Say(r.Intn(9)+1))
}
