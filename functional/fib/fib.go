package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"bytes"
)

//定义一个函数类型
type intGen func() int
type bufIntGen struct {
	g   intGen
	buf bytes.Buffer
}

func (b *bufIntGen) Read(p []byte) (n int, err error) {
	if b.buf.Len() == 0 {
		next := b.g()
		if next > 10000 {
			return 0, io.EOF
		}
		_, err := fmt.Fprintf(&b.buf, "%d\n", next)
		if err != nil {
			return 0, err
		}
	}
	return b.buf.Read(p)
}

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

//函数也能够作为接受者
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	testFib()

	testFibSmallBuf()
}
func testFib() {
	f := fibonacci()
	printFileContents(f)
}

func testFibSmallBuf() {
	var f = fibonacci()
	p := make([]byte, 2)
	//外面缓冲区开的很小p is too small, so we use bufIntGen instead of intGen
	r := bufIntGen{g: f}
	for {
		n, err := r.Read(p)
		if err != nil {
			break
		}
		fmt.Printf("%s", p[:n])
	}
}
