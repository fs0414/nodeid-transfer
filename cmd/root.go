package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var name string

func init() {
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "Your name")
}

var rootCmd = &cobra.Command{
	Use:   "nodeid-transfer",
	Short: "A brief description of your application",
	Long:  `A longer description of your CLI app...`,
	Run: func(cmd *cobra.Command, args []string) {
    if name == "" {
      log.Println("Please provide a name using the --name flag.")
			return
		}
    log.Println("nodeid-transfer use a " + name)
	},
}


func Execute() error {
	return rootCmd.Execute()
}
