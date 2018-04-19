package main

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ghjan/learngo/functional/fib"
	"github.com/ghjan/learngo/functional/typedef"
)

//闭包 adder函数返回了一个闭包（不仅仅是一个函数，闭包还包括一个环境）
func adder() func(int) int {
	sum := 0 //环境中的自由变量
	return func(v int) int { //局部变量
		sum += v
		return sum
	}
}

func main() {
	//fmt.Println("----closure1----")
	//closure1()
	fmt.Println("----closure2----")
	closure2()
	fmt.Println("----testFib----")
	testFib()
	fmt.Println("----testFibSmallBuf----")
	testFibSmallBuf()

}

func closure2() {
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0+1+...+%d=%d\n", i, s)
	}
}
func closure1() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+1+...+%d=%d\n", i, a(i))
	}
}

//正统的函数式编程
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func testFib() {
	f := fib.Fibonacci()
	PrintFileContents(f)
}

func testFibSmallBuf() {
	var f = fib.Fibonacci()
	p := make([]byte, 2)
	//外面缓冲区开的很小p is too small, so we use typedef.intGen instead of typedef.intGen
	r := typedef.BufIntGen{G: f}
	for {
		n, err := r.Read(p)
		if err != nil {
			break
		}
		fmt.Printf("%s", p[:n])
	}
}

func PrintFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
