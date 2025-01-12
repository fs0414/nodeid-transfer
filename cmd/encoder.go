package cmd

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var operator string 
var id string

func init() {
  rootCmd.AddCommand(nodeIdEncoderCmd)
  nodeIdEncoderCmd.Flags().StringVarP(&operator, "operator", "o", "", "Your operator")
  nodeIdEncoderCmd.Flags().StringVarP(&id, "id", "i", "", "Your id")
}

var nodeIdEncoderCmd = &cobra.Command{
  Use:   "enc",
  Short: "Encode a node id",
  Long:  `Encode a node id using base64 encoding.`,
  Run: func(cmd *cobra.Command, args []string) {
    operator, _ := cmd.Flags().GetString("operator")
    id, _ := cmd.Flags().GetString("id")
    if operator == "" {
      log.Println("Please provide a operator using the --operator flag.")
      return
    }
    if id == "" {
      log.Println("Please provide a id using the --id flag.")
      return
    }
    nodeID := fmt.Sprintf("%s:%s", operator, id)
    enc := base64.URLEncoding.EncodeToString([]byte(nodeID))
    log.Println(enc)
  },
}
