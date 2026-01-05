package geometry

type RoundRect struct {
	Left   float32
	Top    float32
	Right  float32
	Bottom float32

	TopLeftCornerRadius     CornerRadius
	TopRightCornerRadius    CornerRadius
	BottomRightCornerRadius CornerRadius
	BottomLeftCornerRadius  CornerRadius
}

func NewRoundRect(rect Rect, topLeftCornerRadius, topRightCornerRadius, bottomRightCornerRadius, bottomLeftCornerRadius CornerRadius) RoundRect {
	return RoundRect{
		Left:                    rect.Left,
		Top:                     rect.Top,
		Right:                   rect.Right,
		Bottom:                  rect.Bottom,
		TopLeftCornerRadius:     topLeftCornerRadius,
		TopRightCornerRadius:    topRightCornerRadius,
		BottomRightCornerRadius: bottomRightCornerRadius,
		BottomLeftCornerRadius:  bottomLeftCornerRadius,
	}
}
