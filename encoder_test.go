package morsegode

import "testing"

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want string
	}{
		{name: "single word", msg: "hello", want: ".... . .-.. .-.. ---"},
		{name: "uppercase is lowercased", msg: "HELLO", want: ".... . .-.. .-.. ---"},
		{name: "multiple words", msg: "hello world", want: ".... . .-.. .-.. --- / .-- --- .-. .-.. -.."},
		{name: "double space collapses", msg: "hello  world", want: ".... . .-.. .-.. --- / .-- --- .-. .-.. -.."},
		{name: "empty message", msg: "", want: ""},
	}

	for _, tt := range tests {
		got, err := AlphabeticEncoder.Encode(tt.msg)
		if err != nil {
			t.Errorf("%s: unexpected error: %v", tt.name, err)
			continue
		}
		if got != tt.want {
			t.Errorf("%s: Encode() = %q, want %q", tt.name, got, tt.want)
		}
	}
}

func TestEncodeInvalid(t *testing.T) {
	// Digits are not part of the alphabetic tree.
	if _, err := AlphabeticEncoder.Encode("hello 123"); err == nil {
		t.Errorf("Encode with unsupported characters should return an error")
	}
}

func TestEncodeInternationalPunctuation(t *testing.T) {
	msg := "hi!"
	want := ".... .. -.-.--"

	got, err := InternationalEncoder.Encode(msg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("Encode() = %q, want %q", got, want)
	}
}
