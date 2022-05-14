package digest

import "crypto/sha256"

func GenerateDigest(data []byte) []byte {
	digest := sha256.Sum256(data)
	return digest[:]
}
