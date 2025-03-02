package functions

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	logger "net-cat/log"
)

func SendingMsgs(sender net.Conn) {
	for {
		// Reads and listen for messages
		msg, err := bufio.NewReader(sender).ReadString('\n')
		if err != nil {
			tempName := Clients[sender]
			fmt.Printf(Red+"🔴%s has left the groupe chat.\n"+Reset, Clients[sender])
			logger.Log(2, "The Client "+sender.LocalAddr().String()+" Has lost connection."+"\n", nil)
			logger.Log(2, fmt.Sprintf("Client `%s` has left the groupe chat...\n", Clients[sender]), nil)
			MU.Lock()
			delete(Clients, sender)
			sender.Close()
			MU.Unlock()
			Broadcast(fmt.Sprintf(Red+"🔴%s has left the chat.\n"+Reset, tempName), sender)
			return
		}
		if len(msg) > 0 {
			msg = strings.ReplaceAll(msg, "\r\n", "")
			msg = strings.ReplaceAll(msg, "\n", "")
			for Client := range Clients {
				// send the msg to all clients except the sender
				if Client != sender {
					Client.Write([]byte(fmt.Sprintf("\n[%s] [%s] : %s", time.Now().Format("2006-01-02 15:04:05"), Clients[sender], msg)))
					Client.Write([]byte(fmt.Sprintf("\n[%s] [%s] : %s", time.Now().Format("2006-01-02 15:04:05"), Clients[Client], "")))
				} else {
					// send this msg to the sender
					Client.Write([]byte(fmt.Sprintf("[%s] [%s] : %s", time.Now().Format("2006-01-02 15:04:05"), Clients[sender], "")))
				}
			}
		}
	}
}
