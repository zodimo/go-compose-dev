package style

// Automatic hyphenation configuration.
//
// Hyphenation is a dash-like punctuation mark used to join two-words into one or separate
// syl-lab-les of a word.
//
// Automatic hyphenation is added between syllables at appropriate hyphenation points, following
// language rules.
//
// However, user can override automatic break point selection, suggesting line break opportunities
// (see Suggesting line break opportunities below).
//
// Suggesting line break opportunities:
//   - "\u2010" ("hard" hyphen) Indicates a visible line break opportunity. Even if the
//     line is not actually broken at that point, the hyphen is still rendered.
//   - "\u00AD" ("soft" hyphen) This character is not rendered visibly; instead, it marks a
//     place where the word can be broken if hyphenation is necessary.
//
// The default configuration for [Hyphens] = [HyphensNone]
type Hyphens int

const (
	// This represents an unset value, a usual replacement for "null" when a primitive value is
	// desired.
	HyphensUnspecified Hyphens = 0

	// Lines will break with no hyphenation.
	//
	// "Hard" hyphens will still be respected. However, no automatic hyphenation will be
	// attempted. If a word must be broken due to being longer than a line, it will break at any
	// character and will not attempt to break at a syllable boundary.
	//
	// +---------+
	// | Experim |
	// | ental   |
	// +---------+
	HyphensNone Hyphens = 1

	// The words will be automatically broken at appropriate hyphenation points.
	//
	// However, suggested line break opportunities (see Suggesting line break opportunities
	// above) will override automatic break point selection when present.
	//
	// +---------+
	// | Experi- |
	// | mental  |
	// +---------+
	HyphensAuto Hyphens = 2
)

func (h Hyphens) String() string {
	switch h {
	case HyphensNone:
		return "Hyphens.None"
	case HyphensAuto:
		return "Hyphens.Auto"
	case HyphensUnspecified:
		return "Hyphens.Unspecified"
	default:
		return "Invalid"
	}
}

// Returns true if it is not [HyphensUnspecified].
func (h Hyphens) IsSpecified() bool {
	return h != HyphensUnspecified
}

// If [IsSpecified] is true then this is returned, otherwise [block] is executed and its result is
// returned.
func (h Hyphens) TakeOrElse(block func() Hyphens) Hyphens {
	if h.IsSpecified() {
		return h
	}
	return block()
}
