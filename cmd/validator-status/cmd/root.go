package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:           "validator-status",
		Short:         "A custom cli for generating CSVs of validator and delegator information on the Cosmos chains",
		SilenceUsage:  true,
		SilenceErrors: false,
	}
	chainName string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&chainName, "chain", "", "name of the cosmos chain")
	_ = rootCmd.MarkPersistentFlagRequired("chain")
}
