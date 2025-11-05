package ctypes

import (
	"encoding/xml"
	"godocx/common/units"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPageMargin_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    PageMargin
		expected string
	}{
		{
			name: "All attributes",
			input: PageMargin{
				Left:   units.Inch(1).TwipsMeasure(),
				Right:  units.Inch(1).TwipsMeasure(),
				Gutter: units.Twip(0).TwipsMeasure(),
				Header: units.Inch(.5).TwipsMeasure(),
				Top:    units.Inch(1).SignedTwipsMeasure(),
				Footer: units.Inch(.5).TwipsMeasure(),
				Bottom: units.Inch(1).SignedTwipsMeasure(),
			},
			expected: `<w:pgMar w:left="1440" w:right="1440" w:gutter="0" w:header="720" w:top="1440" w:footer="720" w:bottom="1440"></w:pgMar>`,
		},
		{
			name: "Some attributes",
			input: PageMargin{
				Left:   units.Inch(1).TwipsMeasure(),
				Right:  units.Inch(1).TwipsMeasure(),
				Top:    units.Inch(1).SignedTwipsMeasure(),
				Bottom: units.Inch(1).SignedTwipsMeasure(),
			},
			expected: `<w:pgMar w:left="1440" w:right="1440" w:top="1440" w:bottom="1440"></w:pgMar>`,
		},
		{
			name:     "No attributes",
			input:    PageMargin{},
			expected: `<w:pgMar></w:pgMar>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			assert.NoError(t, encoder.Encode(&tt.input), "error encoding xml")
			assert.Equal(t, tt.expected, result.String(), "XML output not as expected")
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			assert.NoError(t, encoder.Encode(tt.input), "Error marshalling xml")
			assert.Equal(t, tt.expected, result.String(), "XML output not as expected")
		})
	}
}

func TestPageMargin_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected PageMargin
	}{
		{
			name:     "All attributes",
			inputXML: `<w:pgMar w:left="1440" w:right="1440" w:gutter="0" w:header="720" w:top="1440" w:footer="720" w:bottom="1440"></w:pgMar>`,
			expected: PageMargin{
				Left:   units.Inch(1).TwipsMeasure(),
				Right:  units.Inch(1).TwipsMeasure(),
				Gutter: units.Twip(0).TwipsMeasure(),
				Header: units.Inch(.5).TwipsMeasure(),
				Top:    units.Inch(1).SignedTwipsMeasure(),
				Footer: units.Inch(.5).TwipsMeasure(),
				Bottom: units.Inch(1).SignedTwipsMeasure(),
			},
		},
		{
			name:     "Some attributes",
			inputXML: `<w:pgMar w:left="1440" w:right="1440" w:top="1440" w:bottom="1440"></w:pgMar>`,
			expected: PageMargin{
				Left:   units.Inch(1).TwipsMeasure(),
				Right:  units.Inch(1).TwipsMeasure(),
				Top:    units.Inch(1).SignedTwipsMeasure(),
				Bottom: units.Inch(1).SignedTwipsMeasure(),
			},
		},
		{
			name:     "No attributes",
			inputXML: `<w:pgMar></w:pgMar>`,
			expected: PageMargin{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result PageMargin

			assert.NoError(t, xml.Unmarshal([]byte(tt.inputXML), &result), "Error unmarshalling xml")
			assert.Equal(t, tt.expected, result, "XML output not as expected")
		})
	}
}

func intPtr(i int) *int {
	return &i
}
