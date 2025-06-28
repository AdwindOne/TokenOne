package wallets

import (
	"github.com/tyler-smith/go-bip39"
)

// GenerateMnemonic creates a new BIP39 mnemonic phrase.
// It defaults to 128 bits of entropy (12 words).
func GenerateMnemonic() (string, error) {
	// Default entropy is 128 bits, which results in a 12-word mnemonic.
	// For more words, increase entropy (e.g., 256 bits for 24 words).
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return "", err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}

// ValidateMnemonic checks if the given string is a valid BIP39 mnemonic.
func ValidateMnemonic(mnemonic string) bool {
	return bip39.IsMnemonicValid(mnemonic)
}
