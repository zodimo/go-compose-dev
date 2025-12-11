package clickable

import "go-compose-dev/internal/modifier"

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

func OnClick(onClick func(), options ...ClickableOption) Modifier {

	opt := DefaultClickableOptions()
	for _, option := range options {
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
