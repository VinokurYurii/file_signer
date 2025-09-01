package utils

import (
	"file_signer/logger"
	"fmt"
	"os"

	"golang.org/x/term"
)

func GetPassphrase() ([]byte, error) {
	fmt.Println("Enter passphrase:")

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return nil, fmt.Errorf("failed to set terminal to raw mode: %w", err)
	}
	defer safeRestore(int(os.Stdin.Fd()), oldState)

	passphrase, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return nil, fmt.Errorf("failed to read passphrase: %w", err)
	}

	return passphrase, nil
}

func safeRestore(fd int, state *term.State) {
	if err := term.Restore(fd, state); err != nil {
		logger.HaltOnError(fmt.Errorf("failed to restore terminal state: %v", err), "Terminal restoration failed")
	}
}
