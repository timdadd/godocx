package stypes

import (
	"encoding/xml"
	"strconv"
)

// TwipsMeasure (derived from TWentieth of an Imperial Point) is a typographical measurement,
// defined as 1/20 of a typographical point. From a Microsoft and Postscript point of view
// one twip is 1/1440 inch or 17.639 µm.  Traditional printer's have a slightly different dimension
type TwipsMeasure int64

// SignedTwipsMeasure (derived from TWentieth of an Imperial Point) is a typographical measurement,
// defined as 1/20 of a typographical point. From a Microsoft and Postscript point of view
// one twip is 1/1440 inch or 17.639 µm.  Traditional printer's have a slightly different dimension
type SignedTwipsMeasure uint64

func (twip *TwipsMeasure) XmlAttr(local string) xml.Attr {
	return xml.Attr{
		Name:  xml.Name{Local: local},
		Value: strconv.Itoa(int(*twip))}
}

func (twip *TwipsMeasure) Nvl() int {
	if twip == nil {
		return 0
	}
	return int(*twip)
}

func (twip *SignedTwipsMeasure) XmlAttr(local string) xml.Attr {
	return xml.Attr{
		Name:  xml.Name{Local: local},
		Value: strconv.Itoa(int(*twip))}
}

func (twip *SignedTwipsMeasure) Nvl() int {
	if twip == nil {
		return 0
	}
	return int(*twip)
}

func twipFromStr(s string) (i int, err error) {
	if i, err = strconv.Atoi(s); err != nil {
		return 0, err
	}
	return
}
