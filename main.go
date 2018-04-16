package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"io/ioutil"
)

func main() {
	euler()
	var a, b = 3, 4
	fmt.Println(triangle(a, b))
	a, b = 3, 5
	fmt.Printf("%.3f\n", triangle(a, b))
	consts()
	enums()
	const filename = "abc.txt"
	printFile(filename)
	fmt.Println(bounded(3))

	fmt.Println(grade(0))
	fmt.Println(grade(59))
	fmt.Println(grade(60))
	fmt.Println(grade(90))
	fmt.Println(grade(100))

}

func euler() {
	fmt.Printf("%.3f\n", cmplx.Pow(math.E, 1i*math.Pi)+1)
}

func triangle(a, b int) float64 {
	return math.Sqrt(float64(a*a + b*b))
}

func consts() {
	//golang里面的常量不要大写开头，因为golang里面的大写是有另外的特殊含义（public）
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp        = iota
		_
		python
		golang
		javascript
	)
	const (
		b  = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, python, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func bounded(v int) int {
	if v > 100 {
		return 100
	} else if v > 50 {
		return 50
	} else if v > 10 {
		return 10
	} else {
		return 0
	}
}

func printFile(filename string) {

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

}

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return g
}

