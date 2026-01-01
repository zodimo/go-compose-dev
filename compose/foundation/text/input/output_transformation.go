package input

// OutputTransformation transforms text for visual presentation only.
//
// Output transformations modify how text is DISPLAYED without affecting the
// underlying TextFieldState. They are applied every frame and changes are
// discarded after rendering.
//
// Common use cases:
//   - Adding prefixes/suffixes (e.g., currency symbols)
//   - Formatting (e.g., phone numbers with dashes)
//   - Visual decorations
//
// The selection and cursor positions are automatically mapped through the
// transformation, so users can still edit the underlying text normally.
//
// This is a port of androidx.compose.foundation.text.input.OutputTransformation.
type OutputTransformation interface {
	// TransformOutput modifies the buffer for display purposes.
	// Changes do not affect the underlying TextFieldState.
	TransformOutput(buffer *TextFieldBuffer)
}

// OutputTransformationFunc is a function type that implements OutputTransformation.
type OutputTransformationFunc func(buffer *TextFieldBuffer)

// TransformOutput implements OutputTransformation.
func (f OutputTransformationFunc) TransformOutput(buffer *TextFieldBuffer) {
	f(buffer)
}

// PrefixOutputTransformation adds a prefix to the displayed text.
type PrefixOutputTransformation struct {
	Prefix string
}

// TransformOutput implements OutputTransformation.
func (p *PrefixOutputTransformation) TransformOutput(buffer *TextFieldBuffer) {
	buffer.Insert(0, p.Prefix)
}

// PrefixTransformation creates an OutputTransformation that adds a prefix.
func PrefixTransformation(prefix string) OutputTransformation {
	return &PrefixOutputTransformation{Prefix: prefix}
}

// SuffixOutputTransformation adds a suffix to the displayed text.
type SuffixOutputTransformation struct {
	Suffix string
}

// TransformOutput implements OutputTransformation.
func (s *SuffixOutputTransformation) TransformOutput(buffer *TextFieldBuffer) {
	buffer.Append(s.Suffix)
}

// SuffixTransformation creates an OutputTransformation that adds a suffix.
func SuffixTransformation(suffix string) OutputTransformation {
	return &SuffixOutputTransformation{Suffix: suffix}
}

// ChainedOutputTransformation combines two output transformations.
type ChainedOutputTransformation struct {
	First  OutputTransformation
	Second OutputTransformation
}

// TransformOutput implements OutputTransformation.
func (c *ChainedOutputTransformation) TransformOutput(buffer *TextFieldBuffer) {
	c.First.TransformOutput(buffer)
	c.Second.TransformOutput(buffer)
}

// ChainOutputTransformations creates an OutputTransformation that runs
// first then second.
func ChainOutputTransformations(first, second OutputTransformation) OutputTransformation {
	if first == nil {
		return second
	}
	if second == nil {
		return first
	}
	return &ChainedOutputTransformation{First: first, Second: second}
}

// PhoneNumberOutputTransformation formats text as a US phone number.
// This is an example transformation that formats "1234567890" as "(123) 456-7890".
type PhoneNumberOutputTransformation struct{}

// TransformOutput implements OutputTransformation.
func (p *PhoneNumberOutputTransformation) TransformOutput(buffer *TextFieldBuffer) {
	text := buffer.String()
	if len(text) == 0 {
		return
	}

	// Extract only digits
	var digits []rune
	for _, r := range text {
		if r >= '0' && r <= '9' {
			digits = append(digits, r)
		}
	}

	if len(digits) == 0 {
		return
	}

	// Format as phone number
	var formatted string
	switch {
	case len(digits) <= 3:
		formatted = "(" + string(digits)
	case len(digits) <= 6:
		formatted = "(" + string(digits[:3]) + ") " + string(digits[3:])
	case len(digits) <= 10:
		formatted = "(" + string(digits[:3]) + ") " + string(digits[3:6]) + "-" + string(digits[6:])
	default:
		// Truncate to 10 digits
		formatted = "(" + string(digits[:3]) + ") " + string(digits[3:6]) + "-" + string(digits[6:10])
	}

	if formatted != text {
		buffer.Replace(0, buffer.Length(), formatted)
		buffer.PlaceCursorAtEnd()
	}
}

// PhoneNumberTransformation creates an OutputTransformation that formats
// text as a US phone number.
func PhoneNumberTransformation() OutputTransformation {
	return &PhoneNumberOutputTransformation{}
}
