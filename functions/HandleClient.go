package functions

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	Clients = make(map[net.Conn]string, 10)
	MU      sync.Mutex
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
			fmt.Printf("🟢%s has joined the groupe chat\n", ClientName)
			break
		}
	}
	MU.Lock()
	Clients[conn] = ClientName
	MU.Unlock()

	Broadcast(fmt.Sprintf("🟢%s has joined our chat...\n", ClientName), conn)
	conn.Write([]byte(fmt.Sprintf("[%s] [%s] : ", time.Now().Format("2006-01-02 15:04:05"), Clients[conn])))
	go SendingMsgs(conn)
}
