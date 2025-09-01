package main

import (
	"file_signer/cmd"
	"file_signer/cmd/keys"
	"file_signer/cmd/signatures"
	"file_signer/logger"
)

func main() {
	rootCmd := cmd.RootCmd()
	keys.Init(rootCmd)
	signatures.Init(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.HaltOnError(err, "Initial setup failed")
	}
}
