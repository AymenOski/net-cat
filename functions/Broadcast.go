package functions

import (
	"fmt"
	"net"
	"time"
)

func Broadcast(message string, sender net.Conn) {
	for Client := range Clients {
		if Client != sender {
			_, err := Client.Write([]byte("\n" + message))
			if err != nil {
				return
			}
			Client.Write([]byte(fmt.Sprintf("[%s] [%s] : ", time.Now().Format("2006-01-02 15:04:05"), Clients[Client])))
		}
	}
}
