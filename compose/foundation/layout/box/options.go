package box

type BoxOptions struct {
	Modifier  Modifier
	Alignment Direction
}

type BoxOption func(*BoxOptions)

func DefaultBoxOptions() BoxOptions {
	return BoxOptions{
		Modifier:  EmptyModifier,
		Alignment: NW,
	}
}

func WithModifier(m Modifier) BoxOption {
	return func(o *BoxOptions) {
		o.Modifier = m
	}
}

func WithAlignment(alignment Direction) BoxOption {
	return func(o *BoxOptions) {
		o.Alignment = alignment
	}
}
