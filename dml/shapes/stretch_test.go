package shapes

import (
	"encoding/xml"
	"testing"

	"godocx/dml/dmlct"

	"github.com/stretchr/testify/assert"
)

func TestMarshalStretch(t *testing.T) {
	tests := []struct {
		name        string
		stretch     *Stretch
		expectedXML string
	}{
		{
			name:        "With FillRect",
			stretch:     &Stretch{FillRect: &dmlct.RelativeRect{}},
			expectedXML: `<a:stretch><a:fillRect></a:fillRect></a:stretch>`,
		},
		{
			name:        "Without FillRect",
			stretch:     &Stretch{},
			expectedXML: `<a:stretch></a:stretch>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generatedXML, err := xml.Marshal(tt.stretch)
			assert.NoError(t, err, "Error during MarshalXML")
			assert.Equal(t, tt.expectedXML, string(generatedXML), "Unexpected result")
		})
	}
}

func TestUnmarshalStretch(t *testing.T) {
	tests := []struct {
		name           string
		inputXML       string
		expectedResult Stretch
	}{
		{
			name:     "With FillRect",
			inputXML: `<a:stretch><a:fillRect></a:fillRect></a:stretch>`,
			expectedResult: Stretch{
				FillRect: &dmlct.RelativeRect{},
			},
		},
		{
			name:           "Without FillRect",
			inputXML:       `<a:stretch></a:stretch>`,
			expectedResult: Stretch{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Stretch

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			assert.NoError(t, err, "Error during UnmarshalXML")

			if (result.FillRect == nil) != (tt.expectedResult.FillRect == nil) {
				t.Errorf("Expected FillRect to be %v, but got %v", tt.expectedResult.FillRect, result.FillRect)
			}
		})
	}
}
