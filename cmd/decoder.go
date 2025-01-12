package cmd

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var encodedNodeID string

func init() {
	rootCmd.AddCommand(nodeIdDecoderCmd)
	nodeIdDecoderCmd.Flags().StringVarP(&encodedNodeID, "encoded", "e", "", "Your encoded node ID")
}

var nodeIdDecoderCmd = &cobra.Command{
	Use:   "dec",
	Short: "Decode a node id",
	Long:  `Decode a base64 encoded node id.`,
	Run: func(cmd *cobra.Command, args []string) {
		encodedNodeID, _ := cmd.Flags().GetString("encoded")
		if encodedNodeID == "" {
			fmt.Println("Please provide an encoded node ID using the --encoded flag.")
			return
		}

		decodedBytes, err := base64.URLEncoding.DecodeString(encodedNodeID)
		if err != nil {
			log.Fatalf("Error decoding node ID: %v", err)
		}

		decodedNodeID := string(decodedBytes)
		parts := strings.SplitN(decodedNodeID, ":", 2)

		if len(parts) != 2 {
			log.Fatalf("Invalid node ID format. Expected 'operator:id', got '%s'", decodedNodeID)
		}

		operator := parts[0]
		id := parts[1]

		log.Printf("Operator: %s", operator)
		log.Printf("ID: %s", id)
	},
}
