package modifier

import (
	"fmt"
	"image/color"
	"strings"
)

// DebugModifier returns a formatted string representation of the Modifier chain
// for debugging purposes.
func DebugModifier(m Modifier) string {
	if m == nil {
		return "<none>"
	}

	// Special case for emptyModifier
	if _, ok := m.(emptyModifier); ok {
		return "<none>"
	}

	var sb strings.Builder

	// First check if it's an inspectable modifier
	if im, ok := m.(InspectableModifier); ok {
		info := im.InspectorInfo()
		sb.WriteString(info.Name)
		if len(info.Properties) > 0 {
			sb.WriteString(" {")
			first := true
			for k, v := range info.Properties {
				if !first {
					sb.WriteString(", ")
				}
				// Special handling for color.Color values
				if c, ok := v.(color.Color); ok {
					r, g, b, a := c.RGBA()
					sb.WriteString(fmt.Sprintf("%s: color.RGBA{R:%d, G:%d, B:%d, A:%d}", k, r>>8, g>>8, b>>8, a>>8))
				} else {
					sb.WriteString(fmt.Sprintf("%s: %v", k, v))
				}
				first = false
			}
			sb.WriteString("}")
		}
		sb.WriteString(fmt.Sprintf("  // %T", m))
		return sb.String()
	}

	// Use Fold to traverse the modifier chain for chained modifiers
	m = UnwrapModifier(m)
	if chain, ok := m.(ModifierChain); ok {
		// Count total modifiers first for better formatting
		count := 0
		chain.Fold(nil, func(interface{}, Modifier) interface{} {
			count++
			return nil
		})

		sb.WriteString(fmt.Sprintf("Modifiers (%d):\n", count))

		i := 0
		chain.Fold(nil, func(_ interface{}, mod Modifier) interface{} {
			if mod == nil {
				sb.WriteString("  • <nil>\n")
				return nil
			}

			sb.WriteString(fmt.Sprintf("  %d. ", i+1))
			i++

			// Get the full type name with package
			modType := fmt.Sprintf("%T", mod)
			// Get short type name (without package)
			shortType := modType
			if dot := strings.LastIndex(shortType, "."); dot != -1 {
				shortType = shortType[dot+1:]
			}

			// If it's an inspectable modifier, use its name
			if im, ok := mod.(InspectableModifier); ok {
				info := im.InspectorInfo()
				sb.WriteString(info.Name)
				if len(info.Properties) > 0 {
					sb.WriteString(" {")
					first := true
					for k, v := range info.Properties {
						if !first {
							sb.WriteString(", ")
						}
						// Special handling for color.Color values
						if c, ok := v.(color.Color); ok {
							r, g, b, a := c.RGBA()
							sb.WriteString(fmt.Sprintf("%s: color.RGBA{R:%d, G:%d, B:%d, A:%d}", k, r>>8, g>>8, b>>8, a>>8))
						} else {
							sb.WriteString(fmt.Sprintf("%s: %v", k, v))
						}
						first = false
					}
					sb.WriteString("}")
				}
			} else {
				sb.WriteString(shortType)
			}

			// Show the full type in a comment for debugging
			sb.WriteString(fmt.Sprintf("  // %s\n", modType))
			return nil
		})
	} else {
		// Handle non-chain modifiers
		modType := fmt.Sprintf("%T", m)
		shortType := modType
		if dot := strings.LastIndex(shortType, "."); dot != -1 {
			shortType = shortType[dot+1:]
		}
		sb.WriteString(fmt.Sprintf("Modifiers (1):\n  1. %s  // %s\n", shortType, modType))
	}

	return strings.TrimSpace(sb.String())
}

// UnwrapModifier extracts the underlying Modifier from wrapper types
func UnwrapModifier(m Modifier) Modifier {
	if m == nil {
		return nil
	}

	// Handle inspectable modifier wrapper
	if im, ok := m.(InspectableModifier); ok {
		return im.(*inspectableModifier).Modifier
	}

	return m
}
