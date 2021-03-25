package zk

import (
	"errors"
	"fmt"
	"github.com/go-zookeeper/zk"
	"io/ioutil"
	"math/rand"
	"net"
	"testing"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func TestStartClient(t *testing.T) {
	for i := 0; i < 100; i++ {
		startClient()
		time.Sleep(time.Second)
	}
}

func startClient() {
	srvHost, err := getServerHost()
	checkError(err)
	fmt.Println("\n【connect host】: " + srvHost)
	addr, err := net.ResolveTCPAddr("tcp4", srvHost)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, addr)
	checkError(err)
	defer conn.Close()

	_, err = conn.Write([]byte("timestamp"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	return
}

func getServerHost() (host string, err error) {
	connect, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second * 5)
	checkError(err)
	defer connect.Close()

	srvList, err := GetServerList(connect)
	checkError(err)
	count := len(srvList)

	if count == 0 {
		err = errors.New("server list is empty")
		return
	}

	//随机选中一个返回
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	host = srvList[r.Intn(1)]
	return

}
func GetServerList(conn *zk.Conn) (list []string, err error) {
	list, _, err = conn.Children("/go_servers")
	return
}
