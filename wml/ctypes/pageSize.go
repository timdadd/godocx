package ctypes

import (
	"encoding/xml"
	"godocx/wml/stypes"
	"strconv"
)

// PageSize : w:pgSz
type PageSize struct {
	Width  *stypes.TwipsMeasure `xml:"w,attr,omitempty"`
	Height *stypes.TwipsMeasure `xml:"h,attr,omitempty"`
	Orient stypes.PageOrient    `xml:"orient,attr,omitempty"`
	Code   *int                 `xml:"code,attr,omitempty"`
}

func (p PageSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pgSz"

	if p.Width != nil {
		start.Attr = append(start.Attr, p.Width.XmlAttr("w:w"))
	}

	if p.Height != nil {
		start.Attr = append(start.Attr, p.Height.XmlAttr("w:h"))
	}

	if p.Orient != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:orient"}, Value: string(p.Orient)})
	}

	if p.Code != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:code"}, Value: strconv.Itoa(*p.Code)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
