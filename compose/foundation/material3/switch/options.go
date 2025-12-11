package mswitch

type SwitchOptions struct {
	Modifier Modifier
}

type SwitchOption func(*SwitchOptions)

func DefaultSwitchOptions() SwitchOptions {
	return SwitchOptions{
		Modifier: EmptyModifier,
	}
}

func WithModifier(m Modifier) SwitchOption {
	return func(o *SwitchOptions) {
		o.Modifier = m
	}
}
