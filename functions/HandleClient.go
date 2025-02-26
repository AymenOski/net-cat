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
	buff := make([]byte, 1)
	// writes data to the connection
	conn.Write([]byte(Welcoming()))
	for len(buff) < 3 || len(buff) > 15 {
		_, err := conn.Read(buff)
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(buff) < 3 || len(buff) > 15 {
			conn.Write([]byte("The name must be between 3 and 15 characters\n"))
			time.Sleep(2 * time.Second)
		}
	}

	MU.Lock()
	Clients[conn] = string(buff[:len(buff)-1])
	MU.Unlock()

	Broadcast("One member has join the chat "+string(Clients[conn]), conn)

	time := time.Now().Format("2006-01-02 15:04:05")

	conn.Write([]byte(fmt.Sprintf("[%s] [%s] : ", time, string(Clients[conn]))))
}
