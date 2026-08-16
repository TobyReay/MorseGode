package morsegode

// internationalTable holds the full International Morse code character set as
// defined by ITU-R M.1677-1: the 26 Latin letters, the 10 digits and the most
// common punctuation marks. Each key is a dot/dash sequence and each value is
// the character it represents.
var internationalTable = map[string]string{
	// Letters
	"a": ".-",
	"b": "-...",
	"c": "-.-.",
	"d": "-..",
	"e": ".",
	"f": "..-.",
	"g": "--.",
	"h": "....",
	"i": "..",
	"j": ".---",
	"k": "-.-",
	"l": ".-..",
	"m": "--",
	"n": "-.",
	"o": "---",
	"p": ".--.",
	"q": "--.-",
	"r": ".-.",
	"s": "...",
	"t": "-",
	"u": "..-",
	"v": "...-",
	"w": ".--",
	"x": "-..-",
	"y": "-.--",
	"z": "--..",

	// Digits
	"0": "-----",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",

	// Punctuation
	".": ".-.-.-",
	",": "--..--",
	"?": "..--..",
	"'": ".----.",
	"!": "-.-.--",
	"/": "-..-.",
	"(": "-.--.",
	")": "-.--.-",
	"&": ".-...",
	":": "---...",
	";": "-.-.-.",
	"=": "-...-",
	"+": ".-.-.",
	"-": "-....-",
	"_": "..--.-",
	`"`: ".-..-.",
	"$": "...-..-",
	"@": ".--.-",
}

// numericTable holds the International Morse code for the 10 digits, 0-9.
var numericTable = map[string]string{
	"0": "-----",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
}

// wabunTable holds the Wabun (Japanese telegraphic) code for the katakana
// syllabograms together with the punctuation and diacritics commonly used
// with them. Each key is a dot/dash sequence and each value is the kana it
// represents. The codes follow ITU-R M.1677-1.
var wabunTable = map[string]string{
	"ア": "--.--",
	"イ": ".-",
	"ウ": "..-",
	"エ": "-.---",
	"オ": ".-...",
	"カ": ".-..",
	"キ": "-.-..",
	"ク": "...-",
	"ケ": "-.--",
	"コ": "----",
	"サ": "-.-.-",
	"シ": "--.-.",
	"ス": "---.-",
	"セ": ".---.",
	"ソ": "---.",
	"タ": "-.",
	"チ": "..-.",
	"ツ": ".--.",
	"テ": ".-.--",
	"ト": "..-..",
	"ナ": ".-.",
	"ニ": "-.-.",
	"ヌ": "....",
	"ネ": "--.-",
	"ノ": "..--",
	"ハ": "-...",
	"ヒ": "--..-",
	"フ": "--..",
	"ヘ": ".",
	"ホ": "-..",
	"マ": "-..-",
	"ミ": "..-.-",
	"ム": "-",
	"メ": "-...-",
	"モ": "-..-.",
	"ヤ": ".--",
	"ユ": "-..--",
	"ヨ": "--",
	"ラ": "...",
	"リ": "--.",
	"ル": "-.--.-",
	"レ": "---",
	"ロ": ".-.-",
	"ワ": "-.-",
	"ヰ": ".-..-",
	"ヱ": ".--..",
	"ヲ": ".---",
	"ン": ".-.-.",

	// Punctuation and diacritics
	"。": ".-.-..",
	"、": ".-.-.-",
	"ー": ".--.-",
	"゛": "..",
	"゜": "..--.",
}

// InternationalMorseTree is a Morse tree for the full International Morse
// code character set (letters, digits and common punctuation) as defined by
// ITU-R M.1677-1.
var InternationalMorseTree = mustNewTree(internationalTable)

// NumericMorseTree is a Morse tree for the 10 digits, 0-9, using
// International Morse code symbols.
var NumericMorseTree = mustNewTree(numericTable)

// WabunMorseTree is a Morse tree for the Wabun (Japanese telegraphic) code,
// covering the katakana syllabograms and their punctuation and diacritics.
var WabunMorseTree = mustNewTree(wabunTable)

// mustNewTree builds a tree from table using the standard International Morse
// symbols ('.' for dot, '-' for dash). It panics if the table is invalid,
// which cannot happen for the built-in tables.
func mustNewTree(table map[string]string) *GeneralMorseTree {
	tree, err := NewGeneralMorseTree('.', '-', table)
	if err != nil {
		panic(err)
	}
	return tree
}
