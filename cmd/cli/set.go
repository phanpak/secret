package cli

import (
	"log"

	"github.com/phanpak/secret"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set a secret",
	Run: func(cmd *cobra.Command, args []string) {
		handle(args)
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}

func handle(args []string) {
	vault := secret.File(encodingKey)
	key, value := args[0], args[1]
	err := vault.Set(key, value)
	if err != nil {
		log.Fatal(err)
	}
}
