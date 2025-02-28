package functions

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func GetClientName(conn net.Conn, ClientName *string) bool {
	conn.Write([]byte(Welcoming()))
	temp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		MU.Lock()
		// talk back to the server
		delete(Clients, conn)
		MU.Unlock()
		return false
	}
	check, k := FiltringCheck(temp)
	if !check {
		if k == 1 {
			conn.Write([]byte("❌ Invalid name! The name is empty.\n"))
			time.Sleep(1 * time.Second)
			conn.Write([]byte("👉Try again..\n"))
			time.Sleep(2 * time.Second)
		} else if k == 2 {
			conn.Write([]byte("❌ Invalid name! Inprintable caracteres are not allowed in the name.\n"))
			time.Sleep(1 * time.Second)
			conn.Write([]byte("👉Try again..\n"))
			time.Sleep(2 * time.Second)
		} else if k == 3 {
			conn.Write([]byte("❌ Invalid name! Spaces are not allowed in the name.\n"))
			time.Sleep(1 * time.Second)
			conn.Write([]byte("👉Try again..\n"))
			time.Sleep(2 * time.Second)
		}
	}
	if check && k == 0 {
		temp = strings.ReplaceAll(temp, "\r\n", "")
		temp = strings.ReplaceAll(temp, "\n", "")
		if len(temp) < 3 || len(temp) > 13 {
			conn.Write([]byte("❌ Invalid name! Must be between 3 and 15 characters.\n"))
			time.Sleep(1 * time.Second)
			conn.Write([]byte("👉Try again..\n"))
			time.Sleep(2 * time.Second)
		} else {
			*ClientName = temp
			fmt.Printf("🟢%s has joined the groupe chat\n", *ClientName)
			return true
		}
	}
	return false
}
