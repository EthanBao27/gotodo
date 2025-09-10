package network

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/ethanbao27/gotodo/internal/storage"
	"github.com/fatih/color"
)

func StartServer(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("listen error: %w", err)
	}

	color.New(color.FgBlue, color.Bold).Printf("Friend server started on %s\n", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			color.New(color.FgRed).Println("accept error:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		color.New(color.FgRed).Println("read error:", err)
		return
	}

	req := strings.TrimSpace(string(buf[:n]))
	if req != "GET_TODOS" {
		conn.Write([]byte("invalid request"))
		color.New(color.FgYellow).Printf("Invalid request from %s\n", conn.RemoteAddr())
		return
	}

	tasks, err := storage.List()
	if err != nil {
		conn.Write([]byte("failed to load tasks"))
		color.New(color.FgRed).Println("Failed to load local tasks")
		return
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		conn.Write([]byte("json error"))
		color.New(color.FgRed).Println("JSON marshal error")
		return
	}

	conn.Write(data)
	color.New(color.FgGreen).Printf("Shared %d tasks with %s\n", len(tasks), conn.RemoteAddr())
}
