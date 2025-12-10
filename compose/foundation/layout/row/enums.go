package row

import "gioui.org/layout"

type Spacing = layout.Spacing
type Alignment = layout.Alignment

const (
	// SpaceEnd leaves space at the end.
	SpaceEnd Spacing = iota
	// SpaceStart leaves space at the start.
	SpaceStart
	// SpaceSides shares space between the start and end.
	SpaceSides
	// SpaceAround distributes space evenly between children,
	// with half as much space at the start and end.
	SpaceAround
	// SpaceBetween distributes space evenly between children,
	// leaving no space at the start and end.
	SpaceBetween
	// SpaceEvenly distributes space evenly between children and
	// at the start and end.
	SpaceEvenly
)

const (
	Start Alignment = iota
	End
	Middle
	Baseline
)
