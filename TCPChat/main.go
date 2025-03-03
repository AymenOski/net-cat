package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

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
	defer ln.Close()
	fmt.Println("Chat Server Started : server listening for connections on port", port)
	// Cmp := 0
	logger.Log(1, fmt.Sprintf("Chat Server Started : server listening for connections on the port %s\n", port), err)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	// Channeling the os.Interrupt and syscall.SIGTERM signals
	go func() {
		<-sigChan
		fmt.Println("\nThe server is closing")
		logger.Log(1, "Chat Server Closed. The server is no longer listening.", err)
		ln.Close()
		if functions.Clients != nil {
			for conn := range functions.Clients {
				conn.Write([]byte("The server is closed, please disconnect!\n"))
				conn.Close()
			}
		}
		os.Exit(0)
	}()
	for {
		conn, err := ln.Accept() // this accept data that u can manipulate
		if err != nil {
			continue // keep trying to connect
		} else {
			// Cmp++
			logger.Log(2, "New connection from "+conn.LocalAddr().String()+"\n", nil)
		}
		// if 10 clients are connected, send a message to the new client
		// if Cmp > 2 {
		// 	conn.Write([]byte("The group is full 10/10 , please wait for someone to disconnect!\n"))
		// 	conn.Close()
		// 	continue
		// }
		go functions.HandleClient(conn)
	}
}
