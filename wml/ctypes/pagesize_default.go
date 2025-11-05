package ctypes

import (
	"godocx/common/units"
	"godocx/wml/stypes"
)

/*
code  |  name    |  size
---------------------------
1     |  A4      |  210 × 297 mm
2     |  Letter  |  8.5 × 11 inches
3     |  Legal   |  8.5 × 14 inches
4     |  A3      |  297 × 420 mm
5     |  B4      |  250 × 353 mm
6     |  B5      |  176 × 250 mm
9     |  A5      |  148 × 210 mm
11    |  A6      |  105 × 148 mm
*/

var (
	A1Code = 4
	A1     = &PageSize{
		Width:  units.MM(594).TwipsMeasure(),
		Height: units.MM(841).TwipsMeasure(),
		Orient: stypes.PageOrientPortrait,
		Code:   &A1Code,
	}

	A2Code = 4
	A2     = &PageSize{
		Width:  units.MM(420).TwipsMeasure(),
		Height: units.MM(594).TwipsMeasure(),
		Orient: stypes.PageOrientPortrait,
		Code:   &A2Code,
	}

	A3Code = 4
	A3     = &PageSize{
		Width:  units.MM(297).TwipsMeasure(),
		Height: units.MM(420).TwipsMeasure(),
		Orient: stypes.PageOrientPortrait,
		Code:   &A3Code,
	}

	A4Code = 1
	A4     = &PageSize{
		Width:  units.MM(210).TwipsMeasure(),
		Height: units.MM(297).TwipsMeasure(),
		Orient: stypes.PageOrientPortrait,
		Code:   &A4Code,
	}

	A5Code = 1
	A5     = &PageSize{
		Width:  units.MM(148).TwipsMeasure(),
		Height: units.MM(210).TwipsMeasure(),
		Orient: stypes.PageOrientPortrait,
		Code:   &A5Code,
	}
)
