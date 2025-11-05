package units

import "godocx/wml/stypes"

type Units interface {
	ToEmu() Emu
	ToTwip() Twip
	ToEmuUInt64Ptr() *uint64
	ToEmuIntPtr() *int
	TwipsMeasure() *stypes.TwipsMeasure
	SignedTwipsMeasure() *stypes.SignedTwipsMeasure
}

// Emu represents a dimension known as English Metric Units (EMUs).
// An EMU is a unit of measurement used by Microsoft for internal calculations in Word
// and other Office programs. It was created so that both English (inches) and metric (centimeters)
// measurements could be represented as whole numbers, which is more efficient for computer processing.
//
// 1 EMU = */914,400 of an inch & 1/360,000 of a centimeter & 1/12,700 of a point.
type Emu int64

func (emu Emu) ToEmu() Emu { return emu }
func (emu Emu) ToEmuUInt64Ptr() *uint64 {
	ui64 := uint64(emu)
	return &ui64
}
func (emu Emu) ToEmuIntPtr() *int {
	i := int(emu)
	return &i
}

func (emu Emu) Inches() Inch {
	return Inch(emu / 914400.0)
}
func (emu Emu) CMs() CM {
	return CM(emu / 360000.0)
}
func (emu Emu) MMs() MM {
	return MM(emu / 36000.0)
}
func (emu Emu) Points() Point {
	return Point(emu / 12700.0)
}

func (emu Emu) Twips() Twip {
	return Twip(emu / 635)
}
func (emu Emu) ToTwip() Twip { return Twip(emu / 635) }
func (emu Emu) TwipsMeasure() *stypes.TwipsMeasure {
	return emu.ToTwip().TwipsMeasure()
}
func (emu Emu) SignedTwipsMeasure() *stypes.SignedTwipsMeasure {
	return emu.ToTwip().SignedTwipsMeasure()
}

type Twip uint64

func (twip Twip) ToEmu() Emu              { return Emu(twip * 635) }
func (twip Twip) ToEmuUInt64Ptr() *uint64 { return twip.ToEmuUInt64Ptr() }
func (twip Twip) ToEmuIntPtr() *int       { return twip.ToEmuIntPtr() }
func (twip Twip) ToTwip() Twip            { return twip }
func (twip Twip) Inches() Inch            { return Inch(twip / 1440.0) }
func (twip Twip) Points() Point           { return Point(twip / 20) }
func (twip Twip) TwipsMeasure() *stypes.TwipsMeasure {
	sTM := stypes.TwipsMeasure(twip)
	return &sTM
}
func (twip Twip) SignedTwipsMeasure() *stypes.SignedTwipsMeasure {
	sSTM := stypes.SignedTwipsMeasure(twip)
	return &sSTM
}

// Inch represents a unit measure in inches.
type Inch float64

// ToEmu converts inches to EMUs.
func (i Inch) ToEmu() Emu {
	return Emu(i * 914400)
}
func (i Inch) ToEmuUInt64Ptr() *uint64 { return i.ToEmu().ToEmuUInt64Ptr() }
func (i Inch) ToEmuIntPtr() *int       { return i.ToEmu().ToEmuIntPtr() }
func (i Inch) ToTwip() Twip            { return Twip(i * 1440) }
func (i Inch) TwipsMeasure() *stypes.TwipsMeasure {
	return i.ToTwip().TwipsMeasure()
}
func (i Inch) SignedTwipsMeasure() *stypes.SignedTwipsMeasure {
	return i.ToTwip().SignedTwipsMeasure()
}

// CM represents a unit measure in centimeters.
type CM float64

func (cm CM) ToEmu() Emu {
	return Emu(cm * 360000)
}
func (cm CM) ToEmuUInt64Ptr() *uint64 { return cm.ToEmu().ToEmuUInt64Ptr() }
func (cm CM) ToEmuIntPtr() *int       { return cm.ToEmu().ToEmuIntPtr() }
func (cm CM) ToTwip() Twip            { return Twip(cm*360000/635 + .5) }
func (cm CM) TwipsMeasure() *stypes.TwipsMeasure {
	return cm.ToTwip().TwipsMeasure()
}
func (cm CM) SignedTwipsMeasure() *stypes.SignedTwipsMeasure {
	return cm.ToTwip().SignedTwipsMeasure()
}

// MM represents a unit measure in miliimeters.
type MM float64

func (mm MM) ToEmu() Emu {
	return Emu(mm * 36000)
}
func (mm MM) ToEmuUInt64Ptr() *uint64 { return mm.ToEmu().ToEmuUInt64Ptr() }
func (mm MM) ToEmuIntPtr() *int {
	return mm.ToEmu().ToEmuIntPtr()
}
func (mm MM) ToTwip() Twip { return Twip(mm*36000/635 + .5) }
func (mm MM) TwipsMeasure() *stypes.TwipsMeasure {
	return mm.ToTwip().TwipsMeasure()
}
func (mm MM) SignedTwipsMeasure() *stypes.SignedTwipsMeasure {
	return mm.ToTwip().SignedTwipsMeasure()
}

// Point represents a unit measure in points
type Point int64

func (p Point) ToEmu() Emu              { return Emu(p * 12700) }
func (p Point) ToEmuUInt64Ptr() *uint64 { return p.ToEmu().ToEmuUInt64Ptr() }
func (p Point) ToEmuIntPtr() *int {
	return p.ToEmu().ToEmuIntPtr()
}
func (p Point) ToTwip() Twip { return Twip(p * 20) }
func (p Point) TwipsMeasure() *stypes.TwipsMeasure {
	return p.ToTwip().TwipsMeasure()
}
func (p Point) SignedTwipsMeasure() *stypes.SignedTwipsMeasure {
	return p.SignedTwipsMeasure()
}

type Pixel int

// ToEmu assumes 96 pixels per inch
func (p Pixel) ToEmu() Emu              { return p.ToEmuUsingPPI(96) }
func (p Pixel) ToEmuUInt64Ptr() *uint64 { return p.ToEmu().ToEmuUInt64Ptr() }
func (p Pixel) ToEmuIntPtr() *int {
	return p.ToEmu().ToEmuIntPtr()
}
func (p Pixel) ToTwip() Twip { return p.ToTwipUsingPPI(96) }
func (p Pixel) TwipsMeasure() *stypes.TwipsMeasure {
	return p.ToTwip().TwipsMeasure()
}
func (p Pixel) SignedTwipsMeasure(pixel Pixel) *stypes.SignedTwipsMeasure {
	return p.ToTwip().SignedTwipsMeasure()
}

// ToEmuUsingPPI uses the Pixels Per Inch provided in the call
func (p Pixel) ToEmuUsingPPI(PPI int) Emu              { return Inch.ToEmu(Inch(float64(p) / float64(PPI))) }
func (p Pixel) ToEmuUInt64PtrUsingPPI(PPI int) *uint64 { return p.ToEmuUsingPPI(PPI).ToEmuUInt64Ptr() }
func (p Pixel) ToEmuUIntPtrUsingPPI(PPI int) *int      { return p.ToEmuUsingPPI(PPI).ToEmuIntPtr() }
func (p Pixel) ToTwipUsingPPI(PPI int) Twip            { return Inch.ToTwip(Inch(float64(p) / float64(PPI))) }
