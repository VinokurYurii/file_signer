package keys

import "github.com/spf13/cobra"

var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Manage key pairs.",
	Long:  "Use subcommands to create publix/private key pairs in PEM files.",
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(keysCmd)
}
