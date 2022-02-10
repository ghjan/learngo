package rpcsupport

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

//ServeRpc 开一个Rpc server
func ServeRpc(host string, service interface{}, serverReady chan struct{}) (err error) {
	rpc.Register(service)
	if serverReady == nil {
		serverReady = make(chan struct{})
	}

	listener, err := net.Listen("tcp", host)
	if err != nil {
		return
	}
	fmt.Printf("listening in host:%s", host)
	readySent := false
	for {
		var (
			conn net.Conn
		)
		timer1 := time.AfterFunc(2*time.Second, func() {
			if err == nil && !readySent {
				serverReady <- struct{}{}
				readySent = true
			}
		})
		defer timer1.Stop()

		conn, err = listener.Accept()
		if err != nil {
			log.Printf("accept error:%v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)

	}
	return
}

//NewClient rpc client
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}
