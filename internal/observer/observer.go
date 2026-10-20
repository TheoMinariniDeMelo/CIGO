package observer

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/TheoMinariniDeMelo/cigo/internal/protocol"
	"github.com/spf13/cobra"
)

var (
	dispacherIp string = "localhost:8888"
	repository string
	hostIp string = "localhost:8000"	
)

func Run() {
	rootCmd := &cobra.Command{
		Use:   "observer",
		Short: "Observer goci",
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

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()


}
