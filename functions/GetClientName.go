package functions

import (
	"bufio"
	"net"
	"strings"
	"time"

	"net-cat/utils"
)

func GetClientName(conn net.Conn, ClientName *string) (bool, bool) {
	// The second bool ensures that if the client disconnects during the name input phase, we properly handle cleanup and exit the goroutine.
	conn.Write([]byte(utils.Welcoming()))
	temp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		utils.MU.Lock()
		// talk back to the server
		conn.Close()
		utils.Cmp--
		utils.MU.Unlock()
		return false, true
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
			return true, false
		}
	}
	return false, false
}
