# Color Transformation Guide

This document explains how to add new color transformations to the theme system.

## Overview

Color transformations are implemented using the `ColorUpdate` system, which allows color descriptors to carry a chain of transformations that are applied when the color is resolved at runtime.

## Existing Transformations

- **SetOpacity**: Adjusts the alpha channel of a color
- **Lighten**: Increases color brightness by a percentage
- **Darken**: Decreases color brightness by a percentage

## Adding a New Transformation

### Step 1: Add the Action Constant

In `color.go`, add a new constant to `ColorUpdateActions`:

```go
const (
    SetOpacityColorUpdateAction ColorUpdateActions = iota
    LightenColorUpdateAction
    DarkenColorUpdateAction
    YourNewAction  // Add here
)
```

### Step 2: Add Interface Method

Add a method to the `ColorDescriptor` interface:

```go
type ColorDescriptor interface {
    // ... existing methods
    YourNewMethod(param ParamType) ColorDescriptor
}
```

### Step 3: Implement the Method

Add the implementation to `colorDescriptor`:

```go
func (t colorDescriptor) YourNewMethod(param ParamType) ColorDescriptor {
    update := CreateYourUpdate(param)
    return t.AppendUpdate(update.Any())
}
```

### Step 4: Create Helper Functions

Add constructor and getter functions:

```go
func CreateYourUpdate(param ParamType) *ColorUpdateTyped[ParamType] {
    return &ColorUpdateTyped[ParamType]{
        action: YourNewAction,
        value:  param,
    }
}

func GetYourParam(update ColorUpdate) ParamType {
    if update.Action() != YourNewAction {
        panic("update is not your action type")
    }
    return update.Value().(ParamType)
}
```

### Step 5: Implement Resolution Logic

In `color_resolver.go` (or the theme manager), add logic to handle your transformation:

```go
func (tm *themeManager) ResolveColorDescriptor(desc ColorDescriptor) ThemeColor {
    // ... existing code to get base color
    
    for _, update := range desc.Updates() {
        switch update.Action() {
        case YourNewAction:
            param := GetYourParam(update)
            color = applyYourTransformation(color, param)
        // ... other cases
        }
    }
    
    return color
}
```

### Step 6: Add Tests

Create tests in `color_transformations_test.go`:

```go
func TestYourTransformation(t *testing.T) {
    colorHelper := ColorHelper
    baseColor := colorHelper.ColorSelector().Primary
    
    transformed := baseColor.YourNewMethod(testParam)
    
    // Verify the update was added
    updates := transformed.Updates()
    require.Len(t, updates, 1)
    require.Equal(t, YourNewAction, updates[0].Action())
    
    // Test resolution
    manager := GetThemeManager()
    resolved := manager.ResolveColorDescriptor(transformed)
    
    // Assert expected color values
}
```

## Best Practices

### Color Space Conversions

When implementing transformations that adjust color properties:

1. **For brightness/lightness**: Use HSL color space
2. **For saturation**: Use HSL or HSV color space
3. **For opacity**: Work directly with RGBA
4. **For hue rotation**: Use HSL or HSV color space

Example HSL conversion:

```go
func rgbaToHSL(c color.NRGBA) (h, s, l float32) {
    // Normalize RGB to 0-1
    r := float32(c.R) / 255.0
    g := float32(c.G) / 255.0
    b := float32(c.B) / 255.0
    
    // Calculate HSL
    // ... implementation
    return h, s, l
}

func hslToRGBA(h, s, l float32, a uint8) color.NRGBA {
    // Convert HSL back to RGB
    // ... implementation
    return color.NRGBA{R: r, G: g, B: b, A: a}
}
```

### Parameter Ranges

Document expected parameter ranges clearly:

- **Percentages**: Use `float32` in range 0.0 to 1.0 (e.g., 0.2 = 20%)
- **Degrees**: Use `float32` in range 0.0 to 360.0
- **Opacity**: Use `OpacityLevel` type from token package

### Transformation Order

Transformations are applied in the order they're added:

```go
// Lighten first, then reduce opacity
color := baseColor.Lighten(0.2).SetOpacity(Opacity50)

// Different result:
color := baseColor.SetOpacity(Opacity50).Lighten(0.2)
```

Document any order-dependent behavior.

## Example: Lighten Implementation

Here's how `Lighten` is implemented as a reference:

```go
// 1. Constant
const (
    LightenColorUpdateAction ColorUpdateActions = iota + 1
)

// 2. Interface method
func (t colorDescriptor) Lighten(percentage float32) ColorDescriptor {
    update := Lighten(percentage)
    return t.AppendUpdate(update.Any())
}

// 3. Constructor
func Lighten(percentage float32) *ColorUpdateTyped[float32] {
    return &ColorUpdateTyped[float32]{
        action: LightenColorUpdateAction,
        value:  percentage,
    }
}

// 4. Getter
func GetLighten(update ColorUpdate) float32 {
    if update.Action() != LightenColorUpdateAction {
        panic("update is not a lighten update")
    }
    return update.Value().(float32)
}

// 5. Resolution (in theme manager)
case LightenColorUpdateAction:
    percentage := GetLighten(update)
    h, s, l := rgbaToHSL(currentColor)
    l = clamp(l + percentage, 0.0, 1.0)
    currentColor = hslToRGBA(h, s, l, currentColor.A)
```

## Testing Guidelines

Every transformation should have:

1. **Unit tests** for the transformation creation
2. **Integration tests** for color resolution
3. **Visual tests** (demos) showing the transformation in action
4. **Edge case tests** (e.g., clamping, zero values, extreme values)

## Performance Considerations

- Transformations are applied at every layout frame
- Keep transformation logic simple and fast
- Consider caching resolved colors if performance becomes an issue
- Profile before optimizing

## Documentation

When adding a transformation, also document:

- What the transformation does
- Parameter ranges and units
- Example usage
- Visual examples in demos
- Any Material 3 guidelines it supports
