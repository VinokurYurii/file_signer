package keys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"file_signer/utils"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

type PrivateKeyGen struct {
	outputPath string
	keyBitSize int
	saltSize   int
}

func init() {
	keysCmd.AddCommand(keysGenerateCommand)

	keysGenerateCommand.Flags().String("pub-out", "pub_key.pem", "Path to save the public key")
	keysGenerateCommand.Flags().String("priv-out", "priv_key.pem", "Path to save the pprivate key")
	keysGenerateCommand.Flags().Int("priv-size", 2048, "Private key size in bits")
	keysGenerateCommand.Flags().Int("salt-size", 16, "Salt size used in key derivation in bytes")
}

var keysGenerateCommand = &cobra.Command{
	Use:   "generate",
	Short: "Generates key pair.",
	Long:  `Generate an RSA key pair and store it in PEM files. The private key will be encrypted using a passphrase that you'll need to enter. AES encryption with Argon2 key derivation function is utilized.`,
	Run: func(cmd *cobra.Command, args []string) {
		pkOut, _ := cmd.Flags().GetString("priv-out")
		pkSize, _ := cmd.Flags().GetInt("priv-size")
		saltSize, _ := cmd.Flags().GetInt("salt-size")

		pkGenConfig := PrivateKeyGen{
			outputPath: pkOut,
			keyBitSize: pkSize,
			saltSize:   saltSize,
		}
	},
}

func generatePrivKey(pkGenConfig PrivateKeyGen) (*rsa.PrivateKey, error) {
	absPath, err := filepath.Abs(pkGenConfig.outputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %w", err)
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, pkGenConfig.keyBitSize)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %w", err)
	}

	passphrase, err := utils.GetPassphrase()
	if err != nil {
		return nil, err
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	salt, err := makeSalt(pkGenConfig.saltSize)
	if err != nil {
		return nil, err
	}
}

func makeSalt(saltSize int) ([]byte, error) {
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return nil, fmt.Errorf("failed to generate salt: %v", err)
	}

	return salt, nil
}
