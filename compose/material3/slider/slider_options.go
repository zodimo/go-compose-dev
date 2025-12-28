package slider

// SliderOption defines the functional option pattern for Slider.
type SliderOption func(*SliderOptions)

// SliderOptions holds all optional parameters for the Slider.
type SliderOptions struct {
	Modifier              Modifier
	Enabled               bool
	ValueRange            struct{ Min, Max float32 }
	Steps                 int
	OnValueChangeFinished func()
	Colors                SliderColors
}

// DefaultSliderOptions returns the default options.
func DefaultSliderOptions() SliderOptions {
	return SliderOptions{
		Modifier:              EmptyModifier,
		Enabled:               true,
		ValueRange:            struct{ Min, Max float32 }{Min: 0, Max: 1},
		Steps:                 0,
		OnValueChangeFinished: nil,
		Colors:                SliderDefaults.Colors(), // In practice, these should be resolved from theme if empty
	}
}

// WithModifier applies a modifier to the slider.
func WithModifier(m Modifier) SliderOption {
	return func(o *SliderOptions) {
		o.Modifier = m
	}
}

// WithEnabled sets whether the slider is enabled.
func WithEnabled(enabled bool) SliderOption {
	return func(o *SliderOptions) {
		o.Enabled = enabled
	}
}

// WithValueRange sets the range of valid values.
func WithValueRange(min, max float32) SliderOption {
	return func(o *SliderOptions) {
		o.ValueRange.Min = min
		o.ValueRange.Max = max
	}
}

// WithSteps sets the number of discrete steps.
func WithSteps(steps int) SliderOption {
	return func(o *SliderOptions) {
		o.Steps = steps
	}
}

// WithOnValueChangeFinished sets the callback for when dragging finishes.
func WithOnValueChangeFinished(callback func()) SliderOption {
	return func(o *SliderOptions) {
		o.OnValueChangeFinished = callback
	}
}

// WithColors sets the colors for the slider.
func WithColors(colors SliderColors) SliderOption {
	return func(o *SliderOptions) {
		o.Colors = colors
	}
}
