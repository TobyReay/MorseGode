package morsegode

import "testing"

func TestSearchTreeLetters(t *testing.T) {
	expected := map[string]string{
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
	}

	for letter, morse := range expected {
		got, err := AlphabeticMorseTree.SearchTree(letter)
		if err != nil {
			t.Errorf("SearchTree(%q) unexpected error: %v", letter, err)
			continue
		}
		if got != morse {
			t.Errorf("SearchTree(%q) = %q, want %q", letter, got, morse)
		}
	}
}

func TestSearchTreePunctuation(t *testing.T) {
	expected := map[string]string{
		".": ".-.-.-",
		",": "--..--",
		"?": "..--..",
		"!": "-.-.--",
		"@": ".--.-",
		"$": "...-..-",
	}

	for symbol, morse := range expected {
		got, err := InternationalMorseTree.SearchTree(symbol)
		if err != nil {
			t.Errorf("SearchTree(%q) unexpected error: %v", symbol, err)
			continue
		}
		if got != morse {
			t.Errorf("SearchTree(%q) = %q, want %q", symbol, got, morse)
		}
	}
}

func TestSearchTreeNotFound(t *testing.T) {
	for _, tree := range []MorseTree{AlphabeticMorseTree, InternationalMorseTree, NumericMorseTree} {
		if _, err := tree.SearchTree(""); err == nil {
			t.Errorf("SearchTree(\"\") should return an error")
		}
	}

	if _, err := AlphabeticMorseTree.SearchTree("1"); err == nil {
		t.Errorf("SearchTree(\"1\") on alphabetic tree should return an error")
	}

	if _, err := NumericMorseTree.SearchTree("a"); err == nil {
		t.Errorf("SearchTree(\"a\") on numeric tree should return an error")
	}
}

func TestTravelTree(t *testing.T) {
	expected := map[string]string{
		".":      "e",
		"-":      "t",
		"...":    "s",
		"...-":   "v",
		".-":     "a",
		"....":   "h",
		"---":    "o",
		"":       "",
		"-----":  "0",
		".----":  "1",
		".-.-.-": ".",
		"--..--": ",",
		"..--..": "?",
	}

	for morse, want := range expected {
		var tree MorseTree = AlphabeticMorseTree
		if want == "0" || want == "1" {
			tree = NumericMorseTree
		}
		if want == "." || want == "," || want == "?" {
			tree = InternationalMorseTree
		}

		got, err := tree.TravelTree([]byte(morse))
		if err != nil {
			t.Errorf("TravelTree(%q) unexpected error: %v", morse, err)
			continue
		}
		if got != want {
			t.Errorf("TravelTree(%q) = %q, want %q", morse, got, want)
		}
	}
}

func TestTravelTreeInvalid(t *testing.T) {
	tree := AlphabeticMorseTree

	// Walks off the tree: h (....) has no children.
	if _, err := tree.TravelTree([]byte("....-")); err == nil {
		t.Errorf("TravelTree(\"....-\") should return an error")
	}

	// Unknown symbol.
	if _, err := tree.TravelTree([]byte("..x.")); err == nil {
		t.Errorf("TravelTree(\"..x.\") should return an error")
	}

	// Empty tree.
	emptyTree := GeneralMorseTree{}
	if _, err := emptyTree.TravelTree([]byte(".")); err == nil {
		t.Errorf("TravelTree on empty tree should return an error")
	}
}

func TestDecodeWord(t *testing.T) {
	tests := []struct {
		name string
		seq  []string
		want string
	}{
		{name: "hello", seq: []string{"....", ".", ".-..", ".-..", "---"}, want: "hello"},
		{name: "sos", seq: []string{"...", "---", "..."}, want: "sos"},
		{name: "empty tokens skipped", seq: []string{"....", "", "."}, want: "he"},
		{name: "empty sequence", seq: nil, want: ""},
	}

	for _, tt := range tests {
		got, err := AlphabeticMorseTree.DecodeWord(tt.seq)
		if err != nil {
			t.Errorf("%s: unexpected error: %v", tt.name, err)
			continue
		}
		if got != tt.want {
			t.Errorf("%s: DecodeWord() = %q, want %q", tt.name, got, tt.want)
		}
	}
}

func TestDecodeWordInvalid(t *testing.T) {
	// "....-" walks off the alphabetic tree.
	if _, err := AlphabeticMorseTree.DecodeWord([]string{"....", "....-"}); err == nil {
		t.Errorf("DecodeWord with an invalid sequence should return an error")
	}
}

func TestEncodeWord(t *testing.T) {
	tests := []struct {
		name    string
		letters []string
		join    string
		want    string
	}{
		{name: "hello", letters: []string{"h", "e", "l", "l", "o"}, join: " ", want: ".... . .-.. .-.. ---"},
		{name: "sos", letters: []string{"s", "o", "s"}, join: " ", want: "... --- ..."},
		{name: "empty letters skipped", letters: []string{"h", "", "e"}, join: " ", want: ".... ."},
		{name: "custom join", letters: []string{"a", "b"}, join: " | ", want: ".- | -..."},
	}

	for _, tt := range tests {
		got, err := AlphabeticMorseTree.EncodeWord(tt.letters, tt.join)
		if err != nil {
			t.Errorf("%s: unexpected error: %v", tt.name, err)
			continue
		}
		if got != tt.want {
			t.Errorf("%s: EncodeWord() = %q, want %q", tt.name, got, tt.want)
		}
	}
}

func TestEncodeWordInvalid(t *testing.T) {
	// "1" is not in the alphabetic tree.
	if _, err := AlphabeticMorseTree.EncodeWord([]string{"h", "1"}, " "); err == nil {
		t.Errorf("EncodeWord with an unsupported letter should return an error")
	}
}

func TestNewGeneralMorseTree(t *testing.T) {
	table := map[string]string{
		"a": ".-",
		"b": "-...",
	}

	tree, err := NewGeneralMorseTree('.', '-', table)
	if err != nil {
		t.Fatalf("NewGeneralMorseTree unexpected error: %v", err)
	}

	if got, err := tree.SearchTree("a"); err != nil || got != ".-" {
		t.Errorf("SearchTree(\"a\") = %q, %v; want %q, nil", got, err, ".-")
	}

	if got, err := tree.TravelTree([]byte("-...")); err != nil || got != "b" {
		t.Errorf("TravelTree(\"-...\") = %q, %v; want %q, nil", got, err, "b")
	}

	if _, err := tree.SearchTree("c"); err == nil {
		t.Errorf("SearchTree(\"c\") should return an error")
	}
}

func TestNewGeneralMorseTreeInvalid(t *testing.T) {
	// Empty path.
	if _, err := NewGeneralMorseTree('.', '-', map[string]string{"": "x"}); err == nil {
		t.Errorf("empty path should return an error")
	}

	// Symbol that is neither left nor right.
	if _, err := NewGeneralMorseTree('.', '-', map[string]string{"..x..": "x"}); err == nil {
		t.Errorf("invalid symbol should return an error")
	}
}

func TestGeneralMorseTreeNodeItem(t *testing.T) {
	node := GeneralMorseTreeNode{letter: "e"}
	if got := node.Item(); got != "e" {
		t.Errorf("Item() = %v, want %q", got, "e")
	}

	if node.Left() != nil || node.Right() != nil {
		t.Errorf("children should be nil by default")
	}
}
