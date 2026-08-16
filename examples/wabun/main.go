// Command wabun demonstrates the Wabun (Japanese telegraphic) code for
// katakana with the WabunEncoder and WabunDecoder.
package main

import (
	"fmt"
	"log"

	"github.com/TobyReay/morse-gode"
)

func main() {
	msg := "モールス"
	fmt.Println("Text input:", msg)

	encoded, err := morsegode.WabunEncoder.Encode(msg)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Encoded:", encoded)

	decoded, err := morsegode.WabunDecoder.Decode(encoded)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Decoded:", decoded)
}
