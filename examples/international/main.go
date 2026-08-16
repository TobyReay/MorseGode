// Command international demonstrates the full International Morse code
// character set: letters, digits and punctuation.
package main

import (
	"fmt"
	"log"

	"github.com/TobyReay/morse-gode"
)

func main() {
	msg := "HELLO WORLD 123!"
	fmt.Println("Text input:", msg)

	encoded, err := morsegode.InternationalEncoder.Encode(msg)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Encoded:", encoded)

	decoded, err := morsegode.InternationalDecoder.Decode(encoded)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Decoded:", decoded)
}
