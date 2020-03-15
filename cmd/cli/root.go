package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = cobra.Command{
	Use:   "secret",
	Short: "Secret manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello")
	},
}

var encodingKey string

func Execute() {
	RootCmd.PersistentFlags().StringVarP(&encodingKey, "key", "k", "", "the key to use when encoding and decoding secrets")
	RootCmd.Execute()
}
