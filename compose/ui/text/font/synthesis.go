package font

type SynthesisFlag int

const (
	None       SynthesisFlag = 0
	WeightFlag SynthesisFlag = 1
	StyleFlag  SynthesisFlag = 2
	AllFlags   SynthesisFlag = 0xffff
)

func (f SynthesisFlag) String() string {
	switch f {
	case None:
		return "None"
	case WeightFlag:
		return "Weight"
	case StyleFlag:
		return "Style"
	case AllFlags:
		return "All"
	default:
		return "Invalid"
	}
}

// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-text/src/commonMain/kotlin/androidx/compose/ui/text/font/FontSynthesis.kt;drc=182f9f08cf26aaa426d03f39b6dcef1c33e9f41b;l=44
type FontSynthesis struct {
	flag SynthesisFlag
}

func (f FontSynthesis) String() string {
	return f.flag.String()
}

func (f FontSynthesis) IsWeightOn() bool {
	return f.flag == WeightFlag
}

func (f FontSynthesis) IsStyleOn() bool {
	return f.flag == StyleFlag
}

// Perform platform-specific font synthesis such as fake bold or fake italic.
// Platforms are not required to support synthesis, in which case they should return [typeface].
// Platforms that support synthesis should check [FontSynthesis.isWeightOn] and
// [FontSynthesis.isStyleOn] in this method before synthesizing bold or italic, respectively.
// @param typeface a platform-specific typeface
// @param font initial font that generated the typeface via loading
// @param requestedWeight app-requested weight (may be different than the font's weight)
// @param requestedStyle app-requested style (may be different than the font's style)
// @return a synthesized typeface, or the passed [typeface] if synthesis is not needed or supported.
// func (f FontSynthesis) SynthesizeTypeface(typeface any, font Font, requestedWeight FontWeight, requestedStyle FontStyle) any {
// 	return typeface
// }

func FontSynthesisNone() FontSynthesis {
	return FontSynthesis{flag: None}
}

func FontSynthesisWeight() FontSynthesis {
	return FontSynthesis{flag: WeightFlag}
}

func FontSynthesisStyle() FontSynthesis {
	return FontSynthesis{flag: StyleFlag}
}

func FontSynthesisAll() FontSynthesis {
	return FontSynthesis{flag: AllFlags}
}

func NewFontSynthesis(flag SynthesisFlag) FontSynthesis {
	if flag.String() == "Invalid" {
		panic("invalid font flag")
	}
	return FontSynthesis{flag: flag}
}
