package morsegode

import "strings"

// Encoder translates plain text into its Morse code representation.
type Encoder interface {
	// Encode returns the Morse code representation of cypherText.
	Encode(cypherText string) (string, error)
}

// GeneralEncoder is an Encoder backed by a MorseTree. It splits the input
// text into words and letters, encodes each value by searching the tree, and
// joins the results together using the configured separators.
type GeneralEncoder struct {
	// Tree is the Morse tree used to encode values.
	Tree MorseTree

	// WordSplit separates words in the input text (for example " ").
	WordSplit string

	// WordJoin separates words in the encoded output (for example " / ").
	WordJoin string

	// LetterSplit separates individual values within a word in the input text
	// (for example "" to split by rune).
	LetterSplit string

	// LetterJoin separates letters within a word in the encoded output
	// (for example " ").
	LetterJoin string
}

// Encode returns the Morse code representation of cypherText. The input is
// lower-cased before encoding, so upper and lower case letters encode the
// same way. Any value that is not present in the tree results in an error.
func (encoder GeneralEncoder) Encode(cypherText string) (string, error) {
	cypherText = strings.ToLower(cypherText)
	textParts := strings.Split(cypherText, encoder.WordSplit)

	encodedStrings := make([]string, 0)
	for _, part := range textParts {
		if part == "" {
			continue
		}

		letters := strings.Split(part, encoder.LetterSplit)
		encodedWord, encoderErr := encoder.Tree.EncodeWord(letters, encoder.LetterJoin)
		if encoderErr != nil {
			return "", encoderErr
		}

		encodedStrings = append(encodedStrings, encodedWord)
	}

	return strings.Join(encodedStrings, encoder.WordJoin), nil
}
