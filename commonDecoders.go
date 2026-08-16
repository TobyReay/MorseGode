package morsegode

// AlphabeticDecoder decodes International Morse code for the 26 letters of
// the Latin alphabet. Words are separated by "/" in the input and " " in the
// output; letters within a word are separated by " ".
var AlphabeticDecoder = GeneralDecoder{
	Tree: AlphabeticMorseTree,

	WordSep:   "/",
	WordJoin:  " ",
	LetterSep: " ",
}

// InternationalDecoder decodes the full International Morse code character
// set (letters, digits and common punctuation). Words are separated by "/" in
// the input and " " in the output; letters within a word are separated by
// " ".
var InternationalDecoder = GeneralDecoder{
	Tree: InternationalMorseTree,

	WordSep:   "/",
	WordJoin:  " ",
	LetterSep: " ",
}

// NumericDecoder decodes Morse code for the 10 digits, 0-9. Words are
// separated by "/" in the input and " " in the output; digits within a word
// are separated by " ".
var NumericDecoder = GeneralDecoder{
	Tree: NumericMorseTree,

	WordSep:   "/",
	WordJoin:  " ",
	LetterSep: " ",
}

// WabunDecoder decodes the Wabun (Japanese telegraphic) code for katakana
// and its punctuation. Words are separated by "/" in the input and " " in the
// output; kana within a word are separated by " ".
var WabunDecoder = GeneralDecoder{
	Tree: WabunMorseTree,

	WordSep:   "/",
	WordJoin:  " ",
	LetterSep: " ",
}
