# GoCompose Styling Patterns Guide

This guide documents idiomatic patterns for styling components in GoCompose, following Jetpack Compose principles.

## Core Principle: Prefer Modifiers Over Wrappers

When styling a single element, **chain modifiers directly** rather than wrapping in a container component like `Surface`.

### ✅ Idiomatic Pattern: Modifier Chain

```go
// Circle-clipped image with border
fImage.Image(
    imageResource,
    fImage.WithContentScale(uilayout.ContentScaleCrop),
    fImage.WithAlignment(size.Center),
    fImage.WithModifier(
        size.Size(100, 100).
            Then(clip.Clip(shape.ShapeCircle)).
            Then(border.Border(4, colors.PrimaryRoles.Primary, shape.ShapeCircle)),
    ),
)
```

### ❌ Avoid: Unnecessary Surface Wrapper

```go
// Don't do this for simple styling
surface.Surface(
    fImage.Image(imageResource, ...),
    surface.WithShape(shape.ShapeCircle),
    surface.WithBorder(4, borderColor),
    surface.WithModifier(size.Size(108, 108)),
)
```

## When to Use Each Pattern

| Use Case | Pattern |
|----------|---------|
| Single element with clip/border | **Modifier chain** |
| Element with background color | **Modifier chain** with `background.Background()` |
| Container with multiple children + shared styling | **Surface** |
| Card with elevation, padding, and content | **Surface** |
| Semantic elevation (FAB, Dialog, etc.) | **Surface** |

## Card vs Surface

Use `card.Elevated/Filled/Outlined` when:
- Building a **simple content card** with title, body, optional image
- Following the standard **M3 Card layout pattern**
- Content flows **linearly** (no overlapping elements)

```go
card.Elevated(
    card.CardContents(
        card.Image(gioImage),
        card.Content(titleAndBody),
    ),
)
```

Use `surface.Surface` when:
- Building **custom layouts** with overlapping elements
- Need **partial rounded corners** (e.g., only top corners rounded)
- Implementing **non-standard card designs** (profile cards with overlapping avatars)
- Creating **dialog/sheet backgrounds**

```go
// Custom profile card with overlapping header
surface.Surface(
    column.Column(
        c.Sequence(
            // Header with only top corners rounded
            surface.Surface(headerContent,
                surface.WithShape(shape.RoundedCornerShape{TopStart: 16, TopEnd: 16}),
            ),
            // Overlapping profile image
            profileImageComposable,
            // Content
            contentComposable,
        ),
    ),
    surface.WithShape(shape.RoundedCornerShape{Radius: 16}),
    surface.WithShadowElevation(4),
)
```

## Available Styling Modifiers

| Modifier | Purpose | Example |
|----------|---------|---------|
| `clip.Clip(shape)` | Clip to shape | `clip.Clip(shape.ShapeCircle)` |
| `border.Border(width, color, shape)` | Draw border | `border.Border(2, color, shape.ShapeCircle)` |
| `background.Background(color)` | Fill background | `background.Background(colors.SurfaceRoles.Surface)` |
| `shadow.Simple(elevation, shape)` | Add shadow | `shadow.Simple(4, shape.RoundedCornerShape{Radius: 8})` |
| `padding.All(dp)` | Add padding | `padding.All(16)` |
| `offset.Offset(x, y)` | Translate position | `offset.OffsetY(-50)` — overlapping layouts |

## Color Pattern: Use ColorSelector

Always use the `ColorSelector()` pattern for theme-reactive colors:

```go
colors := theme.ColorHelper.ColorSelector()

// Theme role references (reactive to theme changes)
colors.PrimaryRoles.Primary
colors.SurfaceRoles.OnSurface
colors.SurfaceRoles.OnVariant

// For explicit colors, use SpecificColor
theme.ColorHelper.SpecificColor(color.NRGBA{R: 255, G: 0, B: 0, A: 255})
```

## Conditional Values (go-ternary)

Go lacks ternary operators, so use the `go-ternary` package for conditional values:

```go
import "github.com/zodimo/go-ternary"

colors := theme.ColorHelper.ColorSelector()
specificColor := theme.ColorHelper.SpecificColor

// Conditional color based on state
surface.WithColor(ternary.Ternary(
    selected,
    colors.SecondaryRoles.Container,
    specificColor(color.NRGBA{A: 0}), // Transparent
)),
```

This mirrors Kotlin's inline if-expression:
```kotlin
color = if (selected) colorScheme.secondaryContainer else Color.Transparent
```

## Common Shape Options

```go
shape.ShapeRectangle                    // No rounding
shape.ShapeCircle                       // Full circle/ellipse
shape.RoundedCornerShape{Radius: 16}    // Uniform corners
shape.RoundedCornerShape{              // Per-corner control
    TopStart: 16,
    TopEnd:   16,
    BottomStart: 0,
    BottomEnd: 0,
}
```
