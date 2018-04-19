package typedef

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

//定义一个函数类型
type IntGen func() int
type BufIntGen struct {
	G   IntGen
	Buf bytes.Buffer
}

//函数也能够作为接受者
func (g IntGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	return strings.NewReader(s).Read(p)
}

func (b *BufIntGen) Read(p []byte) (n int, err error) {
	if b.Buf.Len() == 0 {
		next := b.G()
		if next > 10000 {
			return 0, io.EOF
		}
		_, err := fmt.Fprintf(&b.Buf, "%d\n", next)
		if err != nil {
			return 0, err
		}
	}
	return b.Buf.Read(p)
}
