package functions

import (
	"bufio"
	"net"
	"strings"
	"time"
)

func GetClientName(conn net.Conn, ClientName *string) bool {
	MX.Lock()
	conn.Write([]byte(Welcoming()))
	temp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		MU.Lock()
		// talk back to the server
		delete(Clients, conn)
		conn.Close()
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
			MX.Unlock()
		} else if k == 2 {
			conn.Write([]byte("❌ Invalid name! Inprintable caracteres are not allowed in the name.\n"))
			time.Sleep(1 * time.Second)
			conn.Write([]byte("👉Try again..\n"))
			time.Sleep(2 * time.Second)
			MX.Unlock()
		} else if k == 3 {
			conn.Write([]byte("❌ Invalid name! Spaces are not allowed in the name.\n"))
			time.Sleep(1 * time.Second)
			conn.Write([]byte("👉Try again..\n"))
			time.Sleep(2 * time.Second)
			MX.Unlock()
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
			MX.Unlock()
		} else {
			*ClientName = temp
			MX.Unlock()
			return true
		}
	}
	return false
}
