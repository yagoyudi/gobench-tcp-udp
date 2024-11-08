package cmd

import (
	"github.com/spf13/cobra"
)

var udpCmd = &cobra.Command{
	Use:   "udp",
	Short: "UDP benchmarking commands",
	Long:  `Run UDP client and server for benchmarking purposes.`,
}

func init() {
	udpCmd.AddCommand(udpClientCmd)
	udpCmd.AddCommand(udpServerCmd)
}
