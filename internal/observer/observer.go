package observer

import (
	"net"
	"os"
	"github.com/spf13/cobra"
)

var (
	dispacherIp string = "localhost:8888"
	repository 	string
	hostIp 		string = "localhost:8000"
	configFile 	string = "./cigo.json"
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

	configCmd := &cobra.Command{
		Use: "config [configFile]",
		Short: "Set config file for CI",
		RunE: func(cmd *cobra.Command, args []string) error {
			configFile = args[0]
			return nil
		},
	}

	rootCmd.AddCommand(dispacherCmd)
	rootCmd.AddCommand(repositoryCmd)
	rootCmd.AddCommand(configCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
