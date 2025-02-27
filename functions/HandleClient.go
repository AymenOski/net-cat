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
		if check {
			break
		}
	}
	MU.Lock()
	Clients[conn] = ClientName
	MU.Unlock()

	Broadcast(fmt.Sprintf("%s has joined our chat...\n", ClientName), conn)
	conn.Write([]byte(fmt.Sprintf("[%s] [%s] : ", time.Now().Format("2006-01-02 15:04:05"), Clients[conn])))
	go SendingMsgs(conn)
}
