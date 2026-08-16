# morsegode

A standalone Go package for encoding and decoding Morse code, built on a
binary tree representation. Each character (letter, digit, punctuation mark
or kana) is stored at a node of a binary tree; a dot (`'.'`) descends to the
left child and a dash (`'-'`) descends to the right child, so the path to a
node *is* its Morse code.

## Install

```sh
go get github.com/TobyReay/morse-gode
```

## Quick start

```go
package main

import (
	"fmt"

	"github.com/TobyReay/morse-gode"
)

func main() {
	// Decode
	decoded, err := morsegode.AlphabeticDecoder.Decode(".... . .-.. .-.. --- / .-- --- .-. .-.. -..")
	if err != nil {
		panic(err)
	}
	fmt.Println(decoded) // hello world

	// Encode
	encoded, err := morsegode.InternationalEncoder.Encode("HELLO WORLD 123!")
	if err != nil {
		panic(err)
	}
	fmt.Println(encoded) // .... . .-.. .-.. --- / .-- --- .-. .-.. -.. / .---- ..--- ...-- -.-.--
}
```

## Built-in codecs

| Codec                               | Character set                                              |
| ----------------------------------- | ---------------------------------------------------------- |
| `AlphabeticDecoder` / `AlphabeticEncoder`   | The 26 letters of the Latin alphabet (A-Z)          |
| `InternationalDecoder` / `InternationalEncoder` | Letters, digits 0-9 and common punctuation per ITU-R M.1677-1 |
| `NumericDecoder` / `NumericEncoder` | The 10 digits (0-9)                                        |
| `WabunDecoder` / `WabunEncoder`     | Japanese katakana telegraphic code (Wabun)                 |

Each decoder splits incoming Morse code into words and letters using its
configured separators (`WordSep` `/` and `LetterSep` space by default), walks
the tree for every letter, and joins the decoded text with `WordJoin`. Each
encoder does the reverse: it splits text into words and letters, looks each
value up in the tree, and joins the codes with `LetterJoin` and `WordJoin`.

All built-in encoders are case-insensitive: input is lower-cased before
encoding, matching how Morse code treats upper and lower case letters.

## Custom codecs

Build your own tree from a code table with `NewGeneralMorseTree` and wrap it
in a `GeneralDecoder` and `GeneralEncoder`:

```go
tree, err := morsegode.NewGeneralMorseTree('.', '-', map[string]string{
	"a": ".-",
	"b": "-...",
	"c": "-.-.",
})
if err != nil {
	panic(err)
}

encoder := morsegode.GeneralEncoder{
	Tree:        tree,
	WordSplit:   " ",
	WordJoin:    " / ",
	LetterSplit: "",
	LetterJoin:  " ",
}

decoder := morsegode.GeneralDecoder{
	Tree:      tree,
	WordSep:   "/",
	WordJoin:  " ",
	LetterSep: " ",
}
```

## API

- `NewGeneralMorseTree(left, right byte, table map[string]string) (*GeneralMorseTree, error)` — build a tree from a value-to-code table.
- `MorseTree` — tree interface with `TravelTree`, `SearchTree`, `DecodeWord` and `EncodeWord`.
- `GeneralMorseTree` / `GeneralMorseTreeNode` — default tree and node types.
- `Encoder` / `GeneralEncoder` — text-to-Morse; `Encode(text string) (string, error)`.
- `Decoder` / `GeneralDecoder` — Morse-to-text; `Decode(msg string) (string, error)`.

## Examples

Runnable examples live in the `examples/` directory:

```sh
go run ./examples/basic
go run ./examples/international
go run ./examples/numeric
go run ./examples/wabun
```

## Testing

```sh
go test ./...
```

## License

MIT