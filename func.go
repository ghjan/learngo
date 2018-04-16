package main

import (
"fmt"
"reflect"
"runtime"
"math"
)

func main() {
	if result, err := eval(3, 4, "*"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(eval(3, 4, "/"))
	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 3, 4))

	//匿名函数
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//return a / b
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation:%s", op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b

	// do not recommend here
	//q, r = a/b, a%b
	//return
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function :%s with args"+"(%d,%d)\n", opName, a, b)
	return op(a, b)
}
