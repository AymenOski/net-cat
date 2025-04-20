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
---
## Files

- `server.go` - TCP server that handles multiple clients and clean shutdown.
- `client.go` - Simple TCP client to connect to the server.

## How to Run

### Start the Server

```bash
go run ./TCPChat/
```

### Start Clients

In separate terminals:

```bash
nc localhost 2525
```

Enjoy chatting! 💬🚀