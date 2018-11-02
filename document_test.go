package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewDocument(t *testing.T) {
	expected := &Document{
		Element: &Element{
			Type:     "doc",
			Elements: []*Element{},
		},
	}

	doc := NewDocument()

	assert.Equal(t, expected, doc)
}

func TestAppendElement(t *testing.T) {
	doc := NewDocument()
	expected := &Document{
		Element: &Element{
			Type: "doc",
			Elements: []*Element{
				&Element{
					Text:     "Test",
					Type:     "text",
					Elements: []*Element{},
				},
			},
		},
	}

	doc.Append(&Element{
		Text:     "Test",
		Type:     "text",
		Elements: []*Element{},
	})

	assert.Equal(t, expected, doc)
}
