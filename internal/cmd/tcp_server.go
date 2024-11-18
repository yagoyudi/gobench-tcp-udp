package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/yagoyudi/gobench-tcp-udp/internal/benchmark"
)

var tcpServerCmd = &cobra.Command{
	Use:   "server [address]",
	Short: "Run TCP server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addr := args[0]
		err := benchmark.ServerTCP(addr)
		if err != nil {
			log.Fatal(err)
		}
	},
}
