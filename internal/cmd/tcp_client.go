package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yagoyudi/gobench-tcp-udp/internal/benchmark"
	"github.com/yagoyudi/gobench-tcp-udp/internal/logger"
)

var tcpClientCmd = &cobra.Command{
	Use:   "client [address]",
	Short: "Run TCP client",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		addr := args[0]
		totalDataSize, err := strconv.Atoi(args[1])
		if err != nil {
			logger.FatalError(err)
		}
		err = benchmark.ClientTCP(addr, totalDataSize)
		if err != nil {
			logger.FatalError(err)
		}
	},
}
