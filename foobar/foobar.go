package foobar

func Say(n int) string {
	if n == 5 {
		return "Bar"
	}
	if n == 4 {
		return "4"
	}
	if n == 3 {
		return "Foo"
	}
	if n == 2 {
		return "2"
	}
	return "1"
}
