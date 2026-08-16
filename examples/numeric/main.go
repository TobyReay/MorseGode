// Command numeric demonstrates encoding and decoding digits with the
// numeric-only NumericEncoder and NumericDecoder.
package main

import (
	"fmt"
	"log"

	"github.com/TobyReay/morse-gode"
)

func main() {
	digits := "0123456789"
	fmt.Println("Digits input:", digits)

	encoded, err := morsegode.NumericEncoder.Encode(digits)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Encoded:", encoded)

	decoded, err := morsegode.NumericDecoder.Decode(encoded)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Decoded:", decoded)
}
