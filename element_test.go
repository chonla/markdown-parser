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

func TestTry1ColumnTable(t *testing.T) {
	content := "| Header |\n| --- |\n| Body |"

	table, success := tryTable(content)

	assert.True(t, success)
	assert.Equal(t, [][]string{
		[]string{
			"Header",
		},
		[]string{
			"Body",
		},
	}, table)
}

func TestTry2ColumnTable(t *testing.T) {
	content := "| Header 1 | Header 2 |\n| --- | --- |\n| Body 1 | Body 2 |"

	table, success := tryTable(content)

	assert.True(t, success)
	assert.Equal(t, [][]string{
		[]string{
			"Header 1",
			"Header 2",
		},
		[]string{
			"Body 1",
			"Body 2",
		},
	}, table)
}

func TestTryLargeColumnTable(t *testing.T) {
	content := "| Header 1 | Header 2 | Header 3 | Header 4 |\n| --- | --- | --- | --- |\n| Body 1 Row 1 | Body 2 Row 1 | Body 3 Row 1 | Body 4 Row 1 |\n| Body 1 Row 2 | Body 2 Row 2 | Body 3 Row 2 | Body 4 Row 2 |\n| Body 1 Row 3 | Body 2 Row 3 | Body 3 Row 3 | Body 4 Row 3 |"

	table, success := tryTable(content)

	assert.True(t, success)
	assert.Equal(t, [][]string{
		[]string{
			"Header 1",
			"Header 2",
			"Header 3",
			"Header 4",
		},
		[]string{
			"Body 1 Row 1",
			"Body 2 Row 1",
			"Body 3 Row 1",
			"Body 4 Row 1",
		},
		[]string{
			"Body 1 Row 2",
			"Body 2 Row 2",
			"Body 3 Row 2",
			"Body 4 Row 2",
		},
		[]string{
			"Body 1 Row 3",
			"Body 2 Row 3",
			"Body 3 Row 3",
			"Body 4 Row 3",
		},
	}, table)
}

func TestTryUnorderedList(t *testing.T) {
	content := "* test 1\n* test 2\n* test 3"

	list, success := tryUnorderedList(content)

	assert.True(t, success)
	assert.Equal(t, []string{
		"test 1",
		"test 2",
		"test 3",
	}, list)
}

func TestTryOrderedList(t *testing.T) {
	content := "213411. test 1\n0. test 2\n12387192837182738127391287398123. test 3"

	list, success := tryOrderedList(content)

	assert.True(t, success)
	assert.Equal(t, []string{
		"test 1",
		"test 2",
		"test 3",
	}, list)
}
