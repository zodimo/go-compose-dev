package unit

import (
	"fmt"
	"math"
)

// Constraints represents immutable constraints for measuring layouts.
// It packs minWidth, maxWidth, minHeight, and maxHeight into a uint64.
type Constraints uint64

const (
	// Infinity is used when a constraint is considered unbounded.
	Infinity int = math.MaxInt32

	focusMask uint64 = 0x3
)

// Constants for bit allocation
const (
	minFocusBits                 = 16
	maxAllowedForMinFocusBits    = (1 << (31 - minFocusBits)) - 2
	minFocusMask                 = 0xFFFF
	minNonFocusBits              = 15
	maxAllowedForMinNonFocusBits = (1 << (31 - minNonFocusBits)) - 2
	minNonFocusMask              = 0x7FFF
	maxFocusBits                 = 18
	maxAllowedForMaxFocusBits    = (1 << (31 - maxFocusBits)) - 2
	maxFocusMask                 = 0x3FFFF
	maxNonFocusBits              = 13
	maxAllowedForMaxNonFocusBits = (1 << (31 - maxNonFocusBits)) - 2
	maxNonFocusMask              = 0x1FFF
	maxDimensionsAndFocusMask    = 0xFFFFFFFE00000003
)

// --- Construction ---

// NewConstraints creates a new Constraints object.
func NewConstraints(minWidth, maxWidth, minHeight, maxHeight int) Constraints {
	if !(maxWidth >= minWidth && maxHeight >= minHeight && minWidth >= 0 && minHeight >= 0) {
		panic(fmt.Sprintf("Invalid constraints: w(%d-%d), h(%d-%d)", minWidth, maxWidth, minHeight, maxHeight))
	}
	return createConstraints(minWidth, maxWidth, minHeight, maxHeight)
}

func Fixed(width, height int) Constraints {
	return NewConstraints(width, width, height, height)
}

func FixedWidth(width int) Constraints {
	return NewConstraints(width, width, 0, Infinity)
}

func FixedHeight(height int) Constraints {
	return NewConstraints(0, Infinity, height, height)
}

// --- Getters ---

func (c Constraints) focusIndex() int {
	return int(uint64(c) & focusMask)
}

func (c Constraints) MinWidth() int {
	mask := widthMask(indexToBitOffset(c.focusIndex()))
	return int((uint64(c) >> 2)) & int(mask)
}

func (c Constraints) MaxWidth() int {
	mask := widthMask(indexToBitOffset(c.focusIndex()))
	width := int((uint64(c) >> 33)) & int(mask)
	if width == 0 {
		return Infinity
	}
	return width - 1
}

func (c Constraints) MinHeight() int {
	bitOffset := indexToBitOffset(c.focusIndex())
	mask := heightMask(bitOffset)
	offset := minHeightOffsets(bitOffset)
	return int((uint64(c) >> uint(offset))) & int(mask)
}

func (c Constraints) MaxHeight() int {
	bitOffset := indexToBitOffset(c.focusIndex())
	mask := heightMask(bitOffset)
	offset := minHeightOffsets(bitOffset) + 31
	height := int((uint64(c) >> uint(offset))) & int(mask)
	if height == 0 {
		return Infinity
	}
	return height - 1
}

// --- Bounded Checks ---

func (c Constraints) HasBoundedWidth() bool {
	mask := widthMask(indexToBitOffset(c.focusIndex()))
	return (int(uint64(c)>>33) & int(mask)) != 0
}

func (c Constraints) HasBoundedHeight() bool {
	bitOffset := indexToBitOffset(c.focusIndex())
	mask := heightMask(bitOffset)
	offset := minHeightOffsets(bitOffset) + 31
	return (int(uint64(c)>>uint(offset)) & int(mask)) != 0
}

func (c Constraints) HasFixedWidth() bool  { return c.MinWidth() == c.MaxWidth() }
func (c Constraints) HasFixedHeight() bool { return c.MinHeight() == c.MaxHeight() }
func (c Constraints) IsZero() bool         { return c.MaxWidth() == 0 || c.MaxHeight() == 0 }

// --- Utilities ---

func (c Constraints) Copy(minWidth, maxWidth, minHeight, maxHeight int) Constraints {
	return NewConstraints(minWidth, maxWidth, minHeight, maxHeight)
}

func (c Constraints) CopyMaxDimensions() Constraints {
	return Constraints(uint64(c) & maxDimensionsAndFocusMask)
}

func (c Constraints) ConstrainWidth(width int) int {
	return fastCoerceIn(width, c.MinWidth(), c.MaxWidth())
}

func (c Constraints) ConstrainHeight(height int) int {
	return fastCoerceIn(height, c.MinHeight(), c.MaxHeight())
}

// --- Internal Helpers ---

func createConstraints(minWidth, maxWidth, minHeight, maxHeight int) Constraints {
	heightVal := maxHeight
	if maxHeight == Infinity {
		heightVal = minHeight
	}
	heightBits := bitsNeededForSizeUnchecked(heightVal)

	widthVal := maxWidth
	if maxWidth == Infinity {
		widthVal = minWidth
	}
	widthBits := bitsNeededForSizeUnchecked(widthVal)

	if widthBits+heightBits > 31 {
		panic(fmt.Sprintf("Can't represent width %d and height %d in Constraints", widthVal, heightVal))
	}

	// Branchless max(0, x) logic mimicking the Kotlin implementation
	maxWidthValue := maxWidth + 1
	if maxWidth == Infinity {
		maxWidthValue = 0
	}

	maxHeightValue := maxHeight + 1
	if maxHeight == Infinity {
		maxHeightValue = 0
	}

	bitOffset := widthBits - 13
	focus := bitOffsetToIndex(bitOffset)

	minHeightOffset := minHeightOffsets(bitOffset)
	maxHeightOffset := minHeightOffset + 31

	value := uint64(focus) |
		(uint64(minWidth) << 2) |
		(uint64(maxWidthValue) << 33) |
		(uint64(minHeight) << uint(minHeightOffset)) |
		(uint64(maxHeightValue) << uint(maxHeightOffset))

	return Constraints(value)
}

func bitsNeededForSizeUnchecked(size int) int {
	switch {
	case size < maxNonFocusMask:
		return maxNonFocusBits
	case size < minNonFocusMask:
		return minNonFocusBits
	case size < minFocusMask:
		return minFocusBits
	case size < maxFocusMask:
		return maxFocusBits
	default:
		return 255
	}
}

func indexToBitOffset(index int) int {
	return (index&0x1)<<1 + ((index & 0x2) >> 1 * 3)
}

func bitOffsetToIndex(bits int) int {
	return (bits >> 1) + (bits & 0x1)
}

func minHeightOffsets(bitOffset int) int {
	return 15 + bitOffset
}

func widthMask(bitOffset int) uint32 {
	return uint32((1 << (13 + uint(bitOffset))) - 1)
}

func heightMask(bitOffset int) uint32 {
	return uint32((1 << (18 - uint(bitOffset))) - 1)
}

func fastCoerceIn(value, minVal, maxVal int) int {
	if value < minVal {
		return minVal
	}
	if value > maxVal {
		return maxVal
	}
	return value
}

// --- Fitting Logic ---

// FitPrioritizingWidth returns Constraints that match as close as possible to the values passed.
// If the dimensions are outside of those that can be represented, the constraints are limited.
// The width is granted as much space as it needs (up to 18 bits) and the height is given the remaining space.
func FitPrioritizingWidth(minWidth, maxWidth, minHeight, maxHeight int) Constraints {
	minW := min(minWidth, maxFocusMask-1)
	maxW := minW
	if maxWidth == Infinity {
		maxW = Infinity
	} else {
		maxW = min(maxWidth, maxFocusMask-1)
	}

	consumed := minW
	if maxW != Infinity {
		consumed = maxW
	}

	maxAllowed := maxAllowedForSize(consumed)
	maxH := maxAllowed
	if maxHeight != Infinity {
		maxH = min(maxAllowed, maxHeight)
	}
	minH := min(maxAllowed, minHeight)

	return NewConstraints(minW, maxW, minH, maxH)
}

// FitPrioritizingHeight returns Constraints that match as close as possible to the values passed.
// The height is granted as much space as it needs (up to 18 bits) and the width is given the remaining space.
func FitPrioritizingHeight(minWidth, maxWidth, minHeight, maxHeight int) Constraints {
	minH := min(minHeight, maxFocusMask-1)
	maxH := minH
	if maxHeight == Infinity {
		maxH = Infinity
	} else {
		maxH = min(maxHeight, maxFocusMask-1)
	}

	consumed := minH
	if maxH != Infinity {
		consumed = maxH
	}

	maxAllowed := maxAllowedForSize(consumed)
	maxW := maxAllowed
	if maxWidth != Infinity {
		maxW = min(maxAllowed, maxWidth)
	}
	minW := min(maxAllowed, minWidth)

	return NewConstraints(minW, maxW, minH, maxH)
}

// maxAllowedForSize returns the maximum possible value for the other dimension
// given the size of the current focused dimension.
func maxAllowedForSize(size int) int {
	switch {
	case size < maxNonFocusMask:
		return maxAllowedForMaxNonFocusBits
	case size < minNonFocusMask:
		return maxAllowedForMinNonFocusBits
	case size < minFocusMask:
		return maxAllowedForMinFocusBits
	case size < maxFocusMask:
		return maxAllowedForMaxFocusBits
	default:
		panic(fmt.Sprintf("Can't represent a size of %d in Constraints", size))
	}
}

// Helper for Go 1.21+ (if you are on an older version, use a custom min function)
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
