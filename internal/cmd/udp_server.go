package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yagoyudi/gobench-tcp-udp/internal/benchmark"
	"github.com/yagoyudi/gobench-tcp-udp/internal/logger"
)

var udpServerCmd = &cobra.Command{
	Use:   "server [address]",
	Short: "Run UDP server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addr := args[0]
		err := benchmark.ServerUDP(addr)
		if err != nil {
			logger.FatalError(err)
		}
	},
}
