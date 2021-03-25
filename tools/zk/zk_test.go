package zk

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"testing"
	"time"
)

var (
	acls = zk.WorldACL(zk.PermAll) //控制访问权限模式
	//0:永久，除非手动删除
	//zk.FlagEphemeral = 1:短暂，session断开则改节点也被删除
	//zk.FlagSequence  = 2:会自动在节点后面添加序号
	//3:Ephemeral和Sequence，即，短暂且自动添加序号
	flag int32 = 0
)

func GetConnect() (*zk.Conn, <-chan zk.Event, error) {
	hosts := []string{"localhost:2181"}
	option := zk.WithEventCallback(callback)
	return zk.Connect(hosts, time.Second*5,option)
}

func Create(conn *zk.Conn, path string, data string) bool {
	_, err := conn.Create(path, []byte(data), flag, acls)
	if err == nil {
		return true
	}
	return false
}

func Update(conn *zk.Conn, path string, data string) bool {
	_, stat, err := Get(conn, path)
	if err == nil {
		_, err := conn.Set(path, []byte(data), stat.Version)
		if err == nil {
			return true
		}
		fmt.Println(err)
	}
	fmt.Println(err)
	return false
}

func Get(conn *zk.Conn, path string) (data string, stat *zk.Stat, err error) {
	get, stat, err := conn.Get(path)
	return string(get), stat, err
}

func Del(conn *zk.Conn, path string) bool {
	_, stat, err := Get(conn, path)
	if err == nil {
		err := conn.Delete(path, stat.Version)
		if err == nil {
			return true
		}
		fmt.Println(err)
	}
	fmt.Println(err)
	return false
}

func TestZK(t *testing.T) {
	// 获取zk连接
	connect, _, err := GetConnect()
	if err != nil {
		panic(err)
	}

	defer connect.Close()

	// ================ CRUD  ================
	path := "/test"

	//watch
	//connect.ExistsW(path)

	//Create

	var data string = "hello world"
	_,_,ech,_:=connect.ExistsW(path)
	go NodeListener(ech)
	create := Create(connect, path, data)
	fmt.Println("create: ",create)

	data, _, _ = Get(connect, path)
	fmt.Println("get data: ",data)

	//connect.ExistsW(path)
	Update(connect,path,"hello zk")
	data, _, _ = Get(connect, path)
	fmt.Println("update data: ",data)

	//connect.ExistsW(path)
	Del(connect,path)
	data, _, _ = Get(connect, path)
	fmt.Println("delete data: ",data)


}

func NodeListener(ech <-chan zk.Event)  {
	event := <-ech
	fmt.Println("********NodeListener***********")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("*********NodeListener**********")
}

func callback(event zk.Event) {
	fmt.Println("*******************")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("*******************")
}
