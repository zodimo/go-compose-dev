package clickable

import (
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/internal/modifier"
)

type ClickableOptions struct {
	Clickable *GioClickable
}

type ClickableOption func(*ClickableOptions)

func WithClickable(clickable *GioClickable) ClickableOption {
	return func(options *ClickableOptions) {
		options.Clickable = clickable
	}
}

func DefaultClickableOptions() ClickableOptions {
	return ClickableOptions{
		Clickable: nil,
	}
}

func OnClick(onClick func(), options ...ClickableOption) ui.Modifier {

	opt := DefaultClickableOptions()
	for _, option := range options {
		if option == nil {
			continue
		}
		option(&opt)
	}
	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&ClickableElement{
				clickableData: ClickableData{
					OnClick:   onClick,
					Clickable: opt.Clickable,
				},
			},
		),
		modifier.NewInspectorInfo(
			"clickable",
			map[string]any{
				"onClick": onClick,
				"options": opt,
			},
		),
	)
}
