package style

import (
	"fmt"

	"github.com/zodimo/go-compose/compose/ui/utils/lerp"
	"github.com/zodimo/go-compose/pkg/floatutils"
)

// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-text/src/commonMain/kotlin/androidx/compose/ui/text/style/TextGeometricTransform.kt;drc=4970f6e96cdb06089723da0ab8ec93ae3f067c7a;l=33

type GeometricTransform struct {
	ScaleX float32
	SkewX  float32
}

func (gt GeometricTransform) Equals(other GeometricTransform) bool {
	epsilon := floatutils.Float32EqualityThreshold
	return floatutils.Float32Equals(gt.ScaleX, other.ScaleX, epsilon) && floatutils.Float32Equals(gt.SkewX, other.SkewX, epsilon)
}

func (gt GeometricTransform) String() string {
	return fmt.Sprintf("TextGeometricTransform(scaleX=%f, skewX=%f)", gt.ScaleX, gt.SkewX)

}

func LerpGeometricTransform(
	start GeometricTransform,
	stop GeometricTransform,
	fraction float32,
) GeometricTransform {
	return GeometricTransform{
		ScaleX: lerp.Between32(start.ScaleX, stop.ScaleX, fraction),
		SkewX:  lerp.Between32(start.SkewX, stop.SkewX, fraction),
	}
}
