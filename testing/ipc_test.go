package testing

import (
	"fmt"
	"testing"

	"example.com/hello/ipc"
)

func TestIpcServerClients(t *testing.T) {
	server := new(ipc.Server)
	ipcServer := ipc.NewIpcServer(*server)
	client1 := ipc.NewIpcClient(ipcServer)
	client2 := ipc.NewIpcClient(ipcServer)
	resp1, err := client1.Call("GET", "from client1")
	if err != nil {
		return
	}
	fmt.Println(resp1)
	resp2, err := client2.Call("POST", "from client2")
	if err != nil {
		return
	}
	fmt.Println(resp2)
	client1.Close()
	client2.Close()
}
