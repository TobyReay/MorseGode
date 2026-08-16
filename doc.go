// Package morsegode provides encoders and decoders for Morse code built on a
// binary tree representation.
//
// Each Morse code character (a letter, digit, punctuation mark, or kana) is
// stored at a node of a binary tree. Following a dot ('.') descends to the
// left child and following a dash ('-') descends to the right child, so the
// path of dots and dashes that reaches a node is exactly its Morse code.
//
// # Decoding
//
// A GeneralDecoder splits an incoming Morse message into words and letters,
// walks the tree for each letter, and joins the decoded values back together:
//
//	decoded, err := AlphabeticDecoder.Decode(".... . .-.. .-.. ---")
//	// decoded == "hello"
//
// # Encoding
//
// A GeneralEncoder splits plain text into words and letters, searches the
// tree for each value, and joins the resulting code sequences:
//
//	encoded, err := AlphabeticEncoder.Encode("hello")
//	// encoded == ".... . .-.. .-.. ---"
//
// # Built-in codecs
//
// The package ships with several ready-made codecs:
//
//   - AlphabeticDecoder / AlphabeticEncoder: the 26 Latin letters.
//   - InternationalDecoder / InternationalEncoder: letters, digits and common
//     punctuation per ITU-R M.1677-1.
//   - NumericDecoder / NumericEncoder: the 10 digits.
//   - WabunDecoder / WabunEncoder: Japanese katakana telegraphic code.
//
// # Custom codecs
//
// New code tables can be built with NewGeneralMorseTree, which constructs a
// tree from a map of dot/dash sequences to the values they represent, and
// wrapped in a GeneralDecoder and GeneralEncoder.
package morsegode
