// Package text provides rich text styling and annotation capabilities.
package text

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// ============================================================================
// Core Types
// ============================================================================

// AnnotatedString is the basic data structure of text with multiple styles.
// It's immutable; use Builder to construct instances.
type AnnotatedString struct {
	text        string
	annotations []Range[Annotation]

	// Cached derived slices (nil if empty)
	spanStyles      []Range[SpanStyle]
	paragraphStyles []Range[ParagraphStyle]
}

// Range represents an annotation applied to a text range [Start, End).
type Range[T any] struct {
	Item  T
	Start int
	End   int
	Tag   string
}

// NewRange creates a Range with the given item and bounds.
func NewRange[T any](item T, start, end int) Range[T] {
	return Range[T]{Item: item, Start: start, End: end}
}

// NewRangeWithTag creates a Range with an additional tag for identification.
func NewRangeWithTag[T any](item T, start, end int, tag string) Range[T] {
	return Range[T]{Item: item, Start: start, End: end, Tag: tag}
}

// Annotation is a marker interface for all annotation types.
type Annotation interface {
	isAnnotation()
}

// ============================================================================
// Annotation Implementations
// ============================================================================

// SpanStyle specifies character-level styling (color, font, etc.).
type SpanStyle struct {
	// Define fields as needed (Color, FontSize, etc.)
}

func (SpanStyle) isAnnotation() {}

// ParagraphStyle specifies paragraph-level styling (alignment, indent, etc.).
type ParagraphStyle struct {
	// Define fields as needed
}

func (ParagraphStyle) isAnnotation() {}

// StringAnnotation is a simple string metadata attachment.
type StringAnnotation string

func (StringAnnotation) isAnnotation() {}

// TtsAnnotation provides text-to-speech metadata.
type TtsAnnotation struct {
	// Define TTS-specific fields
}

func (TtsAnnotation) isAnnotation() {}

// LinkAnnotation marks clickable links in text.
type LinkAnnotation struct {
	URL string
}

func (LinkAnnotation) isAnnotation() {}

// Bullet represents a bullet list marker.
type Bullet struct {
	// Define bullet styling fields
}

func (Bullet) isAnnotation() {}

// ============================================================================
// Constructors
// ============================================================================

// NewAnnotatedString creates an AnnotatedString from text and optional style ranges.
func NewAnnotatedString(text string, spanStyles []Range[SpanStyle], paragraphStyles []Range[ParagraphStyle]) AnnotatedString {
	annotations := constructAnnotations(spanStyles, paragraphStyles)
	return newAnnotatedStringWithAnnotations(text, annotations)
}

// newAnnotatedStringWithAnnotations creates an AnnotatedString from text and generic annotations.
func newAnnotatedStringWithAnnotations(text string, annotations []Range[Annotation]) AnnotatedString {
	// Validate paragraph styles don't overlap
	if err := validateParagraphStyles(paragraphStylesFromAnnotations(annotations)); err != nil {
		panic(err)
	}

	as := AnnotatedString{
		text:        text,
		annotations: annotations,
	}

	// Build cached slices
	var spanStyles []Range[SpanStyle]
	var paraStyles []Range[ParagraphStyle]

	for _, ann := range annotations {
		switch v := ann.Item.(type) {
		case SpanStyle:
			if spanStyles == nil {
				spanStyles = make([]Range[SpanStyle], 0, 4)
			}
			spanStyles = append(spanStyles, Range[SpanStyle]{
				Item:  v,
				Start: ann.Start,
				End:   ann.End,
				Tag:   ann.Tag,
			})
		case ParagraphStyle:
			if paraStyles == nil {
				paraStyles = make([]Range[ParagraphStyle], 0, 4)
			}
			paraStyles = append(paraStyles, Range[ParagraphStyle]{
				Item:  v,
				Start: ann.Start,
				End:   ann.End,
				Tag:   ann.Tag,
			})
		}
	}

	as.spanStyles = spanStyles
	as.paragraphStyles = paraStyles
	return as
}

// constructAnnotations merges SpanStyle and ParagraphStyle slices into a unified annotation list.
func constructAnnotations(spanStyles []Range[SpanStyle], paragraphStyles []Range[ParagraphStyle]) []Range[Annotation] {
	totalLen := len(spanStyles) + len(paragraphStyles)
	if totalLen == 0 {
		return nil
	}

	annotations := make([]Range[Annotation], 0, totalLen)
	for _, s := range spanStyles {
		annotations = append(annotations, Range[Annotation]{
			Item:  s.Item,
			Start: s.Start,
			End:   s.End,
			Tag:   s.Tag,
		})
	}
	for _, p := range paragraphStyles {
		annotations = append(annotations, Range[Annotation]{
			Item:  p.Item,
			Start: p.Start,
			End:   p.End,
			Tag:   p.Tag,
		})
	}
	return annotations
}

// paragraphStylesFromAnnotations extracts ParagraphStyle ranges from generic annotations.
func paragraphStylesFromAnnotations(annotations []Range[Annotation]) []Range[ParagraphStyle] {
	var paraStyles []Range[ParagraphStyle]
	for _, ann := range annotations {
		if ps, ok := ann.Item.(ParagraphStyle); ok {
			paraStyles = append(paraStyles, Range[ParagraphStyle]{
				Item:  ps,
				Start: ann.Start,
				End:   ann.End,
				Tag:   ann.Tag,
			})
		}
	}
	return paraStyles
}

// validateParagraphStyles ensures no overlapping paragraph ranges.
func validateParagraphStyles(styles []Range[ParagraphStyle]) error {
	if len(styles) < 2 {
		return nil
	}

	// Sort by Start
	sorted := make([]Range[ParagraphStyle], len(styles))
	copy(sorted, styles)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Start < sorted[j].Start
	})

	// Track nested paragraph ends
	ends := make([]int, 0, 4)
	for _, s := range sorted {
		for len(ends) > 0 && s.Start >= ends[len(ends)-1] {
			ends = ends[:len(ends)-1] // Pop finished paragraphs
		}

		if len(ends) > 0 {
			// Current must be fully contained in last paragraph
			if s.End > ends[len(ends)-1] {
				return fmt.Errorf("paragraph overlap not allowed: range %d-%d exceeds parent %d",
					s.Start, s.End, ends[len(ends)-1])
			}
		}
		ends = append(ends, s.End)
	}
	return nil
}

// ============================================================================
// Properties
// ============================================================================

// Len returns the length of the text.
func (as AnnotatedString) Len() int {
	return len(as.text)
}

// String implements fmt.Stringer and returns the plain text.
func (as AnnotatedString) String() string {
	return as.text
}

// CharAt returns the character at the given index.
func (as AnnotatedString) CharAt(index int) byte {
	return as.text[index]
}

// Text returns the plain text.
func (as AnnotatedString) Text() string {
	return as.text
}

// SpanStyles returns all SpanStyle ranges.
func (as AnnotatedString) SpanStyles() []Range[SpanStyle] {
	if as.spanStyles == nil {
		return []Range[SpanStyle]{}
	}
	return as.spanStyles
}

// ParagraphStyles returns all ParagraphStyle ranges.
func (as AnnotatedString) ParagraphStyles() []Range[ParagraphStyle] {
	if as.paragraphStyles == nil {
		return []Range[ParagraphStyle]{}
	}
	return as.paragraphStyles
}

// ============================================================================
// Operations
// ============================================================================

// SubSequence returns a substring with annotations in the range [start, end).
func (as AnnotatedString) SubSequence(start, end int) AnnotatedString {
	if start == 0 && end >= len(as.text) {
		return as
	}
	if start > end {
		panic(fmt.Sprintf("start (%d) must be <= end (%d)", start, end))
	}
	if start == end {
		return Empty()
	}

	return newAnnotatedStringWithAnnotations(
		as.text[start:end],
		filterRanges(as.annotations, start, end),
	)
}

// Append concatenates another AnnotatedString.
func (as AnnotatedString) Append(other AnnotatedString) AnnotatedString {
	b := NewBuilder()
	b.AppendAnnotatedString(as)
	b.AppendAnnotatedString(other)
	return b.ToAnnotatedString()
}

// GetStringAnnotations queries string annotations by tag within a range.
func (as AnnotatedString) GetStringAnnotations(tag string, start, end int) []Range[string] {
	if start >= end {
		return nil
	}

	var result []Range[string]
	for _, ann := range as.annotations {
		if sa, ok := ann.Item.(StringAnnotation); ok && ann.Tag == tag && intersect(start, end, ann.Start, ann.End) {
			result = append(result, Range[string]{
				Item:  string(sa),
				Start: max(start, ann.Start) - start,
				End:   min(end, ann.End) - start,
				Tag:   ann.Tag,
			})
		}
	}
	return result
}

// HasStringAnnotations checks if any string annotations with the given tag exist in range.
func (as AnnotatedString) HasStringAnnotations(tag string, start, end int) bool {
	for _, ann := range as.annotations {
		if _, ok := ann.Item.(StringAnnotation); ok && ann.Tag == tag && intersect(start, end, ann.Start, ann.End) {
			return true
		}
	}
	return false
}

// GetLinkAnnotations queries LinkAnnotations within a range.
func (as AnnotatedString) GetLinkAnnotations(start, end int) []Range[LinkAnnotation] {
	if start >= end {
		return nil
	}

	var result []Range[LinkAnnotation]
	for _, ann := range as.annotations {
		if la, ok := ann.Item.(LinkAnnotation); ok && intersect(start, end, ann.Start, ann.End) {
			result = append(result, Range[LinkAnnotation]{
				Item:  la,
				Start: max(start, ann.Start) - start,
				End:   min(end, ann.End) - start,
				Tag:   ann.Tag,
			})
		}
	}
	return result
}

// HasLinkAnnotations checks if any link annotations exist in range.
func (as AnnotatedString) HasLinkAnnotations(start, end int) bool {
	for _, ann := range as.annotations {
		if _, ok := ann.Item.(LinkAnnotation); ok && intersect(start, end, ann.Start, ann.End) {
			return true
		}
	}
	return false
}

// HasEqualAnnotations compares annotation lists for equality.
func (as AnnotatedString) HasEqualAnnotations(other AnnotatedString) bool {
	if len(as.annotations) != len(other.annotations) {
		return false
	}
	for i, ann := range as.annotations {
		if !rangeEqual(ann, other.annotations[i]) {
			return false
		}
	}
	return true
}

// rangeEqual compares two Range[Annotation] instances.
func rangeEqual(a, b Range[Annotation]) bool {
	return a.Start == b.Start && a.End == b.End && a.Tag == b.Tag
}

// Empty returns an empty AnnotatedString.
func Empty() AnnotatedString {
	return AnnotatedString{text: ""}
}

// ============================================================================
// Builder
// ============================================================================

// Builder constructs AnnotatedString instances.
type Builder struct {
	text        strings.Builder
	styleStack  []mutableRange
	annotations []mutableRange
}

type mutableRange struct {
	item  Annotation
	start int
	end   int // mutable, defaults to intMin
	tag   string
}

// NewBuilder creates a new Builder with optional initial capacity.
func NewBuilder(capacity ...int) *Builder {
	cap := 16
	if len(capacity) > 0 && capacity[0] > 0 {
		cap = capacity[0]
	}
	var builder strings.Builder
	builder.Grow(cap)

	return &Builder{
		text:        builder,
		styleStack:  make([]mutableRange, 0, 8),
		annotations: make([]mutableRange, 0, 16),
	}
}

// Len returns the current text length.
func (b *Builder) Len() int {
	return b.text.Len()
}

// Append adds plain text.
func (b *Builder) Append(s string) {
	b.text.WriteString(s)
}

// AppendAnnotatedString adds another AnnotatedString with offset annotations.
func (b *Builder) AppendAnnotatedString(as AnnotatedString) {
	start := b.text.Len()
	b.text.WriteString(as.text)

	offset := start
	for _, ann := range as.annotations {
		b.annotations = append(b.annotations, mutableRange{
			item:  ann.Item,
			start: offset + ann.Start,
			end:   offset + ann.End,
			tag:   ann.Tag,
		})
	}
}

// AddStyle adds a SpanStyle to the given range.
func (b *Builder) AddStyle(style SpanStyle, start, end int) {
	b.annotations = append(b.annotations, mutableRange{
		item:  style,
		start: start,
		end:   end,
	})
}

// AddParagraphStyle adds a ParagraphStyle to the given range.
func (b *Builder) AddParagraphStyle(style ParagraphStyle, start, end int) {
	b.annotations = append(b.annotations, mutableRange{
		item:  style,
		start: start,
		end:   end,
	})
}

// AddStringAnnotation adds a string annotation with a tag.
func (b *Builder) AddStringAnnotation(tag, annotation string, start, end int) {
	b.annotations = append(b.annotations, mutableRange{
		item:  StringAnnotation(annotation),
		start: start,
		end:   end,
		tag:   tag,
	})
}

// PushStyle pushes a SpanStyle onto the stack for subsequent text.
func (b *Builder) PushStyle(style SpanStyle) int {
	mr := mutableRange{
		item:  style,
		start: b.text.Len(),
		end:   math.MinInt, // marker for unset
	}
	b.styleStack = append(b.styleStack, mr)
	b.annotations = append(b.annotations, mr)
	return len(b.styleStack) - 1
}

// Pop ends the last pushed style/annotation.
func (b *Builder) Pop() {
	if len(b.styleStack) == 0 {
		panic("nothing to pop")
	}
	last := &b.styleStack[len(b.styleStack)-1]
	last.end = b.text.Len()
	b.styleStack = b.styleStack[:len(b.styleStack)-1]
}

// PopTo pops all styles up to and including the given index.
func (b *Builder) PopTo(index int) {
	if index >= len(b.styleStack) {
		panic(fmt.Sprintf("index %d exceeds stack size %d", index, len(b.styleStack)))
	}
	for len(b.styleStack) > index {
		b.Pop()
	}
}

// ToAnnotatedString builds the final AnnotatedString.
func (b *Builder) ToAnnotatedString() AnnotatedString {
	// Close any open ranges
	for i := range b.annotations {
		if b.annotations[i].end == math.MinInt {
			b.annotations[i].end = b.text.Len()
		}
	}
	return newAnnotatedStringWithAnnotations(b.text.String(), mutableToImmutable(b.annotations))
}

// mutableToImmutable converts mutableRange slices to Range[Annotation].
func mutableToImmutable(mr []mutableRange) []Range[Annotation] {
	if len(mr) == 0 {
		return nil
	}
	result := make([]Range[Annotation], len(mr))
	for i, r := range mr {
		result[i] = Range[Annotation]{
			Item:  r.item,
			Start: r.start,
			End:   r.end,
			Tag:   r.tag,
		}
	}
	return result
}

// ============================================================================
// Paragraph Normalization
// ============================================================================

// NormalizedParagraphStyles returns paragraph styles with gaps filled by default style.
// This implements the complex merging logic from the Kotlin version.
func (as AnnotatedString) NormalizedParagraphStyles(defaultStyle ParagraphStyle) []Range[ParagraphStyle] {
	paraStyles := as.ParagraphStyles()
	if len(paraStyles) == 0 {
		return []Range[ParagraphStyle]{
			{Item: defaultStyle, Start: 0, End: len(as.text)},
		}
	}

	sort.Slice(paraStyles, func(i, j int) bool {
		return paraStyles[i].Start < paraStyles[j].Start
	})

	var result []Range[ParagraphStyle]
	lastAdded := 0
	var stack []Range[ParagraphStyle]

	for _, current := range paraStyles {
		merged := current.Item // TODO: Actually merge with defaultStyle
		currentRange := Range[ParagraphStyle]{
			Item:  merged,
			Start: current.Start,
			End:   current.End,
		}

		// Fill gaps
		for lastAdded < currentRange.Start && len(stack) > 0 {
			lastStack := stack[len(stack)-1]
			if currentRange.Start < lastStack.End {
				result = append(result, Range[ParagraphStyle]{
					Item:  lastStack.Item,
					Start: lastAdded,
					End:   currentRange.Start,
				})
				lastAdded = currentRange.Start
				break
			}
			result = append(result, lastStack)
			lastAdded = lastStack.End
			stack = stack[:len(stack)-1]
		}

		if lastAdded < currentRange.Start {
			result = append(result, Range[ParagraphStyle]{
				Item:  defaultStyle,
				Start: lastAdded,
				End:   currentRange.Start,
			})
			lastAdded = currentRange.Start
		}

		// Handle nesting/overlapping
		if len(stack) > 0 {
			lastStack := stack[len(stack)-1]
			if lastStack.Start == currentRange.Start && lastStack.End == currentRange.End {
				// Fully overlapped - merge and replace
				stack[len(stack)-1] = currentRange
			} else if lastStack.End < currentRange.End {
				panic("paragraph overlap not allowed")
			} else {
				stack = append(stack, currentRange)
			}
		} else {
			stack = append(stack, currentRange)
		}
	}

	// Empty remaining stack
	for len(stack) > 0 {
		last := stack[len(stack)-1]
		result = append(result, last)
		lastAdded = last.End
		stack = stack[:len(stack)-1]
	}

	// Add trailing default style
	if lastAdded < len(as.text) {
		result = append(result, Range[ParagraphStyle]{
			Item:  defaultStyle,
			Start: lastAdded,
			End:   len(as.text),
		})
	}

	return result
}

// ============================================================================
// Utility Functions
// ============================================================================

// filterRanges extracts and offsets annotations for subsequence operations.
func filterRanges(annotations []Range[Annotation], start, end int) []Range[Annotation] {
	if len(annotations) == 0 || start >= end {
		return nil
	}

	var result []Range[Annotation]
	for _, ann := range annotations {
		if intersect(start, end, ann.Start, ann.End) {
			result = append(result, Range[Annotation]{
				Item:  ann.Item,
				Start: max(start, ann.Start) - start,
				End:   min(end, ann.End) - start,
				Tag:   ann.Tag,
			})
		}
	}
	return result
}

// intersect checks if two ranges [aStart, aEnd) and [bStart, bEnd) intersect.
func intersect(aStart, aEnd, bStart, bEnd int) bool {
	return (aStart < aEnd && bStart < bEnd && aStart < bEnd && bStart < aEnd) ||
		(aStart == aEnd && aStart >= bStart && aStart <= bEnd) ||
		(bStart == bEnd && bStart >= aStart && bStart <= aEnd)
}

// contains checks if [outerStart, outerEnd) fully contains [innerStart, innerEnd).
func contains(outerStart, outerEnd, innerStart, innerEnd int) bool {
	return outerStart <= innerStart && innerEnd <= outerEnd &&
		(outerEnd != innerEnd || (innerStart == innerEnd) == (outerStart == outerEnd))
}

// max returns the larger of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min returns the smaller of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
