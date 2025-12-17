package progress

import "git.sr.ht/~schnwalter/gio-mw/widget/indicator"

type IndicatorOptions struct {
	Modifier  Modifier
	Indicator *indicator.Indicator
}

type IndicatorOption func(o *IndicatorOptions)

func WithModifier(m Modifier) IndicatorOption {
	return func(o *IndicatorOptions) {
		o.Modifier = m
	}
}

func WithIndicator(ind *indicator.Indicator) IndicatorOption {
	return func(o *IndicatorOptions) {
		o.Indicator = ind
	}
}

func DefaultIndicatorOptions() IndicatorOptions {
	return IndicatorOptions{
		Modifier: EmptyModifier,
	}
}
