package main

import (
	"fmt"
)

func main() {
}

func mainOfPointer() {
	var p *int

	if p != nil {
		fmt.Println(*p)
	}

	i := 42
	p = &i

	fmt.Printf("%x %x\n", p, &i)
	fmt.Println(*p, i)

	*p = 43
	fmt.Println(*p, i)
}

func power(b, x int) int {
	r := 1
	for i := 0; i < x; i++ {
		r *= b
	}
	return r
}

func prime(n int) {
	for i := 2; i <= n; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			fmt.Println(i)
		}
	}
}

func mainOfBasic() {
	var name string
	fmt.Printf("Hello, %q\n", name)

	name = "Pallat"
	fmt.Println("Hello,", name)

	var age int = 45
	fmt.Printf("I'm %d years old\n", age)

	var nickName = "Yod"
	fmt.Printf("call me %s\n", nickName)

	fmt.Printf("type is %T\n", nickName)

	sex := "male"
	fmt.Printf("%v\n", sex)

	fmt.Printf("1 + 2 = %d\n", add(1, 2))

	a, b := swap(1, 2)
	fmt.Println(a, b)

	if ok := IsCorrect(); ok {
		println("It's correct")
	}

}

func IsCorrect() bool {
	return true
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func swap(a, b int) (int, int) {
	return b, a
}

func greeting(name string) {
	fmt.Printf("Hello %s\n", name)
}

func add(a int, b int) int {
	return a + b
}
