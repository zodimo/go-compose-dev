package textfield

import (
	"git.sr.ht/~schnwalter/gio-mw/widget/input"
)

type TextFieldOptions struct {
	Modifier       Modifier
	Input          *input.Input
	SupportingText string
	Enabled        bool
	Error          bool
	SingleLine     bool
}

func DefaultTextFieldOptions() TextFieldOptions {
	return TextFieldOptions{
		Modifier:   EmptyModifier,
		Enabled:    true,
		Error:      false,
		SingleLine: true,
	}
}

type TextFieldOption func(*TextFieldOptions)

func WithModifier(m Modifier) TextFieldOption {
	return func(o *TextFieldOptions) {
		o.Modifier = m
	}
}

func WithSupportingText(text string) TextFieldOption {
	return func(o *TextFieldOptions) {
		o.SupportingText = text
	}
}

func WithEnabled(enabled bool) TextFieldOption {
	return func(o *TextFieldOptions) {
		o.Enabled = enabled
	}
}

func WithError(isError bool) TextFieldOption {
	return func(o *TextFieldOptions) {
		o.Error = isError
	}
}

func WithSingleLine(singleLine bool) TextFieldOption {
	return func(o *TextFieldOptions) {
		o.SingleLine = singleLine
	}
}
