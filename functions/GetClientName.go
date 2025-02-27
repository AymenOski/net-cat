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
		fmt.Println(err)
		return false
	}
	temp = strings.ReplaceAll(temp, "\r\n", "")
	temp = strings.ReplaceAll(temp, "\n", "")
	if len(temp) < 3 || len(temp) > 13 {
		conn.Write([]byte("❌ Invalid name! Must be between 3 and 15 characters.\n"))
		time.Sleep(1 * time.Second)
		conn.Write([]byte("Try again.."))
		time.Sleep(2 * time.Second)
	} else {
		*ClientName = temp
		return true
	}
	return false
}
