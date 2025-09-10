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
		if _, err := conn.Write([]byte("invalid request")); err != nil {
			fmt.Println("write error:", err)
			return
		}
		color.New(color.FgYellow).Printf("Invalid request from %s\n", conn.RemoteAddr())
		return
	}

	tasks, err := storage.List()
	if err != nil {
		if _, err := conn.Write([]byte("failed to load tasks")); err != nil {
			fmt.Println("write error:", err)
			return
		}
		color.New(color.FgRed).Println("Failed to load local tasks")
		return
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		if _, err := conn.Write([]byte("json error")); err != nil {
			fmt.Println("write error:", err)
			return
		}
		color.New(color.FgRed).Println("JSON marshal error")
		return
	}

	if _, err := conn.Write(data); err != nil {
		fmt.Println("write error:", err)
		return
	}
	color.New(color.FgGreen).Printf("Shared %d tasks with %s\n", len(tasks), conn.RemoteAddr())
}
