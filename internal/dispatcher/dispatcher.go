package dispatcher

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/spf13/cobra"
)

var (
	dispacherIp string
	repository string
	hostIp string
)

func Run() {
	rootCmd := &cobra.Command{
		Use:   "dispatcher",
		Short: "dispacher goci",
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
		Use:   "host [host]",
		Short: "Set dispacher ip",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dispacherIp = args[0]
			return nil
		},
	}

	repositoryCmd := &cobra.Command{
		Use: "repository [repo]",
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

	listener, err := net.Listen("tcp", dispacherIp)

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

	var message string
	
	buf := make([]byte, 1024)
	
	for n, err := conn.Read(buf); n > 0; {
		message += buf
	}
}
