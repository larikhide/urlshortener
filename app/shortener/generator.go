package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/itchyny/base58-go"
)

// hash the initial inputs during the process
func sha256of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

// The reasons are mostly related to the user experience, Base58 reduces
// confusion in character output.
// he characters 0,O, I, l are highly confusing when used in certain fonts and
// are even quite harder to differentiate for people with visuality issues .
// Removing ponctuations characters prevent confusion for line breakers.
// Double-clicking selects the whole number as one word if it's all alphanumeric.
func base58Encoded(bytes []byte) (string, error) {
	encode := base58.BitcoinEncoding
	encoded, err := encode.Encode(bytes)
	if err != nil {
		return "", fmt.Errorf("failed base58Encode: %w", err)
	}

	return string(encoded), nil
}

func GenShortLink(longUrl string) (string, error) {
	urlHashBytes := sha256of(longUrl)
	generatedNum := new(big.Int).SetBytes(urlHashBytes).Uint64()
	shortUrl, err := base58Encoded([]byte(fmt.Sprintf("%d", generatedNum)))
	if err != nil {
		return "", fmt.Errorf("failed short link generation: %w", err)
	}
	return shortUrl[:8], nil

}
