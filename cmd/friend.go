/*
Copyright © 2025 Ethan Bao
*/
package cmd

import (
	"fmt"

	"github.com/ethanbao27/gotodo/internal/network"
	"github.com/spf13/cobra"
)

// friendCmd represents the friend command
var friendCmd = &cobra.Command{
	Use:   "friend",
	Short: "Manage friends and shared todo lists",
}

var serveCmd = &cobra.Command{
	Use:   "serve [ip]",
	Short: "Start a friend server to share your todo list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ip := args[0]
		addr := fmt.Sprintf("%s:8088", ip)
		return network.StartServer(addr)
	},
}

var connectCmd = &cobra.Command{
	Use:   "connect [ip]",
	Short: "Connect to a friend and fetch their todo list (ip address)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return network.ConnectAndFetch(args[0])
	},
}

func init() {
	friendCmd.AddCommand(serveCmd)
	friendCmd.AddCommand(connectCmd)
	rootCmd.AddCommand(friendCmd)
}
