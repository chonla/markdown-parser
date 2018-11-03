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

func TestTryH3(t *testing.T) {
	content := "### Title"

	text, success := tryH3(content)

	assert.True(t, success)
	assert.Equal(t, "Title", text)
}

func TestTryH4(t *testing.T) {
	content := "#### Title"

	text, success := tryH4(content)

	assert.True(t, success)
	assert.Equal(t, "Title", text)
}

func TestTryH5(t *testing.T) {
	content := "##### Title"

	text, success := tryH5(content)

	assert.True(t, success)
	assert.Equal(t, "Title", text)
}

func TestTryH6(t *testing.T) {
	content := "###### Title"

	text, success := tryH6(content)

	assert.True(t, success)
	assert.Equal(t, "Title", text)
}

func TestTryCodeBlock(t *testing.T) {
	content := "```\nContent\nWith\nNew Line\n```"

	text, success := tryCode(content)

	assert.True(t, success)
	assert.Equal(t, "Content\nWith\nNew Line", text)
}

func TestTryAlternativeCodeBlock(t *testing.T) {
	content := "~~~\nContent\nWith\nNew Line\n~~~"

	text, success := tryCode(content)

	assert.True(t, success)
	assert.Equal(t, "Content\nWith\nNew Line", text)
}
