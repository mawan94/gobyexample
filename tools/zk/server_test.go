package zk

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"net"
	"testing"
	"time"
)

func TestServerStart(t *testing.T) {
	startServer("127.0.0.1:8897")
	forever := make(chan bool, 1)
	<- forever
}



func RegisterServer(conn *zk.Conn, host string) (err error) {
	conn.Create("/go_servers", nil, 0, zk.WorldACL(zk.PermAll))
	path := fmt.Sprintf("/go_servers/%s", host)
	_, err = conn.Create(path, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return
}


func handleClient(conn net.Conn,port string)  {
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte("---- " + port + ": " + daytime+"---- "))
}



func startServer(port string) {
	addr, err := net.ResolveTCPAddr("tcp4", port)
	checkError(err)

	listener, err := net.ListenTCP("tcp4", addr)
	checkError(err)

	// zk
	connect, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second * 5)
	checkError(err)
	defer connect.Close()

	err = RegisterServer(connect,port)
	checkError(err)

	for  {
		conn,err :=listener.Accept()
		if err != nil {
			checkError(err)
			continue
		}
		go handleClient(conn,port)
	}
}
