package ctypes

import (
	"encoding/xml"
	"godocx/wml/stypes"
)

// PageMargin represents the page margins of a Word document.
type PageMargin struct {
	Left   *stypes.TwipsMeasure       `xml:"left,attr,omitempty"`
	Right  *stypes.TwipsMeasure       `xml:"right,attr,omitempty"`
	Gutter *stypes.TwipsMeasure       `xml:"gutter,attr,omitempty"`
	Header *stypes.TwipsMeasure       `xml:"header,attr,omitempty"`
	Top    *stypes.SignedTwipsMeasure `xml:"top,attr,omitempty"`
	Footer *stypes.TwipsMeasure       `xml:"footer,attr,omitempty"`
	Bottom *stypes.SignedTwipsMeasure `xml:"bottom,attr,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for the PageMargin type.
// It encodes the PageMargin to its corresponding XML representation.
func (p PageMargin) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pgMar"

	attrs := []struct {
		local string
		value interface{}
	}{
		{"w:left", p.Left},
		{"w:right", p.Right},
		{"w:gutter", p.Gutter},
		{"w:header", p.Header},
		{"w:top", p.Top},
		{"w:footer", p.Footer},
		{"w:bottom", p.Bottom},
	}

	for _, attr := range attrs {
		switch v := attr.value.(type) {
		case *stypes.TwipsMeasure:
			if v != nil {
				start.Attr = append(start.Attr, v.XmlAttr(attr.local))
			}
		case *stypes.SignedTwipsMeasure:
			if v != nil {
				start.Attr = append(start.Attr, v.XmlAttr(attr.local))
			}
		}
	}

	return e.EncodeElement("", start)
}
