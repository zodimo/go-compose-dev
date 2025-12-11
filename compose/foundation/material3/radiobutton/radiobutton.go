package radiobutton

import (
	"fmt"
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"

	"go-compose-dev/internal/layoutnode"

	"git.sr.ht/~schnwalter/gio-mw/wdk"
)

const Material3RadioButtonNodeID = "Material3RadioButton"

// RadioButton creates a Material3 radio button.
// RadioButton creates a Material3 radio button.
func RadioButton(
	selected bool,
	onClick func(),
	options ...RadioButtonOption,
) Composable {
	return func(c Composer) Composer {
		// Resolve options
		opts := DefaultRadioButtonOptions()
		for _, option := range options {
			option(&opts)
		}

		key := c.GenerateID()
		path := c.GetPath()

		// Fix closure capture for handler
		handlerWrapperState := c.State(fmt.Sprintf("%d/%s/handler_wrapper", key, path), func() any {
			return &HandlerWrapper{Func: onClick}
		})
		handlerWrapper := handlerWrapperState.Get().(*HandlerWrapper)
		handlerWrapper.Func = onClick

		// State for Clickable
		clickableState := c.State(fmt.Sprintf("%d/%s/clickable", key, path), func() any {
			return &widget.Clickable{}
		})
		clickable := clickableState.Get().(*widget.Clickable)

		c.StartBlock(Material3RadioButtonNodeID)
		c.Modifier(func(m Modifier) Modifier {
			// Apply user modifier
			return m.Then(opts.Modifier)
		})

		c.SetWidgetConstructor(radioButtonWidgetConstructor(
			clickable,
			selected,
			opts.Enabled,
			opts.Colors,
			handlerWrapper,
		))

		return c.EndBlock()
	}
}

func radioButtonWidgetConstructor(
	clickable *widget.Clickable,
	selected bool,
	enabled bool,
	colors RadioButtonColors,
	handler *HandlerWrapper,
) layoutnode.LayoutNodeWidgetConstructor {
	return layoutnode.NewLayoutNodeWidgetConstructor(func(node layoutnode.LayoutNode) layoutnode.GioLayoutWidget {
		return func(gtx layoutnode.LayoutContext) layoutnode.LayoutDimensions {
			if !enabled {
				// We don't enable the clickable if disabled
			} else if clickable.Clicked(gtx) {
				if handler.Func != nil {
					handler.Func()
				}
			}

			// Resolve colors from theme if not provided
			theme := wdk.GetMaterialTheme(gtx)

			// Unselected Color
			unselectedColorVal := colors.UnselectedColor
			if unselectedColorVal == nil {
				unselectedColorVal = theme.Scheme.SurfaceVariant.OnColor.AsNRGBA()
			}
			// Selected Color
			selectedColorVal := colors.SelectedColor
			if selectedColorVal == nil {
				selectedColorVal = theme.Scheme.Primary.Color.AsNRGBA()
			}
			// Disabled Color
			disabledColorVal := colors.DisabledColor
			if disabledColorVal == nil {
				// Material 3 Disabled state: 38% opacity of OnSurface
				// We manually apply 38% opacity (approx 97/255)
				c := theme.Scheme.Surface.OnColor.AsNRGBA()
				c.A = 97
				disabledColorVal = c
			}

			// Helper to convert to NRGBA
			asNRGBA := func(c color.Color) color.NRGBA {
				if c == nil {
					return color.NRGBA{}
				}
				if n, ok := c.(color.NRGBA); ok {
					return n
				}
				r, g, b, a := c.RGBA()
				return color.NRGBA{
					R: uint8(r >> 8),
					G: uint8(g >> 8),
					B: uint8(b >> 8),
					A: uint8(a >> 8),
				}
			}

			// Layout constants based on Material 3
			const (
				iconSize       = 20
				stateLayerSize = 40
				strokeWidth    = 2
			)

			sizeDp := unit.Dp(stateLayerSize)
			sizePx := gtx.Dp(sizeDp)

			// Center the icon within the state layer
			iconSizePx := gtx.Dp(unit.Dp(iconSize))
			iconOffset := (sizePx - iconSizePx) / 2

			return layout.Stack{}.Layout(gtx,
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					// Draw state layer (hover/press)
					if enabled {
						if clickable.Hovered() || clickable.Pressed() {
							// Draw circle
							c := unselectedColorVal
							if selected {
								c = selectedColorVal
							}

							targetColor := asNRGBA(c)
							// Apply opacity: ~10% for state layer
							targetColor.A = 25

							defer op.Offset(image.Pt(0, 0)).Push(gtx.Ops).Pop()

							circle := clip.Ellipse{
								Max: image.Pt(sizePx, sizePx),
							}.Path(gtx.Ops)
							paint.FillShape(gtx.Ops, targetColor, clip.Outline{Path: circle}.Op())
						}

					}
					return layout.Dimensions{Size: image.Pt(sizePx, sizePx)}
				}),
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					// Handle clicks
					if enabled {
						return clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Dimensions{Size: image.Pt(sizePx, sizePx)}
						})
					}
					return layout.Dimensions{Size: image.Pt(sizePx, sizePx)}
				}),
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					// Draw Radio Icon
					defer op.Offset(image.Pt(iconOffset, iconOffset)).Push(gtx.Ops).Pop()

					strokeWidthPx := gtx.Dp(unit.Dp(strokeWidth))

					// Outer constant circle
					outerCircle := clip.Ellipse{
						Max: image.Pt(iconSizePx, iconSizePx),
					}.Path(gtx.Ops)

					borderColor := unselectedColorVal
					if !enabled {
						borderColor = disabledColorVal
					} else if selected {
						borderColor = selectedColorVal
					}

					paint.FillShape(gtx.Ops, asNRGBA(borderColor), clip.Stroke{
						Path:  outerCircle,
						Width: float32(strokeWidthPx),
					}.Op())

					// Inner circle if selected
					if selected {
						innerCircleSize := iconSizePx - strokeWidthPx*4
						if innerCircleSize < 0 {
							innerCircleSize = 0
						}

						innerOffset := strokeWidthPx * 2

						defer op.Offset(image.Pt(innerOffset, innerOffset)).Push(gtx.Ops).Pop()

						innerCircle := clip.Ellipse{
							Max: image.Pt(innerCircleSize, innerCircleSize),
						}.Path(gtx.Ops)

						fillColor := selectedColorVal
						if !enabled {
							fillColor = disabledColorVal
						}

						paint.FillShape(gtx.Ops, asNRGBA(fillColor), clip.Outline{Path: innerCircle}.Op())
					}

					return layout.Dimensions{Size: image.Pt(sizePx, sizePx)}
				}),
			)
		}
	})
}
