package main

import (
	"net"
	"net/rpc/jsonrpc"
	"github.com/ghjan/learngo/lang/rpc"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64
	if err = client.Call("DemoService.Div", rpcdemo.Args{10, 3}, &result); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	if err = client.Call("DemoService.Div", rpcdemo.Args{10, 0}, &result); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
