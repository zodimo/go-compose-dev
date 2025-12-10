package size

const NotSet = -1

type SizeData struct {
	Width    int
	Height   int
	Required bool
}

func DefaultSize() SizeData {
	return SizeData{
		Width:    NotSet,
		Height:   NotSet,
		Required: false,
	}
}
