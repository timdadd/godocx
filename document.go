package godocx

import (
	_ "embed"
	"os"
	"path/filepath"

	"godocx/docx"
	"godocx/packager"
)

//go:embed "templates/meaty_doc.docx"
var defaultDocx []byte

// NewDocument creates a new document from the default template.
func NewDocument() (*docx.RootDoc, error) {
	return InitialDocument(defaultDocx)
}

// OpenDocument opens a document from the given file name.
func OpenDocument(fileName string) (rd *docx.RootDoc, err error) {
	var docxContent []byte
	if docxContent, err = os.ReadFile(filepath.Clean(fileName)); err != nil {
		return
	}
	return InitialDocument(docxContent)
}

// InitialDocument starts a document using the document template given
func InitialDocument(docTemplate []byte) (*docx.RootDoc, error) {
	return packager.Unpack(docTemplate)
}
