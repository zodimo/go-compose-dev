package tokens

type ShapeTokenKey int

const (
	ShapeTokenKeyUnspecified ShapeTokenKey = iota
	ShapeTokenKeyCornerExtraExtraLarge
	ShapeTokenKeyCornerExtraLarge
	ShapeTokenKeyCornerExtraLargeIncreased
	ShapeTokenKeyCornerExtraLargeTop
	ShapeTokenKeyCornerExtraSmall
	ShapeTokenKeyCornerExtraSmallTop
	ShapeTokenKeyCornerFull
	ShapeTokenKeyCornerLarge
	ShapeTokenKeyCornerLargeEnd
	ShapeTokenKeyCornerLargeIncreased
	ShapeTokenKeyCornerLargeStart
	ShapeTokenKeyCornerLargeTop
	ShapeTokenKeyCornerMedium
	ShapeTokenKeyCornerNone
	ShapeTokenKeyCornerSmall
)

func (s ShapeTokenKey) String() string {
	switch s {
	case ShapeTokenKeyCornerExtraExtraLarge:
		return "CornerExtraExtraLarge"
	case ShapeTokenKeyCornerExtraLarge:
		return "CornerExtraLarge"
	case ShapeTokenKeyCornerExtraLargeIncreased:
		return "CornerExtraLargeIncreased"
	case ShapeTokenKeyCornerExtraLargeTop:
		return "CornerExtraLargeTop"
	case ShapeTokenKeyCornerExtraSmall:
		return "CornerExtraSmall"
	case ShapeTokenKeyCornerExtraSmallTop:
		return "CornerExtraSmallTop"
	case ShapeTokenKeyCornerFull:
		return "CornerFull"
	case ShapeTokenKeyCornerLarge:
		return "CornerLarge"
	case ShapeTokenKeyCornerLargeEnd:
		return "CornerLargeEnd"
	case ShapeTokenKeyCornerLargeIncreased:
		return "CornerLargeIncreased"
	case ShapeTokenKeyCornerLargeStart:
		return "CornerLargeStart"
	case ShapeTokenKeyCornerLargeTop:
		return "CornerLargeTop"
	case ShapeTokenKeyCornerMedium:
		return "CornerMedium"
	case ShapeTokenKeyCornerNone:
		return "CornerNone"
	case ShapeTokenKeyCornerSmall:
		return "CornerSmall"
	default:
		return "Unspecified"
	}
}

var ShapeKeyTokens = ShapeKeyTokensData{
	Unspecified:               ShapeTokenKeyUnspecified,
	CornerExtraExtraLarge:     ShapeTokenKeyCornerExtraExtraLarge,
	CornerExtraLarge:          ShapeTokenKeyCornerExtraLarge,
	CornerExtraLargeIncreased: ShapeTokenKeyCornerExtraLargeIncreased,
	CornerExtraLargeTop:       ShapeTokenKeyCornerExtraLargeTop,
	CornerExtraSmall:          ShapeTokenKeyCornerExtraSmall,
	CornerExtraSmallTop:       ShapeTokenKeyCornerExtraSmallTop,
	CornerFull:                ShapeTokenKeyCornerFull,
	CornerLarge:               ShapeTokenKeyCornerLarge,
	CornerLargeEnd:            ShapeTokenKeyCornerLargeEnd,
	CornerLargeIncreased:      ShapeTokenKeyCornerLargeIncreased,
	CornerLargeStart:          ShapeTokenKeyCornerLargeStart,
	CornerLargeTop:            ShapeTokenKeyCornerLargeTop,
	CornerMedium:              ShapeTokenKeyCornerMedium,
	CornerNone:                ShapeTokenKeyCornerNone,
	CornerSmall:               ShapeTokenKeyCornerSmall,
}

type ShapeKeyTokensData struct {
	Unspecified               ShapeTokenKey
	CornerExtraExtraLarge     ShapeTokenKey
	CornerExtraLarge          ShapeTokenKey
	CornerExtraLargeIncreased ShapeTokenKey
	CornerExtraLargeTop       ShapeTokenKey
	CornerExtraSmall          ShapeTokenKey
	CornerExtraSmallTop       ShapeTokenKey
	CornerFull                ShapeTokenKey
	CornerLarge               ShapeTokenKey
	CornerLargeEnd            ShapeTokenKey
	CornerLargeIncreased      ShapeTokenKey
	CornerLargeStart          ShapeTokenKey
	CornerLargeTop            ShapeTokenKey
	CornerMedium              ShapeTokenKey
	CornerNone                ShapeTokenKey
	CornerSmall               ShapeTokenKey
}
