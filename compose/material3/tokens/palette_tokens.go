package tokens

import (
	"github.com/zodimo/go-compose/compose/ui/graphics"
)

var PaletteTokens = PaletteTokensData{
	Black:             graphics.NewColorSrgb(0, 0, 0, 255),
	Error0:            graphics.NewColorSrgb(0, 0, 0, 255),
	Error10:           graphics.NewColorSrgb(65, 14, 11, 255),
	Error100:          graphics.NewColorSrgb(255, 255, 255, 255),
	Error20:           graphics.NewColorSrgb(96, 20, 16, 255),
	Error30:           graphics.NewColorSrgb(140, 29, 24, 255),
	Error40:           graphics.NewColorSrgb(179, 38, 30, 255),
	Error50:           graphics.NewColorSrgb(220, 54, 46, 255),
	Error60:           graphics.NewColorSrgb(228, 105, 98, 255),
	Error70:           graphics.NewColorSrgb(236, 146, 142, 255),
	Error80:           graphics.NewColorSrgb(242, 184, 181, 255),
	Error90:           graphics.NewColorSrgb(249, 222, 220, 255),
	Error95:           graphics.NewColorSrgb(252, 238, 238, 255),
	Error99:           graphics.NewColorSrgb(255, 251, 249, 255),
	Neutral0:          graphics.NewColorSrgb(0, 0, 0, 255),
	Neutral10:         graphics.NewColorSrgb(29, 27, 32, 255),
	Neutral100:        graphics.NewColorSrgb(255, 255, 255, 255),
	Neutral12:         graphics.NewColorSrgb(33, 31, 38, 255),
	Neutral17:         graphics.NewColorSrgb(43, 41, 48, 255),
	Neutral20:         graphics.NewColorSrgb(50, 47, 53, 255),
	Neutral22:         graphics.NewColorSrgb(54, 52, 59, 255),
	Neutral24:         graphics.NewColorSrgb(59, 56, 62, 255),
	Neutral30:         graphics.NewColorSrgb(72, 70, 76, 255),
	Neutral4:          graphics.NewColorSrgb(15, 13, 19, 255),
	Neutral40:         graphics.NewColorSrgb(96, 93, 100, 255),
	Neutral50:         graphics.NewColorSrgb(121, 118, 125, 255),
	Neutral6:          graphics.NewColorSrgb(20, 18, 24, 255),
	Neutral60:         graphics.NewColorSrgb(147, 143, 150, 255),
	Neutral70:         graphics.NewColorSrgb(174, 169, 177, 255),
	Neutral80:         graphics.NewColorSrgb(202, 197, 205, 255),
	Neutral87:         graphics.NewColorSrgb(222, 216, 225, 255),
	Neutral90:         graphics.NewColorSrgb(230, 224, 233, 255),
	Neutral92:         graphics.NewColorSrgb(236, 230, 240, 255),
	Neutral94:         graphics.NewColorSrgb(243, 237, 247, 255),
	Neutral95:         graphics.NewColorSrgb(245, 239, 247, 255),
	Neutral96:         graphics.NewColorSrgb(247, 242, 250, 255),
	Neutral98:         graphics.NewColorSrgb(254, 247, 255, 255),
	Neutral99:         graphics.NewColorSrgb(255, 251, 255, 255),
	NeutralVariant0:   graphics.NewColorSrgb(0, 0, 0, 255),
	NeutralVariant10:  graphics.NewColorSrgb(29, 26, 34, 255),
	NeutralVariant100: graphics.NewColorSrgb(255, 255, 255, 255),
	NeutralVariant20:  graphics.NewColorSrgb(50, 47, 55, 255),
	NeutralVariant30:  graphics.NewColorSrgb(73, 69, 79, 255),
	NeutralVariant40:  graphics.NewColorSrgb(96, 93, 102, 255),
	NeutralVariant50:  graphics.NewColorSrgb(121, 116, 126, 255),
	NeutralVariant60:  graphics.NewColorSrgb(147, 143, 153, 255),
	NeutralVariant70:  graphics.NewColorSrgb(174, 169, 180, 255),
	NeutralVariant80:  graphics.NewColorSrgb(202, 196, 208, 255),
	NeutralVariant90:  graphics.NewColorSrgb(231, 224, 236, 255),
	NeutralVariant95:  graphics.NewColorSrgb(245, 238, 250, 255),
	NeutralVariant99:  graphics.NewColorSrgb(255, 251, 254, 255),
	Primary0:          graphics.NewColorSrgb(0, 0, 0, 255),
	Primary10:         graphics.NewColorSrgb(33, 0, 93, 255),
	Primary100:        graphics.NewColorSrgb(255, 255, 255, 255),
	Primary20:         graphics.NewColorSrgb(56, 30, 114, 255),
	Primary30:         graphics.NewColorSrgb(79, 55, 139, 255),
	Primary40:         graphics.NewColorSrgb(103, 80, 164, 255),
	Primary50:         graphics.NewColorSrgb(127, 103, 190, 255),
	Primary60:         graphics.NewColorSrgb(154, 130, 219, 255),
	Primary70:         graphics.NewColorSrgb(182, 157, 248, 255),
	Primary80:         graphics.NewColorSrgb(208, 188, 255, 255),
	Primary90:         graphics.NewColorSrgb(234, 221, 255, 255),
	Primary95:         graphics.NewColorSrgb(246, 237, 255, 255),
	Primary99:         graphics.NewColorSrgb(255, 251, 254, 255),
	Secondary0:        graphics.NewColorSrgb(0, 0, 0, 255),
	Secondary10:       graphics.NewColorSrgb(29, 25, 43, 255),
	Secondary100:      graphics.NewColorSrgb(255, 255, 255, 255),
	Secondary20:       graphics.NewColorSrgb(51, 45, 65, 255),
	Secondary30:       graphics.NewColorSrgb(74, 68, 88, 255),
	Secondary40:       graphics.NewColorSrgb(98, 91, 113, 255),
	Secondary50:       graphics.NewColorSrgb(122, 114, 137, 255),
	Secondary60:       graphics.NewColorSrgb(149, 141, 165, 255),
	Secondary70:       graphics.NewColorSrgb(176, 167, 192, 255),
	Secondary80:       graphics.NewColorSrgb(204, 194, 220, 255),
	Secondary90:       graphics.NewColorSrgb(232, 222, 248, 255),
	Secondary95:       graphics.NewColorSrgb(246, 237, 255, 255),
	Secondary99:       graphics.NewColorSrgb(255, 251, 254, 255),
	Tertiary0:         graphics.NewColorSrgb(0, 0, 0, 255),
	Tertiary10:        graphics.NewColorSrgb(49, 17, 29, 255),
	Tertiary100:       graphics.NewColorSrgb(255, 255, 255, 255),
	Tertiary20:        graphics.NewColorSrgb(73, 37, 50, 255),
	Tertiary30:        graphics.NewColorSrgb(99, 59, 72, 255),
	Tertiary40:        graphics.NewColorSrgb(125, 82, 96, 255),
	Tertiary50:        graphics.NewColorSrgb(152, 105, 119, 255),
	Tertiary60:        graphics.NewColorSrgb(181, 131, 146, 255),
	Tertiary70:        graphics.NewColorSrgb(210, 157, 172, 255),
	Tertiary80:        graphics.NewColorSrgb(239, 184, 200, 255),
	Tertiary90:        graphics.NewColorSrgb(255, 216, 228, 255),
	Tertiary95:        graphics.NewColorSrgb(255, 236, 241, 255),
	Tertiary99:        graphics.NewColorSrgb(255, 251, 250, 255),
	White:             graphics.NewColorSrgb(255, 255, 255, 255),
}

type PaletteTokensData struct {
	Black             graphics.Color
	Error0            graphics.Color
	Error10           graphics.Color
	Error100          graphics.Color
	Error20           graphics.Color
	Error30           graphics.Color
	Error40           graphics.Color
	Error50           graphics.Color
	Error60           graphics.Color
	Error70           graphics.Color
	Error80           graphics.Color
	Error90           graphics.Color
	Error95           graphics.Color
	Error99           graphics.Color
	Neutral0          graphics.Color
	Neutral10         graphics.Color
	Neutral100        graphics.Color
	Neutral12         graphics.Color
	Neutral17         graphics.Color
	Neutral20         graphics.Color
	Neutral22         graphics.Color
	Neutral24         graphics.Color
	Neutral30         graphics.Color
	Neutral4          graphics.Color
	Neutral40         graphics.Color
	Neutral50         graphics.Color
	Neutral6          graphics.Color
	Neutral60         graphics.Color
	Neutral70         graphics.Color
	Neutral80         graphics.Color
	Neutral87         graphics.Color
	Neutral90         graphics.Color
	Neutral92         graphics.Color
	Neutral94         graphics.Color
	Neutral95         graphics.Color
	Neutral96         graphics.Color
	Neutral98         graphics.Color
	Neutral99         graphics.Color
	NeutralVariant0   graphics.Color
	NeutralVariant10  graphics.Color
	NeutralVariant100 graphics.Color
	NeutralVariant20  graphics.Color
	NeutralVariant30  graphics.Color
	NeutralVariant40  graphics.Color
	NeutralVariant50  graphics.Color
	NeutralVariant60  graphics.Color
	NeutralVariant70  graphics.Color
	NeutralVariant80  graphics.Color
	NeutralVariant90  graphics.Color
	NeutralVariant95  graphics.Color
	NeutralVariant99  graphics.Color
	Primary0          graphics.Color
	Primary10         graphics.Color
	Primary100        graphics.Color
	Primary20         graphics.Color
	Primary30         graphics.Color
	Primary40         graphics.Color
	Primary50         graphics.Color
	Primary60         graphics.Color
	Primary70         graphics.Color
	Primary80         graphics.Color
	Primary90         graphics.Color
	Primary95         graphics.Color
	Primary99         graphics.Color
	Secondary0        graphics.Color
	Secondary10       graphics.Color
	Secondary100      graphics.Color
	Secondary20       graphics.Color
	Secondary30       graphics.Color
	Secondary40       graphics.Color
	Secondary50       graphics.Color
	Secondary60       graphics.Color
	Secondary70       graphics.Color
	Secondary80       graphics.Color
	Secondary90       graphics.Color
	Secondary95       graphics.Color
	Secondary99       graphics.Color
	Tertiary0         graphics.Color
	Tertiary10        graphics.Color
	Tertiary100       graphics.Color
	Tertiary20        graphics.Color
	Tertiary30        graphics.Color
	Tertiary40        graphics.Color
	Tertiary50        graphics.Color
	Tertiary60        graphics.Color
	Tertiary70        graphics.Color
	Tertiary80        graphics.Color
	Tertiary90        graphics.Color
	Tertiary95        graphics.Color
	Tertiary99        graphics.Color
	White             graphics.Color
}
