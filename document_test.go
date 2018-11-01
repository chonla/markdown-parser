package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewDocument(t *testing.T) {
	doc := NewDocument()
	expected := Document{
		Elements: []Element{},
	}

	assert.Equal(t, expected, doc)
}

func TestAppendElement(t *testing.T) {
	doc := NewDocument()
	expected := Document{
		Elements: []Element{
			Element{
				Text:     "Test",
				Elements: []Element{},
			},
		},
	}

	doc.Append(Element{
		Text:     "Test",
		Elements: []Element{},
	})

	assert.Equal(t, expected, doc)
}
