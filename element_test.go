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

func TestTryH1(t *testing.T) {
	content := "# Title"

	text, success := tryH1(content)

	assert.True(t, success)
	assert.Equal(t, "Title", text)
}

func TestTryAlternateH1(t *testing.T) {
	content := "Title\n=="

	text, success := tryH1(content)

	assert.True(t, success)
	assert.Equal(t, "Title", text)
}

func TestTryH2(t *testing.T) {
	content := "## Title"

	text, success := tryH2(content)

	assert.True(t, success)
	assert.Equal(t, "Title", text)
}

func TestTryAlternateH2(t *testing.T) {
	content := "Title\n--"

	text, success := tryH2(content)

	assert.True(t, success)
	assert.Equal(t, "Title", text)
}
