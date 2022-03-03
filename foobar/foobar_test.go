package foobar_test

import (
	"fmt"
	"testing"

	"gitlab.com/cjexpress/tildi/signac/learngo/foobar"
)

func TestFooBar(t *testing.T) {
	givenWant := map[int]string{
		1:  "1",
		2:  "2",
		4:  "4",
		7:  "7",
		8:  "8",
		3:  "Foo",
		6:  "Foo",
		9:  "Foo",
		12: "Foo",
		5:  "Bar",
		10: "Bar",
		20: "Bar",
		25: "Bar",
		15: "FooBar",
		30: "FooBar",
		45: "FooBar",
		60: "FooBar",
	}

	for given, wants := range givenWant {
		t.Run(fmt.Sprintf("given %d wants %q", given, wants), func(t *testing.T) {
			get := foobar.Say(given)

			if wants != get {
				t.Errorf("given %d wants %q but got %q\n", given, wants, get)
			}
		})
	}
}

type fakeIntner struct {
	count  int
	series []int
}

func (f *fakeIntner) Intn(n int) int {
	defer func() { f.count++ }()

	return f.series[f.count]
}

func TestRandomSay(t *testing.T) {
	given := &fakeIntner{
		count:  0,
		series: []int{0, 1, 2, 4},
	}

	wants := "1-2-Foo-Bar"

	get := foobar.RandomSay(given)

	if wants != get {
		t.Errorf("given %d as Intn returns value want %q but got %q\n", given, wants, get)
	}
}
