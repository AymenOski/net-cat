package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"net-cat/functions"
	logger "net-cat/log"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	port := "8989"
	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Log(1, "", err)
		log.Fatal(err)
	}
	fmt.Println("Chat Server Started : server listening for connections on port", port)
	logger.Log(0, port, err)
	defer ln.Close()
	for {
		conn, err := ln.Accept() // this accept data that u can manipulate
		if err != nil {
			log.Println(err)
			continue // keep trying to connect
		} else {
			logger.Log(3, conn.RemoteAddr().String()+"\n", nil)
		}
		go functions.HandleClient(conn)
	}
}
