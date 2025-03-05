package functions

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	logger "net-cat/log"
	"net-cat/utils"
)

func SendingMsgs(sender net.Conn) {
	for {
		// Reads and listen for messages
		msg, err := bufio.NewReader(sender).ReadString('\n')
		if err != nil {
			utils.MU.Lock()
			tempName := utils.Clients[sender]
			fmt.Printf(utils.Red+"🔴%s has left the groupe chat.\n"+utils.Reset, utils.Clients[sender])
			logger.Log(2, "The Client "+sender.LocalAddr().String()+" Has lost connection."+"\n", nil)
			logger.Log(2, fmt.Sprintf("Client `%s` has left the groupe chat...\n", utils.Clients[sender]), nil)
			delete(utils.Clients, sender)
			sender.Close()
			utils.Cmp--
			fmt.Println(utils.Cmp)
			Broadcast(fmt.Sprintf(utils.Red+"🔴%s has left the chat.\n"+utils.Reset, tempName), sender)
			utils.MU.Unlock()
			return
		}
		if len(msg) > 0 {
			msg = strings.ReplaceAll(msg, "\r\n", "")
			msg = strings.ReplaceAll(msg, "\n", "")
			for Client := range utils.Clients {
				// send the msg to all clients except the sender
				if Client != sender {
					Client.Write([]byte(fmt.Sprintf("\n[%s] [%s] : %s", time.Now().Format("2006-01-02 15:04:05"), utils.Clients[sender], msg)))
					Client.Write([]byte(fmt.Sprintf("\n[%s] [%s] : %s", time.Now().Format("2006-01-02 15:04:05"), utils.Clients[Client], "")))
				} else {
					// send this msg to the sender
					Client.Write([]byte(fmt.Sprintf("[%s] [%s] : %s", time.Now().Format("2006-01-02 15:04:05"), utils.Clients[sender], "")))
				}
			}
		}
	}
}
