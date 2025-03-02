package functions

import (
	"fmt"
	"net"
	"sync"
	"time"

	logger "net-cat/log"
)

var (
	Clients = make(map[net.Conn]string, 10)
	MU      sync.Mutex
)

const (
	Red   = "\033[31m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

func HandleClient(conn net.Conn) {
	var ClientName string
	for {
		check := GetClientName(conn, &ClientName)
		for _, Client := range Clients {
			if Client == ClientName {
				conn.Write([]byte("This name is already taken.\n"))
				time.Sleep(1 * time.Second)
				conn.Write([]byte("👉Try again..\n"))
				time.Sleep(2 * time.Second)
				check = false
				continue
			}
		}
		if check {
			fmt.Printf(Green+"🟢%s has joined the groupe chat\n"+Reset, ClientName)
			break
		}
	}
	MU.Lock()
	Clients[conn] = ClientName
	Broadcast(fmt.Sprintf(Green+"🟢%s has joined our chat...\n"+Reset, ClientName), conn)
	logger.Log(2, fmt.Sprintf("Client `%s` has joined the groupe chat...\n", ClientName), nil)
	MU.Unlock()
	conn.Write([]byte(fmt.Sprintf("[%s] [%s] : ", time.Now().Format("2006-01-02 15:04:05"), Clients[conn])))
	go SendingMsgs(conn)
}
