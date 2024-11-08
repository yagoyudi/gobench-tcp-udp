package cmd

import (
	"github.com/spf13/cobra"
)

var tcpCmd = &cobra.Command{
	Use:   "tcp",
	Short: "TCP benchmarking commands",
	Long:  `Run TCP client and server for benchmarking purposes.`,
}

func init() {
	tcpCmd.AddCommand(tcpClientCmd)
	tcpCmd.AddCommand(tcpServerCmd)
}
