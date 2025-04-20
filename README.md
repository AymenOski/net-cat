# 🧠 TCPChat – NetCat Clone in Go

## 📌 Description

**TCPChat** is a simple group chat application that mimics the behavior of the famous `netcat (nc)` tool. This project implements a server-client architecture using TCP connections. It allows multiple clients to join a chat room and exchange messages in real-time.

The main goal is to understand and apply concepts like TCP networking, sockets, Go concurrency (goroutines, channels, mutexes), and how to handle client-server communication in an efficient and safe way.

---

## 🎯 Features

- ✅ TCP server that listens for incoming client connections
- ✅ Each client has a unique name (e.g., user1234)
- ✅ Server broadcasts all received messages to all connected clients
- ✅ Server handles unexpected client disconnections gracefully
- ✅ Uses Go concurrency (goroutines) to manage multiple clients simultaneously
- ✅ Messages are timestamped
- ✅Broadcast messages from one client to all others.
- ✅Server supports clean shutdown using OS signals (Ctrl+C).
- ✅Clients are notified when the server shuts down.
- ✅ Implemented `.log` file to track all connections and server status, including client notifications when the server shuts down.
---
## Files

- `server.go` - TCP server that handles multiple clients and clean shutdown.
- `client.go` - Simple TCP client to connect to the server.

## Usage

### 1. Start the Server
```bash
$ go run ./TCPChat/
```
Expected Output:
```
➜  net-cat git:(main) go run ./TCPChat 
Chat Server Started : server listening for connections on port 8989
🟢aymen has joined the groupe chat
🟢hosin has joined the groupe chat
🔴hosin has left the groupe chat.
🔴aymen has left the groupe chat.
^C
The server is closing
```

### 2. Connect as a Client (New Terminal)
```bash
$ nc localhost 8080
```
Expected Output:
```
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]:hosin
[2025-04-20 11:55:00] [hosin] : hello
[2025-04-20 11:55:07] [hosin] : 
[2025-04-20 11:55:13] [aymen] : hi there!
[2025-04-20 11:55:13] [hosin] : see u man    
[2025-04-20 11:55:30] [hosin] : 
[2025-04-20 11:55:40] [aymen] : oki
[2025-04-20 11:55:40] [hosin] : ^C
```

### 3. Another Client Connection (New Terminal)
```bash
$ nc localhost 8080
```
Expected Output:
```
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]:aymen
[2025-04-20 11:54:56] [aymen] : 
🟢hosin has joined our chat...
[2025-04-20 11:55:00] [aymen] : 
[2025-04-20 11:55:07] [hosin] : hello
[2025-04-20 11:55:07] [aymen] : hi there!
[2025-04-20 11:55:13] [aymen] : 
[2025-04-20 11:55:30] [hosin] : see u man
[2025-04-20 11:55:30] [aymen] : oki
[2025-04-20 11:55:40] [aymen] : 
🔴hosin has left the chat.
[2025-04-20 11:55:42] [aymen] : ^C
```

---

## Default Settings
In `main.go`, you can configure:
```go
MaxConnections = 10 // Max connections count
Port = ":8989"      // Default port if the user didnt set
```

Feel free to modify these as per your requirements.

Enjoy chatting! 💬🚀