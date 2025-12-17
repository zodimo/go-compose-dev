package card

type CardOptions struct {
	Modifier Modifier
}

type CardOption func(o *CardOptions)

func WithModifier(m Modifier) CardOption {
	return func(o *CardOptions) {
		o.Modifier = m
	}
}
