package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)
var helloCmd = &cobra.Command{
	Use: "hello",
	Short: "hello cli",
	Run: func (cmd *cobra.Command, args []string)  {
		name := viper.GetString("name")
		fmt.Println("Hello, ", name)
	},
}

func init() {
	helloCmd.Flags().StringP("name", "n", "Masum", "Your name")
	viper.BindPFlag("name", helloCmd.Flags().Lookup("name"))
	rootCmd.AddCommand(helloCmd)
}