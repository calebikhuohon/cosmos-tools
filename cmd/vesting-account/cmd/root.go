package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:           "vesting-account",
		Short:         "A custom cli for generating vesting account analysis from the Umee chain Genesis files",
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
