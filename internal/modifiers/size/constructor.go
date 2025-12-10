package size

import (
	"go-compose-dev/internal/modifier"
)

type SizeOptions struct {
	Required bool
}

func DefaultSizeOptions() SizeOptions {
	return SizeOptions{
		Required: false,
	}
}

type SizeOption func(options *SizeOptions)

func Size(width, height int, options ...SizeOption) Modifier {

	opt := DefaultSizeOptions()
	for _, option := range options {
		option(&opt)
	}
	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&SizeElement{
				size: SizeData{
					Width:    width,
					Height:   height,
					Required: opt.Required,
				},
			},
		),
		modifier.NewInspectorInfo(
			"size",
			map[string]any{
				"width":   width,
				"height":  height,
				"options": opt,
			},
		),
	)
}
