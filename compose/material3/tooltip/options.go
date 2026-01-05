package tooltip

import (
	"github.com/zodimo/go-compose/compose/ui"
)

// TooltipOptions contains configuration for tooltips.
type TooltipOptions struct {
	Modifier ui.Modifier
}

// TooltipOption is a functional option for TooltipOptions.
type TooltipOption func(*TooltipOptions)

// DefaultTooltipOptions returns options with an empty modifier.
func DefaultTooltipOptions() TooltipOptions {
	return TooltipOptions{
		Modifier: ui.EmptyModifier,
	}
}

// WithModifier appends a modifier to the tooltip.
func WithModifier(m ui.Modifier) TooltipOption {
	return func(o *TooltipOptions) {
		o.Modifier = m
	}
}
