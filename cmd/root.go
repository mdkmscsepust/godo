package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "godo",
	Short: "godo cli",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}