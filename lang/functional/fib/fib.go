package fib

import "github.com/ghjan/learngo/lang/functional/typedef"

func Fibonacci() typedef.IntGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}
