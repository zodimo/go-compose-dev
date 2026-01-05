package textfield

import (
	"github.com/zodimo/go-compose/compose/foundation/layout"
	"github.com/zodimo/go-compose/compose/foundation/text/input"
	"github.com/zodimo/go-compose/compose/material3"
	"github.com/zodimo/go-compose/compose/material3/tokens"
	"github.com/zodimo/go-compose/compose/ui/graphics/shape"
	"github.com/zodimo/go-compose/compose/ui/next/unit"
)

type TextFieldLabelPosition int

const (
	TextFieldLabelPositionAttached TextFieldLabelPosition = iota
	TextFieldLabelPositionAbove
)

var _ input.TextFieldDecorator = (*OutlinedTextFieldDecorator)(nil)

type OutlinedTextFieldDecorator struct {
	State                input.TextFieldState
	Enabled              bool
	LineLimits           input.TextFieldLineLimits
	OutputTransformation input.OutputTransformation
	LabelPosition        TextFieldLabelPosition
	Label                Composable
	Placeholder          Composable
	LeadingIcon          Composable
	TrailingIcon         Composable
	Prefix               Composable
	Suffix               Composable
	SupportingText       Composable
	IsError              bool
	Colors               TextFieldColors
	ContentPadding       layout.PaddingValues
	Container            Composable
}

func (d OutlinedTextFieldDecorator) Decoration(innerTextField Composable) Composable {
	return innerTextField
}

func OutlinedTextFieldDefaults(c Composer) *OutlinedTextFieldDefaultOptions {
	shapes := material3.LocalShapes.Current(c)
	return &OutlinedTextFieldDefaultOptions{
		Shape:                    shapes.FromToken(tokens.FilledTextFieldTokens.ContainerShape),
		MinHeight:                unit.Dp(48),
		MinWidth:                 unit.Dp(0),
		UnfocusedBorderThickness: unit.Dp(1),
		FocusedBorderThickness:   unit.Dp(1),
		Decorator:                OutlinedTextFieldDecorator{},
	}
}

type OutlinedTextFieldDefaultOptions struct {
	Shape                    shape.Shape
	MinHeight                unit.Dp
	MinWidth                 unit.Dp
	UnfocusedBorderThickness unit.Dp
	FocusedBorderThickness   unit.Dp
	Decorator                input.TextFieldDecorator
}
