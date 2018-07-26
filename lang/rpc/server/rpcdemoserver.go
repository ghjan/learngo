package main

import (
	"net/rpc"
	"github.com/ghjan/learngo/lang/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error:%v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

/* use telnet to test rpc server
[root@izuf6go9cwac1w1u5nxv6sz ~]# telnet localhost 1234
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
{"method":"DemoService.Div","params":[{"A":3, "B":4}],"id":1}
{"id":1,"result":0.75,"error":null}
{"method":"DemoService.Div","params":[{"A":3, "B":0}],"id":1234}
{"id":1234,"result":null,"error":"division by zero"}

 */
