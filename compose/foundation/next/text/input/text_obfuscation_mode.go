package input

// TextObfuscationMode defines how text is obscured in secure fields.
//
// This is used for password fields and other sensitive input where
// the actual characters should be hidden from view.
//
// This is a port of androidx.compose.foundation.text.input.TextObfuscationMode.
type TextObfuscationMode int

const (
	// TextObfuscationModeVisible shows all content.
	// Use this when implementing a "show password" toggle.
	TextObfuscationModeVisible TextObfuscationMode = iota

	// TextObfuscationModeRevealLastTyped briefly reveals the last typed character
	// before masking it. This is the typical behavior on mobile devices.
	// Note: On Android, this also respects the system TEXT_SHOW_PASSWORD setting.
	TextObfuscationModeRevealLastTyped

	// TextObfuscationModeHidden masks all characters immediately.
	// This is more secure but less user-friendly.
	TextObfuscationModeHidden
)

// String returns a string representation of the TextObfuscationMode.
func (m TextObfuscationMode) String() string {
	switch m {
	case TextObfuscationModeVisible:
		return "Visible"
	case TextObfuscationModeRevealLastTyped:
		return "RevealLastTyped"
	case TextObfuscationModeHidden:
		return "Hidden"
	default:
		return "Unknown"
	}
}

// TextObfuscationModeDefault is the default obfuscation mode.
// This reveals the last typed character briefly before masking.
var TextObfuscationModeDefault = TextObfuscationModeRevealLastTyped
