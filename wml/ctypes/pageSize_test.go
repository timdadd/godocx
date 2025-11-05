package ctypes

import (
	"encoding/xml"
	"godocx/common/units"
	"strings"
	"testing"

	"godocx/wml/stypes"

	"github.com/stretchr/testify/assert"
)

func TestPageSize_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    PageSize
		expected string
	}{
		{
			name: "All attributes",
			input: PageSize{
				Width:  units.Inch(8.5).TwipsMeasure(),
				Height: units.Inch(11).TwipsMeasure(),
				Orient: stypes.PageOrientLandscape,
				Code:   intPtr(1),
			},
			expected: `<w:pgSz w:w="12240" w:h="15840" w:orient="landscape" w:code="1"></w:pgSz>`,
		},
		{
			name: "Some attributes",
			input: PageSize{
				Width:  units.Inch(8.5).TwipsMeasure(),
				Height: units.Inch(11).TwipsMeasure(),
			},
			expected: `<w:pgSz w:w="12240" w:h="15840"></w:pgSz>`,
		},
		{
			name:     "No attributes",
			input:    PageSize{},
			expected: `<w:pgSz></w:pgSz>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			//tt.input.XMLName.Local = tt.input.XMLName.Space + ":" + tt.input.XMLName.Local
			//tt.input.XMLName.Space = "w"
			encoder := xml.NewEncoder(&result)
			assert.NoError(t, encoder.Encode(&tt.input), "error encoding xml")
			assert.Equal(t, tt.expected, result.String(), "XML output not as expected")
		})
	}
}

func TestPageSize_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected PageSize
	}{
		{
			name:     "All attributes",
			inputXML: `<w:pgSz w:w="12240" w:h="15840" w:orient="landscape" w:code="1"></w:pgSz>`,
			expected: PageSize{
				Width:  units.Inch(8.5).TwipsMeasure(),
				Height: units.Inch(11).TwipsMeasure(),
				Orient: stypes.PageOrientLandscape,
				Code:   intPtr(1),
			},
		},
		{
			name:     "Some attributes",
			inputXML: `<w:pgSz w:w="12240" w:h="15840"></w:pgSz>`,
			expected: PageSize{
				Width:  units.Inch(8.5).TwipsMeasure(),
				Height: units.Inch(11).TwipsMeasure(),
			},
		},
		{
			name:     "No attributes",
			inputXML: `<w:pgSz></w:pgSz>`,
			expected: PageSize{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result PageSize

			assert.NoError(t, xml.Unmarshal([]byte(tt.inputXML), &result), "Error unmarshalling xml")
			assert.Equal(t, tt.expected, result, "XML output not as expected")
		})
	}
}
