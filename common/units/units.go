package units

type Units interface {
	ToEmu() Emu
}

// Emu represents a dimension known as English Metric Units (EMUs).
// An EMU is a unit of measurement used by Microsoft for internal calculations in Word
// and other Office programs. It was created so that both English (inches) and metric (centimeters)
// measurements could be represented as whole numbers, which is more efficient for computer processing.
//
// 1 EMU = */914,400 of an inch & 1/360,000 of a centimeter & 1/12,700 of a point.
type Emu int64

func (emu Emu) ToEmu() Emu { return emu }

// Inch represents a unit measure in inches.
type Inch float64

// ToEmu converts inches to EMUs.
func (i Inch) ToEmu() Emu {
	return Emu(i * 914400)
}

// CM represents a unit measure in centimeters.
type CM float64

func (cm CM) ToEmu() Emu {
	return Emu(cm * 360000)
}

// MM represents a unit measure in miliimeters.
type MM float64

func (mm MM) ToEmu() Emu {
	return Emu(mm * 36000)
}

// Point represents a unit measure in points
type Point int64

func (p Point) ToEmu() Emu { return Emu(p * 12700) }

type Pixel int

// ToEmu assumes 96 pixels per inch
func (p Pixel) ToEmu() Emu { return Inch.ToEmu(Inch(p / 96)) }

// ToEmuUsingPPI uses the Pixels Per Inch factor
func (p Pixel) ToEmuUsingPPI(PPI int) Emu { return Inch.ToEmu(Inch(int(p) / PPI)) }
