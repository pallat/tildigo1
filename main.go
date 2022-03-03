package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gitlab.com/cjexpress/tildi/signac/learngo/foobar"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/foobar/:number", func(c *gin.Context) {
		number := c.Param("number")

		c.JSON(200, gin.H{
			"foobar": foobar.SayAny(number),
		})
	})

	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		pic[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			pic[y][x] = uint8(x ^ y*x&y)
		}
	}

	return pic
}

func mainOfClosure() {
	fibonacci := fibonacciFunc()

	fmt.Println(fibonacci(), fibonacci(), fibonacci(), fibonacci(), fibonacci(), fibonacci(), fibonacci(), fibonacci(), fibonacci(), fibonacci())
}

func fibonacciFunc() func() int {
	a, b := 0, 1

	return func() int {
		defer func() { a, b = b, a+b }()

		return a
	}
}

func Show(fn func(dx, dy int) [][]uint8) {
	colorPixels := fn(256, 256)
	fmt.Println(colorPixels)
}

func mainOfRandomSay() {
	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)
	fmt.Println(foobar.RandomSay(r))
}

func mainOfEmptyInterface() {
	var obj interface{}

	fmt.Println(obj == nil)

	obj = struct{}{}
	fmt.Printf("type: %T, value %#v\n", obj, obj)

	obj = 10
	fmt.Printf("type: %T, value %#v\n", obj, obj)

	obj = "ten"
	fmt.Printf("type: %T, value %#v\n", obj, obj)

	if s, ok := obj.(string); ok {
		fmt.Printf("type(s): %T, value(s) %#v\n", s, s)
	}
}

func tryDefer() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("catched ", r)
		}
	}()

	log.Fatal("panic")

	fmt.Println("end tryDefer")
}

type Int int

func (i *Int) toString() string {
	return "convert to string as " + strconv.Itoa(int(*i))
}

func (i *Int) String() string {
	return strconv.Itoa(int(*i))
}

func (i *Int) Set(n int) {
	*i = Int(n)
}

func mainOfMethod() {
	var n Int = 10

	fmt.Printf("%q\n", n.String())

	n.Set(20)
	fmt.Printf("%q\n", n.String())
}

func mainOfOscar() {
	f, err := os.Open("./oscar.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	nameCount := map[string]int{}
	for _, record := range records {
		nameCount[record[3]]++
	}

	for name, count := range nameCount {
		if count > 1 {
			fmt.Println(name)
		}
	}
}

func mainOfMap() {
	s := []int{2, 3, 4, 5}
	s = append(s[:1], s[2:]...)

	fmt.Println(s)

	m := map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "coconut",
	}

	m["d"] = "date"
	m["e"] = "elderberry"

	for k := range m {
		fmt.Println(k)
	}

}

func variadic(a ...int) {
	fmt.Printf("%T %v\n", a, a)
}

func couple(ss string) (r []string) {
	ss += "*"
	s := []rune(ss)

	for ; len(s) > 1; s = s[2:] {
		r = append(r, string(s[:2]))
	}
	return
}

func mainOfSlice() {
	var s []int
	x := [...]int{4, 5, 6, 7, 8}
	s = x[1:3]

	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	x[1] = 15
	fmt.Println(s)

	y := append(s, 20, 21, 22)
	fmt.Println(y)

	y[1] = 16
	fmt.Println(x)
}

func mainOfArray() {
	x := [...]int{4, 5, 6, 7, 8}

	x[1] = 12
	fmt.Printf("%T %v\n", x, x)

	for _, v := range x {
		fmt.Println(v)
	}
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
