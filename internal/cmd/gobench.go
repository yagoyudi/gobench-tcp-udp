package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var gobenchCmd = &cobra.Command{
	Use:   "gobench",
	Short: "Gobench is a benchmarking tool for TCP and UDP.",
	Long:  `Gobench is a CLI tool for benchmarking TCP and UDP performance by running clients and servers.`,
}

func Execute() {
	err := gobenchCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	gobenchCmd.AddCommand(udpCmd)
	gobenchCmd.AddCommand(tcpCmd)
}
