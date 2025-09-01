package keys

import "github.com/spf13/cobra"

func init() {
	keysCmd.AddCommand(keysGenerateCommand)
}

var keysGenerateCommand = &cobra.Command{
	Use:   "generate",
	Short: "Generates key pair.",
	Long:  `Generate an RSA key pair and store it in PEM files. The private key will be encrypted using a passphrase that you'll need to enter. AES encryption with Argon2 key derivation function is utilized.`,
	Run:   func(cmd *cobra.Command, args []string),
}
