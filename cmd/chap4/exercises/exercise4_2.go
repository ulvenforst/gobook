package exercises

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

const (
	SHA256 = "256"
	SHA384 = "384"
	SHA512 = "512"
)

func RunPrintSHA() {

	var sha = flag.String("sha", SHA256, "int Digest for the string")
	flag.Parse()
	if argument := flag.Arg(0); argument != "" {
		switch *sha {
		case SHA256:
			fmt.Printf("SHA%s: %x\n", *sha, sha256.Sum256([]byte(argument)))
		case SHA384:
			fmt.Printf("SHA%s: %x\n", *sha, sha512.Sum384([]byte(argument)))
		case SHA512:
			fmt.Printf("SHA%s: %x\n", *sha, sha512.Sum512([]byte(argument)))
		default:
			fmt.Println("Unknown SHA type")
		}
	} else {
		fmt.Println("Please provide a string to hash")
	}
}
