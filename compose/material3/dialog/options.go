package dialog

import (
	"github.com/zodimo/go-compose/compose/material3/text"
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/pkg/api"

	"git.sr.ht/~schnwalter/gio-mw/token"
)

// DialogOptions configures an AlertDialog.
type DialogOptions struct {
	// Modifier applied to the dialog container.
	Modifier ui.Modifier
	// Icon composable displayed above the title.
	Icon api.Composable
	// Title composable displayed above the text.
	Title api.Composable
	// Text composable displayed as the dialog body.
	Text api.Composable
	// DismissButton composable for the dismiss/cancel action.
	DismissButton api.Composable
}

// DialogOption is a function that modifies DialogOptions.
type DialogOption func(*DialogOptions)

// DefaultDialogOptions returns the default dialog options.
func DefaultDialogOptions() DialogOptions {
	return DialogOptions{
		Modifier: ui.EmptyModifier,
	}
}

// WithModifier sets the modifier for the dialog.
func WithModifier(m ui.Modifier) DialogOption {
	return func(o *DialogOptions) {
		o.Modifier = m
	}
}

// WithIcon sets a composable icon displayed above the title.
func WithIcon(icon api.Composable) DialogOption {
	return func(o *DialogOptions) {
		o.Icon = icon
	}
}

// WithTitle sets a composable title displayed above the text.
func WithTitle(title api.Composable) DialogOption {
	return func(o *DialogOptions) {
		o.Title = title
	}
}

// WithTitleText is a convenience function that sets the title as a text string
// with HeadlineSmall typography.
func WithTitleText(titleStr string) DialogOption {
	return func(o *DialogOptions) {
		o.Title = text.TextWithStyle(titleStr, token.TypestyleHeadlineSmall)
	}
}

// WithText sets a composable text displayed as the dialog body.
func WithText(textComposable api.Composable) DialogOption {
	return func(o *DialogOptions) {
		o.Text = textComposable
	}
}

// WithTextContent is a convenience function that sets the text content as a string
// with BodyMedium typography.
func WithTextContent(content string) DialogOption {
	return func(o *DialogOptions) {
		o.Text = text.TextWithStyle(content, token.TypestyleBodyMedium)
	}
}

// WithDismissButton sets a composable dismiss/cancel button.
func WithDismissButton(button api.Composable) DialogOption {
	return func(o *DialogOptions) {
		o.DismissButton = button
	}
}
