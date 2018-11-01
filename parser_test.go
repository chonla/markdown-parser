package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEmptyContent(t *testing.T) {
	content := ""
	expected := Document{
		Elements: []Element{},
	}

	result := Parse(content)

	assert.Equal(t, expected, result)
}

func TestParseEmptyLineContent(t *testing.T) {
	content := "\n\n\n"
	expected := Document{
		Elements: []Element{},
	}

	result := Parse(content)

	assert.Equal(t, expected, result)
}

func TestParseSimpleDocument(t *testing.T) {
	content := "Test"
	expected := Document{
		Elements: []Element{
			Element{
				Text:     "Test",
				Elements: []Element{},
			},
		},
	}

	result := Parse(content)

	assert.Equal(t, expected, result)
}
