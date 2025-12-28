package internal

// CodepointTransformation provides 1:1 codepoint mapping for visual rendering.
//
// This interface maps each codepoint in the input to another codepoint (or the same).
// The transformation must be 1:1, meaning it preserves text length. This is used
// for effects like password masking where each character is replaced with a dot.
//
// This is a port of androidx.compose.foundation.text.input.internal.CodepointTransformation.
type CodepointTransformation interface {
	// Transform maps a codepoint at the given index to another codepoint.
	// The codepoint is represented as a rune in Go.
	Transform(codepointIndex int, codepoint rune) rune
}

// CodepointTransformationFunc is a function type that implements CodepointTransformation.
type CodepointTransformationFunc func(codepointIndex int, codepoint rune) rune

// Transform implements CodepointTransformation.
func (f CodepointTransformationFunc) Transform(codepointIndex int, codepoint rune) rune {
	return f(codepointIndex, codepoint)
}

// MaskCodepointTransformation masks all characters with the given rune.
type MaskCodepointTransformation struct {
	Mask rune
}

// Transform implements CodepointTransformation.
func (m *MaskCodepointTransformation) Transform(codepointIndex int, codepoint rune) rune {
	return m.Mask
}

// NewMaskCodepointTransformation creates a CodepointTransformation that masks
// all characters with the given rune.
func NewMaskCodepointTransformation(mask rune) CodepointTransformation {
	return &MaskCodepointTransformation{Mask: mask}
}

// PasswordMaskTransformation masks all characters with the password bullet (●).
var PasswordMaskTransformation = NewMaskCodepointTransformation('●')

// DotMaskTransformation masks all characters with a dot (•).
var DotMaskTransformation = NewMaskCodepointTransformation('•')

// AsteriskMaskTransformation masks all characters with an asterisk (*).
var AsteriskMaskTransformation = NewMaskCodepointTransformation('*')

// singleLineCodepointTransformation converts newlines to spaces.
type singleLineCodepointTransformation struct{}

const (
	lineFeed       = '\n'
	carriageReturn = '\r'
	whitespace     = ' '
	zeroWidthSpace = '\uFEFF'
)

// Transform implements CodepointTransformation.
func (s *singleLineCodepointTransformation) Transform(codepointIndex int, codepoint rune) rune {
	if codepoint == lineFeed {
		return whitespace
	}
	if codepoint == carriageReturn {
		return zeroWidthSpace
	}
	return codepoint
}

// SingleLineCodepointTransformation converts newlines to spaces and
// carriage returns to zero-width spaces.
// This ensures content appears as a single line while preserving text length.
var SingleLineCodepointTransformation CodepointTransformation = &singleLineCodepointTransformation{}

// IdentityCodepointTransformation returns each codepoint unchanged.
var IdentityCodepointTransformation = CodepointTransformationFunc(func(codepointIndex int, codepoint rune) rune {
	return codepoint
})

// ApplyCodepointTransformation applies a CodepointTransformation to a string.
// Returns the transformed string.
func ApplyCodepointTransformation(text string, transform CodepointTransformation) string {
	if transform == nil {
		return text
	}

	runes := []rune(text)
	changed := false

	for i, r := range runes {
		newR := transform.Transform(i, r)
		if newR != r {
			runes[i] = newR
			changed = true
		}
	}

	if !changed {
		return text
	}
	return string(runes)
}
