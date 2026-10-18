package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/spf13/cobra"
)

var (
	dispacherIp string = "localhost:8888"
	repository string
	hostIp string = "localhost:8000"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "observer",
		Short: "Observer goci",
	}

	hostCmd := &cobra.Command{
		Use:   "host [ip]",
		Short: "Set host",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			hostIp = args[0]
			return nil
		},
	}
	dispacherCmd := &cobra.Command{
		Use:   "dispacher [ip]",
		Short: "Set dispacher ip",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dispacherIp = args[0]
			return nil
		},
	}

	repositoryCmd := &cobra.Command{
		Use: "dir [repo]",
		Short: "Set repository to CI",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			repository = args[0]
			return nil
		},
	}

	rootCmd.AddCommand(dispacherCmd)
	rootCmd.AddCommand(repositoryCmd)
	rootCmd.AddCommand(hostCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	listener, err := net.Listen("tcp", hostIp)

	if err != nil {
		log.Fatalf("Error while attempting to connect to the ip: %s", dispacherIp)
		os.Exit(1)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Erro:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
}
