package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	dispacherIp string
	repository string
)

func main() {
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

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}

