package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewElement(t *testing.T) {
	expected := &Element{
		Text:     "",
		Type:     "text",
		Elements: []*Element{},
	}

	el := NewElement()

	assert.Equal(t, expected, el)
}
