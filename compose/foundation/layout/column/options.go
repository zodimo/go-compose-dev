package column

import "go-compose-dev/internal/modifier"

type ColumnOptions struct {
	Modifier modifier.Modifier

	// Spacing controls the distribution of space left after
	// layout.
	Spacing Spacing
	// Alignment is the alignment in the cross axis.
	Alignment Alignment
}

type ColumnOption func(o *ColumnOptions)

func WithModifier(modifier modifier.Modifier) ColumnOption {
	return func(o *ColumnOptions) {
		o.Modifier = o.Modifier.Then(modifier)
	}
}

func WithSpacing(spacing Spacing) ColumnOption {
	return func(o *ColumnOptions) {
		o.Spacing = spacing
	}
}

func WithAlignment(alignment Alignment) ColumnOption {
	return func(o *ColumnOptions) {
		o.Alignment = alignment
	}
}
