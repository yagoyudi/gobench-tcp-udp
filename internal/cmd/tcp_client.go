package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yagoyudi/gobench-tcp-udp/internal/benchmark"
	"github.com/yagoyudi/gobench-tcp-udp/internal/logger"
)

func init() {
	tcpClientCmd.Flags().String("total", "10mb", "Total payload to be transfered (10mb|100mb|500mb|1gb)")
}

var tcpClientCmd = &cobra.Command{
	Use:   "client [address]",
	Short: "Run TCP client",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addr := args[0]
		total, err := cmd.Flags().GetString("total")
		if err != nil {
			logger.FatalError(err)
		}
		var totalDataSize int
		switch total {
		case "10mb":
			totalDataSize = tenMB
		case "100mb":
			totalDataSize = hundredMB
		case "500mb":
			totalDataSize = fiveHundredMB
		case "1gb":
			totalDataSize = oneGB
		}
		if err != nil {
			logger.FatalError(err)
		}
		err = benchmark.ClientTCP(addr, totalDataSize)
		if err != nil {
			logger.FatalError(err)
		}
	},
}
