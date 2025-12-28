package internal

// GapBuffer provides efficient text editing via a gap-based data structure.
//
// The gap buffer maintains a "gap" (empty space) in the buffer that moves to the
// insertion/deletion point. This provides O(1) operations near the cursor while
// maintaining O(n) worst case for distant operations.
//
// This is a port of androidx.compose.foundation.text.input.internal.GapBuffer.
type GapBuffer struct {
	buffer   []rune
	gapStart int
	gapEnd   int
}

const (
	// defaultGapSize is the initial/minimum gap size for efficient editing.
	defaultGapSize = 256
)

// NewGapBuffer creates a new GapBuffer with the given initial text.
func NewGapBuffer(text string) *GapBuffer {
	runes := []rune(text)
	bufferLen := len(runes) + defaultGapSize
	buffer := make([]rune, bufferLen)
	copy(buffer, runes)
	return &GapBuffer{
		buffer:   buffer,
		gapStart: len(runes),
		gapEnd:   bufferLen,
	}
}

// Length returns the logical length of the text (excluding the gap).
func (g *GapBuffer) Length() int {
	return len(g.buffer) - (g.gapEnd - g.gapStart)
}

// Get returns the rune at the given index.
func (g *GapBuffer) Get(index int) rune {
	if index < g.gapStart {
		return g.buffer[index]
	}
	return g.buffer[index+(g.gapEnd-g.gapStart)]
}

// Delete removes characters from start (inclusive) to end (exclusive).
func (g *GapBuffer) Delete(start, end int) {
	if start == end {
		return
	}
	g.moveGapTo(start)
	g.gapEnd += end - start
}

// Replace replaces text from start to end with the given text.
func (g *GapBuffer) Replace(start, end int, text string) {
	newRunes := []rune(text)
	deleteLen := end - start
	insertLen := len(newRunes)

	// Move gap to insertion point
	g.moveGapTo(start)

	// Delete the range
	g.gapEnd += deleteLen

	// Ensure sufficient gap size for insertion
	g.ensureGapSize(insertLen)

	// Insert new text
	copy(g.buffer[g.gapStart:], newRunes)
	g.gapStart += insertLen
}

// Append adds text at the end.
func (g *GapBuffer) Append(text string) {
	g.Replace(g.Length(), g.Length(), text)
}

// Insert adds text at the given index.
func (g *GapBuffer) Insert(index int, text string) {
	g.Replace(index, index, text)
}

// String returns the complete text content.
func (g *GapBuffer) String() string {
	result := make([]rune, g.Length())
	// Copy before gap
	copy(result, g.buffer[:g.gapStart])
	// Copy after gap
	copy(result[g.gapStart:], g.buffer[g.gapEnd:])
	return string(result)
}

// SubSequence returns a substring from start (inclusive) to end (exclusive).
func (g *GapBuffer) SubSequence(start, end int) string {
	length := end - start
	result := make([]rune, length)
	for i := 0; i < length; i++ {
		result[i] = g.Get(start + i)
	}
	return string(result)
}

// moveGapTo moves the gap to the specified position.
func (g *GapBuffer) moveGapTo(position int) {
	if position == g.gapStart {
		return
	}

	gapSize := g.gapEnd - g.gapStart

	if position < g.gapStart {
		// Move gap left: shift content right
		distance := g.gapStart - position
		copy(g.buffer[g.gapEnd-distance:g.gapEnd], g.buffer[position:g.gapStart])
		g.gapStart = position
		g.gapEnd = position + gapSize
	} else {
		// Move gap right: shift content left
		distance := position - g.gapStart
		copy(g.buffer[g.gapStart:g.gapStart+distance], g.buffer[g.gapEnd:g.gapEnd+distance])
		g.gapStart = position
		g.gapEnd = position + gapSize
	}
}

// ensureGapSize ensures the gap is at least the specified size.
func (g *GapBuffer) ensureGapSize(needed int) {
	gapSize := g.gapEnd - g.gapStart
	if gapSize >= needed {
		return
	}

	// Need to grow - calculate new size
	newGapSize := needed
	if newGapSize < defaultGapSize {
		newGapSize = defaultGapSize
	}

	// Create larger buffer
	oldLen := len(g.buffer)
	newLen := oldLen + newGapSize - gapSize
	newBuffer := make([]rune, newLen)

	// Copy before gap
	copy(newBuffer, g.buffer[:g.gapStart])
	// Copy after gap (at new position)
	newGapEnd := g.gapStart + newGapSize
	copy(newBuffer[newGapEnd:], g.buffer[g.gapEnd:])

	g.buffer = newBuffer
	g.gapEnd = newGapEnd
}

// ContentEquals returns true if the content equals the given string.
func (g *GapBuffer) ContentEquals(other string) bool {
	runes := []rune(other)
	if len(runes) != g.Length() {
		return false
	}
	for i, r := range runes {
		if g.Get(i) != r {
			return false
		}
	}
	return true
}

// PartialGapBuffer wraps text with a localized gap buffer around the cursor.
//
// Unlike GapBuffer, this implementation doesn't convert the entire text to a mutable
// buffer. Instead, it only converts the cursor-adjacent region, saving construction
// time and memory for large texts. If editing moves away from the buffer region,
// the buffer is flushed and a new region is created.
//
// This is a port of androidx.compose.foundation.text.input.internal.PartialGapBuffer.
type PartialGapBuffer struct {
	text     string     // Original text (immutable regions)
	buffer   *GapBuffer // Mutable region around cursor (nil if not active)
	bufStart int        // Start of buffer region in the current text
	bufEnd   int        // End of buffer region in the current text
}

const (
	// bufferExtent is the number of characters to include before/after the edit point.
	bufferExtent = 256
)

// NewPartialGapBuffer creates a new PartialGapBuffer with the given text.
func NewPartialGapBuffer(text string) *PartialGapBuffer {
	return &PartialGapBuffer{
		text: text,
	}
}

// Length returns the logical length of the text.
func (p *PartialGapBuffer) Length() int {
	if p.buffer == nil {
		return len([]rune(p.text))
	}
	// Original length = len(text) in runes
	// Remove buffer region length, add gap buffer content length
	originalRunes := len([]rune(p.text))
	bufferRegionLen := p.bufEnd - p.bufStart
	return originalRunes - bufferRegionLen + p.buffer.Length()
}

// Get returns the rune at the given index.
func (p *PartialGapBuffer) Get(index int) rune {
	if p.buffer == nil {
		return []rune(p.text)[index]
	}

	if index < p.bufStart {
		return []rune(p.text)[index]
	}

	bufferLen := p.buffer.Length()
	bufferEndInResult := p.bufStart + bufferLen

	if index < bufferEndInResult {
		return p.buffer.Get(index - p.bufStart)
	}

	// After buffer region
	textRunes := []rune(p.text)
	return textRunes[p.bufEnd+(index-bufferEndInResult)]
}

// Replace replaces text from start to end with the given string.
func (p *PartialGapBuffer) Replace(start, end int, replacement string) {
	if p.buffer == nil {
		// Initialize buffer around the edit point
		p.initializeBuffer(start, end)
	} else if !p.isNearBuffer(start, end) {
		// Flush and reinitialize if edit is far from current buffer
		p.flush()
		p.initializeBuffer(start, end)
	}

	// Adjust indices for buffer-local addressing
	localStart := start - p.bufStart
	localEnd := end - p.bufStart

	p.buffer.Replace(localStart, localEnd, replacement)
}

// String returns the complete text content.
func (p *PartialGapBuffer) String() string {
	if p.buffer == nil {
		return p.text
	}
	p.flush()
	return p.text
}

// SubSequence returns a substring from start (inclusive) to end (exclusive).
func (p *PartialGapBuffer) SubSequence(start, end int) string {
	length := end - start
	result := make([]rune, length)
	for i := 0; i < length; i++ {
		result[i] = p.Get(start + i)
	}
	return string(result)
}

// ContentEquals returns true if the content equals the given string.
func (p *PartialGapBuffer) ContentEquals(other string) bool {
	if p.buffer == nil {
		return p.text == other
	}
	// Need to compare character by character
	runes := []rune(other)
	if len(runes) != p.Length() {
		return false
	}
	for i, r := range runes {
		if p.Get(i) != r {
			return false
		}
	}
	return true
}

// initializeBuffer creates a gap buffer around the edit region.
func (p *PartialGapBuffer) initializeBuffer(start, end int) {
	textRunes := []rune(p.text)
	textLen := len(textRunes)

	// Calculate buffer region with extent
	p.bufStart = start - bufferExtent
	if p.bufStart < 0 {
		p.bufStart = 0
	}

	p.bufEnd = end + bufferExtent
	if p.bufEnd > textLen {
		p.bufEnd = textLen
	}

	// Extract the region and create gap buffer
	regionText := string(textRunes[p.bufStart:p.bufEnd])
	p.buffer = NewGapBuffer(regionText)
}

// flush converts the buffer back to a string and rebuilds the text.
func (p *PartialGapBuffer) flush() {
	if p.buffer == nil {
		return
	}

	textRunes := []rune(p.text)
	bufferContent := p.buffer.String()

	// Rebuild: before + buffer content + after
	var result []rune
	result = append(result, textRunes[:p.bufStart]...)
	result = append(result, []rune(bufferContent)...)
	result = append(result, textRunes[p.bufEnd:]...)

	p.text = string(result)
	p.buffer = nil
	p.bufStart = 0
	p.bufEnd = 0
}

// isNearBuffer returns true if the edit range overlaps with or is adjacent to the buffer.
func (p *PartialGapBuffer) isNearBuffer(start, end int) bool {
	bufferContentEnd := p.bufStart + p.buffer.Length()
	return start <= bufferContentEnd && end >= p.bufStart
}
