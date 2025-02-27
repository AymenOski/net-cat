package functions

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func SendingMsgs(sender net.Conn) {
	for {
		msg, err := bufio.NewReader(sender).ReadString('\n')
		if err != nil {
			tempName := Clients[sender]
			MU.Lock()
			delete(Clients, sender)
			MU.Unlock()
			Broadcast(fmt.Sprintf("%s has left the chat.\n", tempName), sender)
			return
		}
		if len(msg) > 0 {
			msg = strings.ReplaceAll(msg, "\r\n", "")
			msg = strings.ReplaceAll(msg, "\n", "")
			for Client := range Clients {
				if Client != sender {
					Client.Write([]byte(fmt.Sprintf("\n[%s] [%s] : %s", time.Now().Format("2006-01-02 15:04:05"), Clients[sender], msg)))
					Client.Write([]byte(fmt.Sprintf("\n[%s] [%s] : %s", time.Now().Format("2006-01-02 15:04:05"), Clients[Client], "")))
				} else {
					Client.Write([]byte(fmt.Sprintf("[%s] [%s] : %s", time.Now().Format("2006-01-02 15:04:05"), Clients[sender], "")))
				}
			}
		}
	}
}
