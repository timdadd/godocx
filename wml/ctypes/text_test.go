package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"godocx/internal"

	"github.com/stretchr/testify/assert"
)

func TestTextMarshalXML(t *testing.T) {
	testCases := []struct {
		input    *Text
		expected string
	}{
		{NewText(), `<w:t></w:t>`},
		{TextFromString("Hello, World!"), `<w:t>Hello, World!</w:t>`},
	}

	for _, tc := range testCases {
		var result strings.Builder
		encoder := xml.NewEncoder(&result)

		start := xml.StartElement{Name: xml.Name{Local: "w:t"}}
		err := tc.input.MarshalXML(encoder, start)

		if err != nil {
			t.Errorf("Error during MarshalXML: %v", err)
		}

		err = encoder.Flush()
		if err != nil {
			t.Errorf("Error flushing encoder: %v", err)
		}

		if result.String() != tc.expected {
			t.Errorf("Expected XML:\n%s\n\nActual XML:\n%s", tc.expected, result.String())
		}
	}
}

func TestTextUnmarshalXML(t *testing.T) {
	testCases := []struct {
		input    string
		expected *Text
	}{
		{`<w:t></w:t>`, NewText()},
		{`<w:t xml:space="preserve">Hello, World!</w:t>`, &Text{
			Text:  "Hello, World!",
			Space: internal.ToPtr("preserve"),
		}},
		{`<w:t xml:space="preserve">Some text</w:t>`, &Text{
			Text:  "Some text",
			Space: internal.ToPtr("preserve"),
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			decoder := xml.NewDecoder(strings.NewReader(tc.input))
			var result Text

			startToken, err := decoder.Token()
			if err != nil {
				t.Fatalf("Error getting start token: %v", err)
			}

			err = result.UnmarshalXML(decoder, startToken.(xml.StartElement))
			assert.NoError(t, err, "Error during UnmarshalXML")
			assert.Equal(t, tc.expected.Text, result.Text, "Unexpected result")
			assert.Equal(t, tc.expected.Space, result.Space, "Unexpected space")
		})
	}
}
