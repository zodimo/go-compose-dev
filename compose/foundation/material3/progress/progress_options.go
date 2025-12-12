package progress

import "git.sr.ht/~schnwalter/gio-mw/widget/indicator"

type IndicatorOptions struct {
	Modifier  Modifier
	Indicator *indicator.Indicator
}

type IndicatorOption func(o *IndicatorOptions)

func WithModifier(modifier Modifier) IndicatorOption {
	return func(o *IndicatorOptions) {
		o.Modifier = o.Modifier.Then(modifier)
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
