package docx

import (
	"github.com/timdadd/godocx/common/units"
	"github.com/timdadd/godocx/dml"
	"github.com/timdadd/godocx/internal"
)

type PicMeta struct {
	Para   *Paragraph
	Inline *dml.Inline
}

// AddPictureFromFile adds a new image to the document.
//
// Example usage:
//
//	// Add a picture to the document
//	_, err = document.AddPicture("gopher.png", units.Inch(2.9), units.Inch(2.9))
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Parameters:
//   - path: The path of the image file to be added.
//   - width: The width of the image in inches.
//   - height: The height of the image in inches.
//
// Returns:
//   - *PicMeta: Metadata about the added picture, including the Paragraph instance and Inline element.
//   - error: An error, if any occurred during the process.
func (rd *RootDoc) AddPictureFromFile(path string, width units.Inch, height units.Inch) (pm *PicMeta, err error) {
	var imgBytes []byte
	if imgBytes, err = internal.FileToByte(path); err != nil {
		return nil, err
	}
	return rd.AddImage(imgBytes, width, height)
}

// AddImage adds a new image to the document.
//
// Example usage:
//
//	// Add a []byte image to the document
//	_, err = document.AddImage([]byte(myImage), units.Inch(2.9), units.Inch(2.9))
//
// Parameters:
//   - image: The image file to be added.
//   - width: The width of the image in inches.
//   - height: The height of the image in inches.
//
// Returns:
//   - *PicMeta: Metadata about the added image, including the Paragraph instance and Inline element.
//   - error: An error, if any occurred during the process.
func (rd *RootDoc) AddImage(imgBytes []byte, width, height units.Inch) (pm *PicMeta, err error) {
	p := newParagraph(rd)
	bodyElem := DocumentChild{
		Para: p,
	}
	rd.Document.Body.Children = append(rd.Document.Body.Children, bodyElem)
	return p.AddImage(imgBytes, width, height)
}
