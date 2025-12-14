package padding

import "github.com/zodimo/go-compose/internal/modifier"

type PaddingOptions struct {
	RtlAware bool
}

func DefaultPaddingOptions() PaddingOptions {
	return PaddingOptions{
		RtlAware: false,
	}
}

type PaddingOption func(options *PaddingOptions)

func Padding(start, top, end, bottom int, options ...PaddingOption) Modifier {

	opt := DefaultPaddingOptions()
	for _, option := range options {
		if option == nil {
			continue
		}
		option(&opt)
	}
	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&paddingElement{
				padding: PaddingData{
					Start:    start,
					Top:      top,
					End:      end,
					Bottom:   bottom,
					RtlAware: opt.RtlAware,
				},
			},
		),
		modifier.NewInspectorInfo(
			"padding",
			map[string]any{
				"Start":   start,
				"Top":     top,
				"End":     end,
				"Bottom":  bottom,
				"options": opt,
			},
		),
	)
}

func All(value int, options ...PaddingOption) Modifier {
	return Padding(value, value, value, value, options...)
}

func Horizontal(start, end int, options ...PaddingOption) Modifier {
	return Padding(start, NotSet, end, NotSet, options...)
}

func Vertical(top, bottom int, options ...PaddingOption) Modifier {
	return Padding(NotSet, top, NotSet, bottom, options...)
}
