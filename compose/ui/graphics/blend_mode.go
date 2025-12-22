package graphics

// BlendMode defines algorithms to use when painting on the canvas.
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-graphics/src/commonMain/kotlin/androidx/compose/ui/graphics/BlendMode.kt
type BlendMode int

const (
	BlendModeClear      BlendMode = 0
	BlendModeSrc        BlendMode = 1
	BlendModeDst        BlendMode = 2
	BlendModeSrcOver    BlendMode = 3
	BlendModeDstOver    BlendMode = 4
	BlendModeSrcIn      BlendMode = 5
	BlendModeDstIn      BlendMode = 6
	BlendModeSrcOut     BlendMode = 7
	BlendModeDstOut     BlendMode = 8
	BlendModeSrcAtop    BlendMode = 9
	BlendModeDstAtop    BlendMode = 10
	BlendModeXor        BlendMode = 11
	BlendModePlus       BlendMode = 12
	BlendModeModulate   BlendMode = 13
	BlendModeScreen     BlendMode = 14
	BlendModeOverlay    BlendMode = 15
	BlendModeDarken     BlendMode = 16
	BlendModeLighten    BlendMode = 17
	BlendModeColorDodge BlendMode = 18
	BlendModeColorBurn  BlendMode = 19
	BlendModeHardlight  BlendMode = 20
	BlendModeSoftlight  BlendMode = 21
	BlendModeDifference BlendMode = 22
	BlendModeExclusion  BlendMode = 23
	BlendModeMultiply   BlendMode = 24
	BlendModeHue        BlendMode = 25
	BlendModeSaturation BlendMode = 26
	BlendModeColor      BlendMode = 27
	BlendModeLuminosity BlendMode = 28
)

func (b BlendMode) String() string {
	switch b {
	case BlendModeClear:
		return "Clear"
	case BlendModeSrc:
		return "Src"
	case BlendModeDst:
		return "Dst"
	case BlendModeSrcOver:
		return "SrcOver"
	case BlendModeDstOver:
		return "DstOver"
	case BlendModeSrcIn:
		return "SrcIn"
	case BlendModeDstIn:
		return "DstIn"
	case BlendModeSrcOut:
		return "SrcOut"
	case BlendModeDstOut:
		return "DstOut"
	case BlendModeSrcAtop:
		return "SrcAtop"
	case BlendModeDstAtop:
		return "DstAtop"
	case BlendModeXor:
		return "Xor"
	case BlendModePlus:
		return "Plus"
	case BlendModeModulate:
		return "Modulate"
	case BlendModeScreen:
		return "Screen"
	case BlendModeOverlay:
		return "Overlay"
	case BlendModeDarken:
		return "Darken"
	case BlendModeLighten:
		return "Lighten"
	case BlendModeColorDodge:
		return "ColorDodge"
	case BlendModeColorBurn:
		return "ColorBurn"
	case BlendModeHardlight:
		return "HardLight"
	case BlendModeSoftlight:
		return "Softlight"
	case BlendModeDifference:
		return "Difference"
	case BlendModeExclusion:
		return "Exclusion"
	case BlendModeMultiply:
		return "Multiply"
	case BlendModeHue:
		return "Hue"
	case BlendModeSaturation:
		return "Saturation"
	case BlendModeColor:
		return "Color"
	case BlendModeLuminosity:
		return "Luminosity"
	default:
		return "Unknown"
	}

}

// IsSupported returns whether the BlendMode is supported.
// SrcOver is guaranteed to be supported.
func (b BlendMode) IsSupported() bool {
	return true
}
