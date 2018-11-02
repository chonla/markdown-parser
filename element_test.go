package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewElement(t *testing.T) {
	expected := &Element{
		Text:     "good",
		Type:     "text",
		Parent:   nil,
		Elements: []*Element{},
	}

	el := NewElement("text", "good")

	assert.Equal(t, expected, el)
}
