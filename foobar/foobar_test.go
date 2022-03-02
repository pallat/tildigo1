package foobar_test

import (
	"testing"

	"gitlab.com/cjexpress/tildi/signac/learngo/foobar"
)

func TestFooBarGiven1WantsString1(t *testing.T) {
	given := 1
	wants := "1"

	get := foobar.Say(given)

	if wants != get {
		t.Errorf("given %d wants %q but got %q\n", given, wants, get)
	}
}

func TestFooBarGiven2WantsString2(t *testing.T) {
	given := 2
	wants := "2"

	get := foobar.Say(given)

	if wants != get {
		t.Errorf("given %d wants %q but got %q\n", given, wants, get)
	}
}
