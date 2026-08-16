package morsegode

import "testing"

func TestDecode(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want string
	}{
		{name: "single word", msg: ".... . .-.. .-.. ---", want: "hello"},
		{name: "multiple words", msg: ".... . .-.. .-.. --- / .-- --- .-. .-.. -..", want: "hello world"},
		{name: "extra whitespace around separator", msg: " .... . .-.. .-.. --- / .-- --- .-. .-.. -.. ", want: "hello world"},
		{name: "empty message", msg: "", want: ""},
		{name: "whitespace only", msg: "   ", want: ""},
		{name: "sos", msg: "... --- ...", want: "sos"},
	}

	for _, tt := range tests {
		got, err := AlphabeticDecoder.Decode(tt.msg)
		if err != nil {
			t.Errorf("%s: unexpected error: %v", tt.name, err)
			continue
		}
		if got != tt.want {
			t.Errorf("%s: Decode() = %q, want %q", tt.name, got, tt.want)
		}
	}
}

func TestDecodeInvalid(t *testing.T) {
	// "....-" walks off the alphabetic tree.
	if _, err := AlphabeticDecoder.Decode(".... . ....-"); err == nil {
		t.Errorf("Decode with an invalid sequence should return an error")
	}

	// "x" is not a dot or dash.
	if _, err := AlphabeticDecoder.Decode(".... . x"); err == nil {
		t.Errorf("Decode with an unknown symbol should return an error")
	}
}

func TestDecodeRoundTripWithWords(t *testing.T) {
	msg := ".... . .-.. .-.. --- / .. / .-.. --- ...- . / -- --- .-. ... ."
	want := "hello i love morse"

	got, err := AlphabeticDecoder.Decode(msg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("Decode() = %q, want %q", got, want)
	}
}
