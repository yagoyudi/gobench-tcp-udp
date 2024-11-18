package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/yagoyudi/gobench-tcp-udp/internal/benchmark"
)

func init() {
	udpClientCmd.Flags().String("total", "10mb", "Total payload to be transfered (10mb|100mb|500mb|1gb)")
}

var udpClientCmd = &cobra.Command{
	Use:   "client [server-address:port]",
	Short: "Run UDP client",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addr := args[0]
		total, err := cmd.Flags().GetString("total")
		if err != nil {
			log.Fatal(err)
		}
		totalDataSize, err := parseTotalFlag(total)
		if err != nil {
			log.Fatal(err)
		}
		err = benchmark.ClientUDP(addr, totalDataSize)
		if err != nil {
			log.Fatal(err)
		}
	},
}
