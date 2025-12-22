package text

import (
	"github.com/zodimo/go-compose/compose/ui/text/font"
	"github.com/zodimo/go-compose/compose/ui/text/style"
	"github.com/zodimo/go-maybe"
)

type TextStyle struct {
	Color      maybe.Maybe[Color]      // Color.Unspecified,
	FontSize   maybe.Maybe[uiTextUnit] //TextUnit.Unspecified,
	FontWeight maybe.Maybe[font.FontWeight]
	FontStyle  maybe.Maybe[font.FontStyle]
	// FontSynthesis          any
	FontFamily maybe.Maybe[font.FontFamily]
	// FontFeatureSettings    any
	LetterSpacing          maybe.Maybe[uiTextUnit] //TextUnit.Unspecified,
	BaselineShift          maybe.Maybe[style.BaselineShift]
	TextGeometricTransform maybe.Maybe[uiTextGeometricTransform]
	LocaleList             maybe.Maybe[uiLocaleList]
	Background             maybe.Maybe[Color]
	TextDecoration         maybe.Maybe[uiTextDecoration]
	Shadow                 maybe.Maybe[uiShadow]
	TextAlign              maybe.Maybe[style.TextAlign]
	TextDirection          maybe.Maybe[uiDirection]
	LineHeight             maybe.Maybe[uiTextUnit] //TextUnit.Unspecified,
	TextIndent             maybe.Maybe[uiIndent]

	// color: Color = Color.Unspecified,
	// fontSize: TextUnit = TextUnit.Unspecified,
	// fontWeight: FontWeight? = null,
	// fontStyle: FontStyle? = null,
	// fontSynthesis: FontSynthesis? = null,
	// fontFamily: FontFamily? = null,
	// fontFeatureSettings: String? = null,
	// letterSpacing: TextUnit = TextUnit.Unspecified,
	// baselineShift: BaselineShift? = null,
	// textGeometricTransform: TextGeometricTransform? = null,
	// localeList: LocaleList? = null,
	// background: Color = Color.Unspecified,
	// textDecoration: TextDecoration? = null,
	// shadow: Shadow? = null,
	// textAlign: TextAlign? = null,
	// textDirection: TextDirection? = null,
	// lineHeight: TextUnit = TextUnit.Unspecified,
	// textIndent: TextIndent? = null,

}

func TextStyleDefaults() TextStyle {

	//   color: Color = Color.Unspecified,
	//     fontSize: TextUnit = TextUnit.Unspecified,
	//     fontWeight: FontWeight? = null,
	//     fontStyle: FontStyle? = null,
	//     fontSynthesis: FontSynthesis? = null,
	//     fontFamily: FontFamily? = null,
	//     fontFeatureSettings: String? = null,
	//     letterSpacing: TextUnit = TextUnit.Unspecified,
	//     baselineShift: BaselineShift? = null,
	//     textGeometricTransform: TextGeometricTransform? = null,
	//     localeList: LocaleList? = null,
	//     background: Color = Color.Unspecified,
	//     textDecoration: TextDecoration? = null,
	//     shadow: Shadow? = null,
	//     textAlign: TextAlign? = null,
	//     textDirection: TextDirection? = null,
	//     lineHeight: TextUnit = TextUnit.Unspecified,
	//     textIndent: TextIndent? = null,

	return TextStyle{}
}
