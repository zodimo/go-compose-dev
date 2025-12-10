package row

type RowOptions struct {
	Modifier Modifier

	// Spacing controls the distribution of space left after
	// layout.
	Spacing Spacing
	// Alignment is the alignment in the cross axis.
	Alignment Alignment
}

type RowOption func(o *RowOptions)

func WithModifier(modifier Modifier) RowOption {
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
