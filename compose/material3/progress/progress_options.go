package progress

import (
	"git.sr.ht/~schnwalter/gio-mw/widget/indicator"
	"github.com/zodimo/go-compose/compose/ui"
)

type IndicatorOptions struct {
	Modifier  ui.Modifier
	Indicator *indicator.Indicator
}

type IndicatorOption func(o *IndicatorOptions)

func WithModifier(m ui.Modifier) IndicatorOption {
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
		Modifier: ui.EmptyModifier,
	}
}
