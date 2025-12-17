package card

import (
	"github.com/zodimo/go-compose/compose/ui/graphics/shape"
	"github.com/zodimo/go-compose/internal/layoutnode"
	"github.com/zodimo/go-compose/modifiers/background"
	"github.com/zodimo/go-compose/modifiers/clip"
	"github.com/zodimo/go-compose/modifiers/shadow"
	"github.com/zodimo/go-compose/theme"

	"gioui.org/layout"
	"gioui.org/unit"
	"git.sr.ht/~schnwalter/gio-mw/wdk/block"
)

const Material3CardNodeID = "Material3Card"
const Material3CardImageSlotID = "Material3CardImage"

// Card corner radius (Material 3 Medium = 12dp)
var cardCornerShape = shape.RoundedCornerShape{Radius: unit.Dp(12)}

func Elevated(contents CardContentContainer, options ...CardOption) Composable {
	return cardComposable(cardElevated, contents, options...)
}

func Filled(contents CardContentContainer, options ...CardOption) Composable {
	return cardComposable(cardFilled, contents, options...)
}

func Outlined(contents CardContentContainer, options ...CardOption) Composable {
	return cardComposable(cardOutlined, contents, options...)
}

func cardComposable(kind cardKind, contents CardContentContainer, options ...CardOption) Composable {
	return func(c Composer) Composer {
		opts := DefaultCardOptions()
		for _, option := range options {
			if option == nil {
				continue
			}
			option(&opts)
		}

		args := CardConstructorArgs{
			Kind:     kind,
			Contents: contents,
		}

		c.StartBlock(Material3CardNodeID)

		// Build modifier chain: user modifier + card styling modifiers
		c.Modifier(func(modifier Modifier) Modifier {
			// Start with user-provided modifier
			m := modifier.Then(opts.Modifier)

			// Get card colors from theme
			colorRoles := theme.ColorHelper.ColorSelector()

			// Add background and clip based on card kind
			switch kind {
			case cardElevated:
				// Elevated card: shadow + background + rounded clip
				m = m.Then(shadow.Simple(unit.Dp(2), cardCornerShape))
				m = m.Then(background.Background(colorRoles.SurfaceRoles.ContainerLow, background.WithShape(cardCornerShape)))
				m = m.Then(clip.Clip(cardCornerShape))
			case cardOutlined:
				// Outlined card: background + rounded clip (TODO: add border)
				m = m.Then(background.Background(colorRoles.SurfaceRoles.Surface, background.WithShape(cardCornerShape)))
				m = m.Then(clip.Clip(cardCornerShape))
			default: // Filled
				// Filled card: background + rounded clip
				m = m.Then(background.Background(colorRoles.SurfaceRoles.ContainerHighest, background.WithShape(cardCornerShape)))
				m = m.Then(clip.Clip(cardCornerShape))
			}

			return m
		})

		images := []*GioImage{}
		for _, child := range contents.children {
			if child.contentType == CardContentImage {
				images = append(images, child.image)
				continue
			}
			c.WithComposable(child.composable)
		}

		c.EmitSlot(Material3CardImageSlotID, images)

		c.SetWidgetConstructor(cardWidgetConstructor(args))

		return c.EndBlock()
	}
}

type CardConstructorArgs struct {
	Kind     cardKind
	Contents CardContentContainer
}

func getCardImages(node layoutnode.LayoutNode) []*GioImage {
	return node.FindSlot(Material3CardImageSlotID).UnwrapUnsafe().([]*GioImage)
}

func cardWidgetConstructor(args CardConstructorArgs) layoutnode.LayoutNodeWidgetConstructor {
	return layoutnode.NewLayoutNodeWidgetConstructor(func(node layoutnode.LayoutNode) layoutnode.GioLayoutWidget {
		return func(gtx layoutnode.LayoutContext) layoutnode.LayoutDimensions {

			nodeChildren := node.Children()
			images := getCardImages(node)

			// Build layout widgets for children
			var widgets []block.Segment
			for _, indexedChild := range args.Contents.children {
				switch indexedChild.contentType {
				case CardContentCover:
					nodeIndex := indexedChild.childIndex
					widgets = append(widgets, block.NewSegment(nodeChildren[nodeIndex].(layoutnode.NodeCoordinator).Layout))
				case CardContentImage:
					imageIndex := indexedChild.childIndex
					img := images[imageIndex]
					widgets = append(widgets, block.NewSegment(func(gtx layout.Context) layout.Dimensions {
						return img.Layout(gtx)
					}))
				default:
					// content
					nodeIndex := indexedChild.childIndex
					widgets = append(widgets, block.NewSegment(nodeChildren[nodeIndex].(layoutnode.NodeCoordinator).Layout))
				}
			}

			// Simple layout: just lay out children vertically
			// NO Clickable wrapper - this fixes focus navigation!
			return block.Line{
				Axis:     block.AxisVertical,
				Overflow: block.OverflowClip,
			}.Layout(gtx, widgets...)
		}
	})
}
