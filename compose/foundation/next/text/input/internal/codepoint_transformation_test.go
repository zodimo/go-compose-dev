package internal

import (
	"testing"
)

func TestCodepointTransformation_Mask(t *testing.T) {
	transform := NewMaskCodepointTransformation('●')

	tests := []struct {
		input rune
		want  rune
	}{
		{'H', '●'},
		{'e', '●'},
		{'1', '●'},
		{'!', '●'},
	}

	for _, tt := range tests {
		got := transform.Transform(0, tt.input)
		if got != tt.want {
			t.Errorf("Transform(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestCodepointTransformation_SingleLine(t *testing.T) {
	transform := SingleLineCodepointTransformation

	tests := []struct {
		input rune
		want  rune
	}{
		{'H', 'H'},
		{'\n', ' '},      // Newline becomes space
		{'\r', '\uFEFF'}, // CR becomes zero-width space
		{'a', 'a'},
	}

	for _, tt := range tests {
		got := transform.Transform(0, tt.input)
		if got != tt.want {
			t.Errorf("Transform(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestCodepointTransformation_Identity(t *testing.T) {
	transform := IdentityCodepointTransformation

	inputs := []rune{'H', 'e', 'l', 'l', 'o', '\n', '世', '界'}
	for _, r := range inputs {
		got := transform.Transform(0, r)
		if got != r {
			t.Errorf("Identity Transform(%q) = %q, want %q", r, got, r)
		}
	}
}

func TestApplyCodepointTransformation(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		transform CodepointTransformation
		want      string
	}{
		{"password mask", "Hello", PasswordMaskTransformation, "●●●●●"},
		{"single line", "Hello\nWorld", SingleLineCodepointTransformation, "Hello World"},
		{"identity", "Hello", IdentityCodepointTransformation, "Hello"},
		{"nil transform", "Hello", nil, "Hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ApplyCodepointTransformation(tt.text, tt.transform)
			if got != tt.want {
				t.Errorf("got '%s', want '%s'", got, tt.want)
			}
		})
	}
}

func TestApplyCodepointTransformation_Unicode(t *testing.T) {
	got := ApplyCodepointTransformation("Hello 世界", PasswordMaskTransformation)
	// 8 characters -> 8 bullets
	if got != "●●●●●●●●" {
		t.Errorf("got '%s', expected 8 bullets", got)
	}
}

func TestCodepointTransformation_Func(t *testing.T) {
	// Custom transformation: rot13
	rot13 := CodepointTransformationFunc(func(index int, r rune) rune {
		if r >= 'a' && r <= 'z' {
			return 'a' + (r-'a'+13)%26
		}
		if r >= 'A' && r <= 'Z' {
			return 'A' + (r-'A'+13)%26
		}
		return r
	})

	got := ApplyCodepointTransformation("Hello", rot13)
	if got != "Uryyb" {
		t.Errorf("rot13('Hello') = '%s', want 'Uryyb'", got)
	}
}
