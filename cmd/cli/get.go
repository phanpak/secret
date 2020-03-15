package cli

import (
	"fmt"
	"log"

	"github.com/phanpak/secret"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a secret",
	Run: func(cmd *cobra.Command, args []string) {
		get(args)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}

func get(args []string) {
	vault := secret.File(encodingKey)
	key := args[0]
	value, err := vault.Get(key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Value of '%s': '%s'", key, value)
}
