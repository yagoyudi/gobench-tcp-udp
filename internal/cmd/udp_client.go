package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/yagoyudi/gobench-tcp-udp/internal/benchmark"
)

func init() {
	udpClientCmd.Flags().String("total", "250mb", "Total payload to be transfered (250mb|500mb|1gb|2gb|4gb)")
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
		totalDataSize := parseTotalFlag(total)
		err = benchmark.ClientUDP(addr, totalDataSize)
		if err != nil {
			log.Fatal(err)
		}
	},
}
