package morsegode

import "testing"

func TestAlphabeticRoundTrip(t *testing.T) {
	msg := "hello i love morse code"

	encoded, err := AlphabeticEncoder.Encode(msg)
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}

	decoded, err := AlphabeticDecoder.Decode(encoded)
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}

	if decoded != msg {
		t.Errorf("round trip = %q, want %q", decoded, msg)
	}
}

func TestInternationalEncodeGolden(t *testing.T) {
	tests := []struct {
		msg  string
		want string
	}{
		{msg: "sos", want: "... --- ..."},
		{msg: "hello world 123", want: ".... . .-.. .-.. --- / .-- --- .-. .-.. -.. / .---- ..--- ...--"},
		{msg: "save 100 dollars", want: "... .- ...- . / .---- ----- ----- / -.. --- .-.. .-.. .- .-. ..."},
	}

	for _, tt := range tests {
		got, err := InternationalEncoder.Encode(tt.msg)
		if err != nil {
			t.Errorf("Encode(%q): %v", tt.msg, err)
			continue
		}
		if got != tt.want {
			t.Errorf("Encode(%q) = %q, want %q", tt.msg, got, tt.want)
		}
	}
}

func TestInternationalRoundTrip(t *testing.T) {
	msg := "hello world 123 !"

	encoded, err := InternationalEncoder.Encode(msg)
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}

	decoded, err := InternationalDecoder.Decode(encoded)
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}

	if decoded != msg {
		t.Errorf("round trip = %q, want %q", decoded, msg)
	}
}

func TestInternationalDigits(t *testing.T) {
	encoded, err := InternationalEncoder.Encode("0123456789")
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}
	want := "----- .---- ..--- ...-- ....- ..... -.... --... ---.. ----."
	if encoded != want {
		t.Errorf("Encode() = %q, want %q", encoded, want)
	}

	decoded, err := InternationalDecoder.Decode(want)
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}
	if decoded != "0123456789" {
		t.Errorf("Decode() = %q, want %q", decoded, "0123456789")
	}
}

func TestInternationalPunctuation(t *testing.T) {
	msg := "where are you? i am here!"

	encoded, err := InternationalEncoder.Encode(msg)
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}

	decoded, err := InternationalDecoder.Decode(encoded)
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}

	if decoded != msg {
		t.Errorf("round trip = %q, want %q", decoded, msg)
	}
}

func TestNumericRoundTrip(t *testing.T) {
	msg := "0123456789"

	encoded, err := NumericEncoder.Encode(msg)
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}

	decoded, err := NumericDecoder.Decode(encoded)
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}

	if decoded != msg {
		t.Errorf("round trip = %q, want %q", decoded, msg)
	}
}

func TestNumericEncode(t *testing.T) {
	got, err := NumericEncoder.Encode("123")
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}
	want := ".---- ..--- ...--"
	if got != want {
		t.Errorf("Encode() = %q, want %q", got, want)
	}
}

func TestNumericDecode(t *testing.T) {
	got, err := NumericDecoder.Decode(".---- ..--- ...-- / ....- ..... -....")
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}
	want := "123 456"
	if got != want {
		t.Errorf("Decode() = %q, want %q", got, want)
	}
}

func TestWabunRoundTrip(t *testing.T) {
	msg := "モールス"

	encoded, err := WabunEncoder.Encode(msg)
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}
	want := "-..-. .--.- -.--.- ---.-"
	if encoded != want {
		t.Errorf("Encode() = %q, want %q", encoded, want)
	}

	decoded, err := WabunDecoder.Decode(encoded)
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}
	if decoded != msg {
		t.Errorf("round trip = %q, want %q", decoded, msg)
	}
}

func TestWabunKana(t *testing.T) {
	tests := []struct {
		kana string
		want string
	}{
		{kana: "アイウ", want: "--.-- .- ..-"},
		{kana: "トウキヨウ", want: "..-.. ..- -.-.. -- ..-"},
		{kana: "ン", want: ".-.-."},
		{kana: "ヘ", want: "."},
		{kana: "ム", want: "-"},
	}

	for _, tt := range tests {
		encoded, err := WabunEncoder.Encode(tt.kana)
		if err != nil {
			t.Errorf("Encode(%q): %v", tt.kana, err)
			continue
		}
		if encoded != tt.want {
			t.Errorf("Encode(%q) = %q, want %q", tt.kana, encoded, tt.want)
		}
	}
}

func TestWabunPunctuation(t *testing.T) {
	msg := "ヘイ。ハロー、"

	encoded, err := WabunEncoder.Encode(msg)
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}

	decoded, err := WabunDecoder.Decode(encoded)
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}

	if decoded != msg {
		t.Errorf("round trip = %q, want %q", decoded, msg)
	}
}
