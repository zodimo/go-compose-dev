package row

import "github.com/zodimo/go-compose/compose/ui"

type RowOptions struct {
	Modifier ui.Modifier

	// Spacing controls the distribution of space left after
	// layout.
	Spacing Spacing
	// Alignment is the alignment in the cross axis.
	Alignment Alignment
}

type RowOption func(o *RowOptions)

func WithModifier(modifier ui.Modifier) RowOption {
	return func(o *RowOptions) {
		o.Modifier = o.Modifier.Then(modifier)
	}
}

func WithSpacing(spacing Spacing) RowOption {
	return func(o *RowOptions) {
		o.Spacing = spacing
	}
}

func WithAlignment(alignment Alignment) RowOption {
	return func(o *RowOptions) {
		o.Alignment = alignment
	}
}
