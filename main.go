package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "hello",
		Short: "Simple greeting CLI",
		Run: func(cmd *cobra.Command, args []string) {
			name := viper.GetString("name")
			fmt.Println("hello, ", name)
		},
	}

	// rootCmd.Flags().String("name", "Masum", "Your name")
	// viper.BindPFlag("name", rootCmd.Flags().Lookup("name"))

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.ReadInConfig()
	rootCmd.Execute();
}