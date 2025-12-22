package font

// FontWeight represents the thickness of the typeface.
type FontWeight int

const (
	FontWeightThin       FontWeight = 100
	FontWeightExtraLight FontWeight = 200
	FontWeightLight      FontWeight = 300
	FontWeightNormal     FontWeight = 400
	FontWeightMedium     FontWeight = 500
	FontWeightSemiBold   FontWeight = 600
	FontWeightBold       FontWeight = 700
	FontWeightExtraBold  FontWeight = 800
	FontWeightBlack      FontWeight = 900
	FontWeightW100       FontWeight = 100
	FontWeightW200       FontWeight = 200
	FontWeightW300       FontWeight = 300
	FontWeightW400       FontWeight = 400
	FontWeightW500       FontWeight = 500
	FontWeightW600       FontWeight = 600
	FontWeightW700       FontWeight = 700
	FontWeightW800       FontWeight = 800
	FontWeightW900       FontWeight = 900
)

func (w FontWeight) Compare(other FontWeight) int {
	if w < other {
		return -1
	} else if w > other {
		return 1
	}
	return 0
}

// LerpFontWeight interpolates between two FontWeights.
func LerpFontWeight(start, stop FontWeight, fraction float32) FontWeight {
	return FontWeight(float32(start) + (float32(stop)-float32(start))*fraction)
}

// FontStyle represents the style of the typeface.
type FontStyle int

const (
	FontStyleNormal FontStyle = 0
	FontStyleItalic FontStyle = 1
)

func (s FontStyle) String() string {
	switch s {
	case FontStyleNormal:
		return "Normal"
	case FontStyleItalic:
		return "Italic"
	default:
		return "Invalid"
	}
}

// FontFamily represents a family of fonts.
type FontFamily interface {
}

type GenericFontFamily struct {
	Name string
}

var (
	FontFamilyDefault   FontFamily = GenericFontFamily{Name: "Default"}
	FontFamilySansSerif FontFamily = GenericFontFamily{Name: "SansSerif"}
	FontFamilySerif     FontFamily = GenericFontFamily{Name: "Serif"}
	FontFamilyMonospace FontFamily = GenericFontFamily{Name: "Monospace"}
	FontFamilyCursive   FontFamily = GenericFontFamily{Name: "Cursive"}
)
