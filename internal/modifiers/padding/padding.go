package padding

import (
	"gioui.org/io/system"
)

const NotSet = -1

const RTL = system.RTL

type PaddingData struct {
	Start    int
	Top      int
	End      int
	Bottom   int
	RtlAware bool // future proofing for RTL support
}

func DefaultPadding() PaddingData {
	return PaddingData{
		Start:    NotSet,
		Top:      NotSet,
		End:      NotSet,
		Bottom:   NotSet,
		RtlAware: false,
	}
}
