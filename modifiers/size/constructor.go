package size

import (
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/internal/modifier"
)

type SizeOptions struct {
	Required bool
}

func DefaultSizeOptions() SizeOptions {
	return SizeOptions{
		Required: false,
	}
}

func SizeRequired() SizeOption {
	return func(options *SizeOptions) {
		options.Required = true
	}
}

type SizeOption func(options *SizeOptions)

func Size(width, height int, options ...SizeOption) ui.Modifier {

	opt := DefaultSizeOptions()
	for _, option := range options {
		if option == nil {
			continue
		}
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

func FillMax() ui.Modifier {

	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&SizeElement{
				size: SizeData{
					Width:   NotSet,
					Height:  NotSet,
					FillMax: true,
				},
			},
		),
		modifier.NewInspectorInfo(
			"size",
			map[string]any{
				"FillMax": true,
			},
		),
	)
}

func FillMaxWidth() ui.Modifier {

	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&SizeElement{
				size: SizeData{
					Width:        NotSet,
					Height:       NotSet,
					FillMaxWidth: true,
				},
			},
		),
		modifier.NewInspectorInfo(
			"size",
			map[string]any{
				"FillMaxWidth": true,
			},
		),
	)
}

func FillMaxHeight() ui.Modifier {

	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&SizeElement{
				size: SizeData{
					Width:         NotSet,
					Height:        NotSet,
					FillMaxHeight: true,
				},
			},
		),
		modifier.NewInspectorInfo(
			"size",
			map[string]any{
				"FillMaxHeight": true,
			},
		),
	)
}

func WrapContentSize(align ...Alignment) ui.Modifier {
	var a Alignment = Center
	if len(align) > 0 {
		a = align[0]
	}

	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&SizeElement{
				size: SizeData{
					Width:      NotSet,
					Height:     NotSet,
					WrapWidth:  true,
					WrapHeight: true,
					Alignment:  a,
					Unbounded:  false,
				},
			},
		),
		modifier.NewInspectorInfo(
			"wrapContentSize",
			map[string]any{
				"align":     a,
				"unbounded": false,
			},
		),
	)
}

func WrapContentWidth(align ...Alignment) ui.Modifier {
	var a Alignment = Center
	if len(align) > 0 {
		a = align[0]
	}

	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&SizeElement{
				size: SizeData{
					Width:      NotSet,
					Height:     NotSet,
					WrapWidth:  true,
					WrapHeight: false,
					Alignment:  a,
					Unbounded:  false,
				},
			},
		),
		modifier.NewInspectorInfo(
			"wrapContentWidth",
			map[string]any{
				"align":     a,
				"unbounded": false,
			},
		),
	)
}

func WrapContentHeight(align ...Alignment) ui.Modifier {
	var a Alignment = Center
	if len(align) > 0 {
		a = align[0]
	}

	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&SizeElement{
				size: SizeData{
					Width:      NotSet,
					Height:     NotSet,
					WrapWidth:  false,
					WrapHeight: true,
					Alignment:  a,
					Unbounded:  false,
				},
			},
		),
		modifier.NewInspectorInfo(
			"wrapContentHeight",
			map[string]any{
				"align":     a,
				"unbounded": false,
			},
		),
	)
}

func Width(width int, options ...SizeOption) ui.Modifier {

	opt := DefaultSizeOptions()
	for _, option := range options {
		if option == nil {
			continue
		}
		option(&opt)
	}

	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&SizeElement{
				size: SizeData{
					Width:    width,
					Height:   NotSet,
					Required: opt.Required,
				},
			},
		),
		modifier.NewInspectorInfo(
			"size",
			map[string]any{
				"width":   width,
				"options": opt,
			},
		),
	)
}

func Height(height int, options ...SizeOption) ui.Modifier {

	opt := DefaultSizeOptions()
	for _, option := range options {
		if option == nil {
			continue
		}
		option(&opt)
	}

	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&SizeElement{
				size: SizeData{
					Width:    NotSet,
					Height:   height,
					Required: opt.Required,
				},
			},
		),
		modifier.NewInspectorInfo(
			"size",
			map[string]any{
				"height":  height,
				"options": opt,
			},
		),
	)
}

// MinWidth sets the minimum width constraint for the element.
func MinWidth(minWidth int) ui.Modifier {
	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&SizeElement{
				size: SizeData{
					Width:    NotSet,
					Height:   NotSet,
					MinWidth: minWidth,
				},
			},
		),
		modifier.NewInspectorInfo(
			"minWidth",
			map[string]any{
				"minWidth": minWidth,
			},
		),
	)
}

// MaxWidth sets the maximum width constraint for the element.
func MaxWidth(maxWidth int) ui.Modifier {
	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&SizeElement{
				size: SizeData{
					Width:    NotSet,
					Height:   NotSet,
					MaxWidth: maxWidth,
				},
			},
		),
		modifier.NewInspectorInfo(
			"maxWidth",
			map[string]any{
				"maxWidth": maxWidth,
			},
		),
	)
}

// MinHeight sets the minimum height constraint for the element.
func MinHeight(minHeight int) ui.Modifier {
	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&SizeElement{
				size: SizeData{
					Width:     NotSet,
					Height:    NotSet,
					MinHeight: minHeight,
				},
			},
		),
		modifier.NewInspectorInfo(
			"minHeight",
			map[string]any{
				"minHeight": minHeight,
			},
		),
	)
}

// MaxHeight sets the maximum height constraint for the element.
func MaxHeight(maxHeight int) ui.Modifier {
	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&SizeElement{
				size: SizeData{
					Width:     NotSet,
					Height:    NotSet,
					MaxHeight: maxHeight,
				},
			},
		),
		modifier.NewInspectorInfo(
			"maxHeight",
			map[string]any{
				"maxHeight": maxHeight,
			},
		),
	)
}
