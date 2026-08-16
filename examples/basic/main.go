// Command basic demonstrates encoding and decoding text with the
// letters-only AlphabeticEncoder and AlphabeticDecoder.
package main

import (
	"fmt"
	"log"

	"github.com/TobyReay/morse-gode"
)

func main() {
	morseString := ".... . .-.. .-.. --- / .-- --- .-. .-.. -.."
	fmt.Println("Morse input:", morseString)

	decoded, err := morsegode.AlphabeticDecoder.Decode(morseString)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Decoded:", decoded)

	encoded, err := morsegode.AlphabeticEncoder.Encode(decoded)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Re-encoded:", encoded)
}
