package morsegode

import "strings"

// Decoder translates Morse code into plain text.
type Decoder interface {
	// Decode returns the plain text representation of Msg.
	Decode(Msg string) (string, error)
}

// GeneralDecoder is a Decoder backed by a MorseTree. It splits the input
// Morse code into words and letters, decodes each letter by walking the tree,
// and joins the results together using the configured separators.
type GeneralDecoder struct {
	// Tree is the Morse tree used to decode symbols.
	Tree MorseTree

	// WordSep separates words in the input Morse code (for example "/").
	WordSep string

	// WordJoin separates words in the decoded output (for example " ").
	WordJoin string

	// LetterSep separates letters within a word in the input Morse code
	// (for example " ").
	LetterSep string
}

// Decode returns the plain text representation of Msg. An empty message
// decodes to an empty string. Malformed Morse code (a sequence that walks off
// the tree, or one containing symbols the tree does not know) results in an
// error.
func (decoder GeneralDecoder) Decode(Msg string) (string, error) {
	if len(Msg) <= 0 {
		return "", nil
	}

	words := strings.Split(Msg, decoder.WordSep)
	decodedMessage := make([]string, 0)

	for _, word := range words {
		word = strings.TrimSpace(word)
		if word == "" {
			continue
		}

		letters := strings.Split(word, decoder.LetterSep)
		decodedLetter, err := decoder.Tree.DecodeWord(letters)
		if err != nil {
			return "", err
		}

		decodedMessage = append(decodedMessage, decodedLetter)
	}

	return strings.Join(decodedMessage, decoder.WordJoin), nil
}
