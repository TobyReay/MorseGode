package morsegode

// AlphabeticEncoder encodes plain text into International Morse code using
// the 26 letters of the Latin alphabet. Words are separated by " " in the
// input and " / " in the output; letters within a word are joined with " ".
var AlphabeticEncoder = GeneralEncoder{
	Tree: AlphabeticMorseTree,

	WordSplit: " ",
	WordJoin:  " / ",

	LetterSplit: "",
	LetterJoin:  " ",
}

// InternationalEncoder encodes plain text into the full International Morse
// code character set (letters, digits and common punctuation). Words are
// separated by " " in the input and " / " in the output; letters within a
// word are joined with " ".
var InternationalEncoder = GeneralEncoder{
	Tree: InternationalMorseTree,

	WordSplit: " ",
	WordJoin:  " / ",

	LetterSplit: "",
	LetterJoin:  " ",
}

// NumericEncoder encodes a sequence of digits into Morse code. Digits are
// split on "" in the input and joined with " " in the output.
var NumericEncoder = GeneralEncoder{
	Tree: NumericMorseTree,

	WordSplit: " ",
	WordJoin:  " / ",

	LetterSplit: "",
	LetterJoin:  " ",
}

// WabunEncoder encodes plain text into the Wabun (Japanese telegraphic) code
// for katakana and its punctuation. Words are separated by " " in the input
// and " / " in the output; kana within a word are joined with " ".
var WabunEncoder = GeneralEncoder{
	Tree: WabunMorseTree,

	WordSplit: " ",
	WordJoin:  " / ",

	LetterSplit: "",
	LetterJoin:  " ",
}
