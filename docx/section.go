package docx

import (
	"godocx/common/units"
	"godocx/wml/ctypes"
	"godocx/wml/stypes"
)

// SectionBreak starts a new section.
//
// # The current section is stored as a properties of a paragraph
//
// The details of the previous section are copied into the new section
func (rd *RootDoc) SectionBreak() {
	p := rd.AddEmptyParagraph()
	p.ensureProp()
	p.ct.Property.SectPr = rd.Document.Body.SectPr
	rd.Document.Body.SectPr = ctypes.NewSectionProper()
}

func (rd *RootDoc) SectionProp() *ctypes.SectionProp {
	return rd.Document.Body.SectPr
}

// PageSize returns the Complex Type PageSize for the current Section
func (rd *RootDoc) PageSize() (ps *ctypes.PageSize) {
	if ps = rd.SectionProp().PageSize; ps != nil {
		return
	}
	ps = new(ctypes.PageSize)
	rd.SectionProp().PageSize = ps
	return
}

func (rd *RootDoc) SetPageHeight(units units.Units) {
	rd.PageSize().Height = units.TwipsMeasure()
}

func (rd *RootDoc) PageHeight() units.Twip {
	return units.Twip(int(*rd.PageSize().Height))
}

func (rd *RootDoc) SetPageWidth(units units.Units) {
	rd.PageSize().Width = units.TwipsMeasure()
}

func (rd *RootDoc) PageWidth() units.Twip {
	return units.Twip(int(*rd.PageSize().Width))
}

func (rd *RootDoc) SetPageOrientation(po stypes.PageOrient) {
	rd.PageSize().Orient = po
}

func (rd *RootDoc) PageOrientation() stypes.PageOrient {
	return rd.PageSize().Orient
}

// PageMargin returns the Complex Type PageMargin for the current Document Section
func (rd *RootDoc) PageMargin() (pm *ctypes.PageMargin) {
	if pm = rd.SectionProp().PageMargin; pm != nil {
		return
	}
	pm = new(ctypes.PageMargin)
	rd.SectionProp().PageMargin = pm
	return
}

// SetLeftMargin sets the left margin of the PageMargin for the current Document Section
func (rd *RootDoc) SetLeftMargin(units units.Units) {
	rd.PageMargin().Left = units.TwipsMeasure()
}

// LeftMargin returns the left margin of the PageMargin for the current Document Section
func (rd *RootDoc) LeftMargin() units.Twip {
	return units.Twip(*rd.PageMargin().Left)
}

// SetRightMargin sets the right margin of the PageMargin for the current Document Section
func (rd *RootDoc) SetRightMargin(units units.Units) {
	rd.PageMargin().Right = units.TwipsMeasure()
}

// RightMargin returns the right margin of the PageMargin for the current Document Section
func (rd *RootDoc) RightMargin() units.Twip {
	return units.Twip(*rd.PageMargin().Right)
}

// SetTopMargin sets the top margin of the PageMargin for the current Document Section
func (rd *RootDoc) SetTopMargin(units units.Units) {
	rd.PageMargin().Top = units.SignedTwipsMeasure()
}

// TopMargin returns the top margin of the PageMargin for the current Document Section
func (rd *RootDoc) TopMargin() units.Twip {
	return units.Twip(*rd.PageMargin().Top)
}

// SetBottomMargin sets the bottom margin of the PageMargin for the current Document Section
func (rd *RootDoc) SetBottomMargin(units units.Units) {
	rd.PageMargin().Bottom = units.SignedTwipsMeasure()
}

// BottomMargin returns the bottom margin of the PageMargin for the current Document Section
func (rd *RootDoc) BottomMargin() units.Twip {
	return units.Twip(*rd.PageMargin().Bottom)
}

// SetGutterMargin sets the gutter margin of the PageMargin for the current Document Section.
// The gutter is an offset to the edge of the page (usually left/ for portrait and top/ for landscape)
// the gutter ensures that text is not lost in the binding or stapled area when a document is printed and bound.
func (rd *RootDoc) SetGutterMargin(units units.Units) {
	rd.PageMargin().Gutter = units.TwipsMeasure()
}

// GutterMargin returns the gutter margin of the PageMargin for the current Document Section
func (rd *RootDoc) GutterMargin() units.Twip {
	return units.Twip(*rd.PageMargin().Gutter)
}

// SetHeaderMargin sets the header margin of the PageMargin for the current Document Section
func (rd *RootDoc) SetHeaderMargin(units units.Units) {
	rd.PageMargin().Header = units.TwipsMeasure()
}

// HeaderMargin returns the header margin of the PageMargin for the current Document Section
func (rd *RootDoc) HeaderMargin() units.Twip {
	return units.Twip(*rd.PageMargin().Header)
}

// SetFooterMargin sets the footer margin of the PageMargin for the current Document Section
func (rd *RootDoc) SetFooterMargin(units units.Units) {
	rd.PageMargin().Footer = units.TwipsMeasure()
}

// FooterMargin returns the footer margin of the PageMargin for the current Document Section
func (rd *RootDoc) FooterMargin() units.Twip {
	return units.Twip(*rd.PageMargin().Footer)
}

// UsableTextArea determines the width and height after all margins and gutters removed
func (rd *RootDoc) UsableTextArea() (w, h units.Twip) {
	section := rd.SectionProp()
	if section == nil || section.PageSize == nil {
		return 0, 0
	}
	var width, height int
	width = section.PageSize.Width.Nvl()
	height = section.PageSize.Height.Nvl()
	if section.PageMargin != nil {
		width = width - section.PageMargin.Left.Nvl() - section.PageMargin.Right.Nvl()
		height = height - section.PageMargin.Top.Nvl() - section.PageMargin.Bottom.Nvl()
	}
	return units.Twip(width), units.Twip(height)
}
