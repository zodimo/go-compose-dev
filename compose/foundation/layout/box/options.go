package box

import "github.com/zodimo/go-compose/compose/ui"

type BoxOptions struct {
	Modifier  ui.Modifier
	Alignment Direction
}

type BoxOption func(*BoxOptions)

func DefaultBoxOptions() BoxOptions {
	return BoxOptions{
		Modifier:  ui.EmptyModifier,
		Alignment: NW,
	}
}

func WithModifier(m ui.Modifier) BoxOption {
	return func(o *BoxOptions) {
		o.Modifier = m
	}
}

func WithAlignment(alignment Direction) BoxOption {
	return func(o *BoxOptions) {
		o.Alignment = alignment
	}
}
