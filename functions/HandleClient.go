package functions

import (
	"fmt"
	"net"
	"time"

	logger "net-cat/log"
	"net-cat/utils"
)

func HandleClient(conn net.Conn) {
	var ClientName string
	for {
		check, clientLost := GetClientName(conn, &ClientName)
		if clientLost {
			logger.Log(1, "The Client "+conn.LocalAddr().String()+" has lost connection"+"\n", nil)
			return
		}
		// to munimise time complexity we should put continue
		if !check {
			continue
		}
		for _, Client := range utils.Clients {
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
			fmt.Printf(utils.Green+"🟢%s has joined the groupe chat\n"+utils.Reset, ClientName)
			break
		}
	}
	utils.MU.Lock()
	utils.Clients[conn] = ClientName
	Broadcast(fmt.Sprintf(utils.Green+"🟢%s has joined our chat...\n"+utils.Reset, ClientName), conn)
	logger.Log(2, fmt.Sprintf("Client `%s` has joined the groupe chat...\n", ClientName), nil)
	conn.Write([]byte(fmt.Sprintf("[%s] [%s] : ", time.Now().Format("2006-01-02 15:04:05"), utils.Clients[conn])))
	utils.MU.Unlock()
	go SendingMsgs(conn)
}
