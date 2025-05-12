package exercises

import (
	"crypto/sha256"
	"github.com/ulvenforst/gobook/cmd/chap2"
)

// Hamming distance, basically
func RunCompareSHA256(a, b string) int {
	digestA := sha256.Sum256([]byte(a))
	digestB := sha256.Sum256([]byte(b))

	var counter int

	for i := range digestA {
		counter += int(chap2.Pc[digestA[i]^digestB[i]])
	}

	return counter

}
