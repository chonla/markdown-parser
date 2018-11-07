package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEmptyContent(t *testing.T) {
	content := ""
	expected := &Document{
		Element: &Element{
			Type:     "doc",
			Elements: []*Element{},
		},
	}

	result := Parse(content)

	assert.Equal(t, expected, result)
}

func TestParseEmptyLineContent(t *testing.T) {
	content := "\n\n\n"
	expected := &Document{
		Element: &Element{
			Type:     "doc",
			Elements: []*Element{},
		},
	}

	result := Parse(content)

	assert.Equal(t, expected, result)
}

func TestParseSimpleDocument(t *testing.T) {
	content := "Test"

	Text := &Element{
		Text:     "Test",
		Type:     "text",
		Elements: []*Element{},
	}

	Doc := &Element{
		Type: "doc",
		Elements: []*Element{
			Text,
		},
	}

	Text.Parent = Doc

	expected := &Document{
		Element: Doc,
	}

	result := Parse(content)

	assert.Equal(t, expected, result)
}

func TestParseSimpleDocumentWith2SimpleParagraph(t *testing.T) {
	content := "Test\n\nTest2"

	Text1 := &Element{
		Text:     "Test",
		Type:     "text",
		Elements: []*Element{},
	}

	Text2 := &Element{
		Text:     "Test2",
		Type:     "text",
		Elements: []*Element{},
	}

	Doc := &Element{
		Type: "doc",
		Elements: []*Element{
			Text1,
			Text2,
		},
	}

	Text1.Parent = Doc
	Text2.Parent = Doc

	expected := &Document{
		Element: Doc,
	}

	result := Parse(content)

	assert.Equal(t, expected, result)
}

func TestParseSimpleDocumentWith1SimpleParagraphWithNewLine(t *testing.T) {
	content := "Test\nTest2"

	Text1 := &Element{
		Text:     "Test\nTest2",
		Type:     "text",
		Elements: []*Element{},
	}

	Doc := &Element{
		Type: "doc",
		Elements: []*Element{
			Text1,
		},
	}

	Text1.Parent = Doc

	expected := &Document{
		Element: Doc,
	}

	result := Parse(content)

	assert.Equal(t, expected, result)
}

func TestParseH1Document(t *testing.T) {
	content := "# Title"

	H1 := &Element{
		Text:     "Title",
		Type:     "h1",
		Elements: []*Element{},
	}

	Doc := &Element{
		Type: "doc",
		Elements: []*Element{
			H1,
		},
	}

	H1.Parent = Doc

	expected := &Document{
		Element: Doc,
	}

	result := Parse(content)

	assert.Equal(t, expected, result)
}

func TestParseAlternateH1Document(t *testing.T) {
	content := "Title\n=="

	H1 := &Element{
		Text:     "Title",
		Type:     "h1",
		Elements: []*Element{},
	}

	Doc := &Element{
		Type: "doc",
		Elements: []*Element{
			H1,
		},
	}

	H1.Parent = Doc

	expected := &Document{
		Element: Doc,
	}

	result := Parse(content)

	assert.Equal(t, expected, result)
}

func TestParseH2Document(t *testing.T) {
	content := "## Title"

	H2 := &Element{
		Text:     "Title",
		Type:     "h2",
		Elements: []*Element{},
	}

	Doc := &Element{
		Type: "doc",
		Elements: []*Element{
			H2,
		},
	}

	H2.Parent = Doc

	expected := &Document{
		Element: Doc,
	}

	result := Parse(content)

	assert.Equal(t, expected, result)
}

func TestParseAlternateH2Document(t *testing.T) {
	content := "Title\n--"

	H2 := &Element{
		Text:     "Title",
		Type:     "h2",
		Elements: []*Element{},
	}

	Doc := &Element{
		Type: "doc",
		Elements: []*Element{
			H2,
		},
	}

	H2.Parent = Doc

	expected := &Document{
		Element: Doc,
	}

	result := Parse(content)

	assert.Equal(t, expected, result)
}

func TestParseH1H2DocumentWithHierarchy(t *testing.T) {
	content := "# H1 Title\n\n## H2 Title"

	H2 := &Element{
		Text:     "H2 Title",
		Type:     "h2",
		Elements: []*Element{},
	}

	H1 := &Element{
		Text: "H1 Title",
		Type: "h1",
		Elements: []*Element{
			H2,
		},
	}

	Doc := &Element{
		Type: "doc",
		Elements: []*Element{
			H1,
		},
	}

	H2.Parent = H1
	H1.Parent = Doc

	expected := &Document{
		Element: Doc,
	}

	result := Parse(content)

	assert.Equal(t, expected, result)
}

func TestParseTableDocument(t *testing.T) {
	content := "# Table Document\n\n| Header 1 | Header 2 | Header 3 | Header 4 |\n| --- | --- | --- | --- |\n| Body 1 Row 1 | Body 2 Row 1 | Body 3 Row 1 | Body 4 Row 1 |\n| Body 1 Row 2 | Body 2 Row 2 | Body 3 Row 2 | Body 4 Row 2 |\n| Body 1 Row 3 | Body 2 Row 3 | Body 3 Row 3 | Body 4 Row 3 |"

	Row1Col1 := &Element{
		Type:     "cell",
		Text:     "Header 1",
		Parent:   nil,
		Elements: []*Element{},
	}

	Row1Col2 := &Element{
		Type:     "cell",
		Text:     "Header 2",
		Parent:   nil,
		Elements: []*Element{},
	}

	Row1Col3 := &Element{
		Type:     "cell",
		Text:     "Header 3",
		Parent:   nil,
		Elements: []*Element{},
	}

	Row1Col4 := &Element{
		Type:     "cell",
		Text:     "Header 4",
		Parent:   nil,
		Elements: []*Element{},
	}

	Row2Col1 := &Element{
		Type:     "cell",
		Text:     "Body 1 Row 1",
		Parent:   nil,
		Elements: []*Element{},
	}

	Row2Col2 := &Element{
		Type:     "cell",
		Text:     "Body 2 Row 1",
		Parent:   nil,
		Elements: []*Element{},
	}

	Row2Col3 := &Element{
		Type:     "cell",
		Text:     "Body 3 Row 1",
		Parent:   nil,
		Elements: []*Element{},
	}

	Row2Col4 := &Element{
		Type:     "cell",
		Text:     "Body 4 Row 1",
		Parent:   nil,
		Elements: []*Element{},
	}

	Row3Col1 := &Element{
		Type:     "cell",
		Text:     "Body 1 Row 2",
		Parent:   nil,
		Elements: []*Element{},
	}

	Row3Col2 := &Element{
		Type:     "cell",
		Text:     "Body 2 Row 2",
		Parent:   nil,
		Elements: []*Element{},
	}

	Row3Col3 := &Element{
		Type:     "cell",
		Text:     "Body 3 Row 2",
		Parent:   nil,
		Elements: []*Element{},
	}

	Row3Col4 := &Element{
		Type:     "cell",
		Text:     "Body 4 Row 2",
		Parent:   nil,
		Elements: []*Element{},
	}

	Row4Col1 := &Element{
		Type:     "cell",
		Text:     "Body 1 Row 3",
		Parent:   nil,
		Elements: []*Element{},
	}

	Row4Col2 := &Element{
		Type:     "cell",
		Text:     "Body 2 Row 3",
		Parent:   nil,
		Elements: []*Element{},
	}

	Row4Col3 := &Element{
		Type:     "cell",
		Text:     "Body 3 Row 3",
		Parent:   nil,
		Elements: []*Element{},
	}

	Row4Col4 := &Element{
		Type:     "cell",
		Text:     "Body 4 Row 3",
		Parent:   nil,
		Elements: []*Element{},
	}

	Row1Cols := []*Element{
		Row1Col1,
		Row1Col2,
		Row1Col3,
		Row1Col4,
	}

	Row2Cols := []*Element{
		Row2Col1,
		Row2Col2,
		Row2Col3,
		Row2Col4,
	}

	Row3Cols := []*Element{
		Row3Col1,
		Row3Col2,
		Row3Col3,
		Row3Col4,
	}

	Row4Cols := []*Element{
		Row4Col1,
		Row4Col2,
		Row4Col3,
		Row4Col4,
	}

	Row1 := &Element{
		Type:     "row",
		Text:     "",
		Parent:   nil,
		Elements: Row1Cols,
	}

	Row2 := &Element{
		Type:     "row",
		Text:     "",
		Parent:   nil,
		Elements: Row2Cols,
	}

	Row3 := &Element{
		Type:     "row",
		Text:     "",
		Parent:   nil,
		Elements: Row3Cols,
	}

	Row4 := &Element{
		Type:     "row",
		Text:     "",
		Parent:   nil,
		Elements: Row4Cols,
	}

	Rows := []*Element{
		Row1,
		Row2,
		Row3,
		Row4,
	}

	Table := &Element{
		Text:     "",
		Type:     "table",
		Elements: Rows,
	}

	H1 := &Element{
		Text: "Table Document",
		Type: "h1",
		Elements: []*Element{
			Table,
		},
	}

	Doc := &Element{
		Type: "doc",
		Elements: []*Element{
			H1,
		},
	}

	H1.Parent = Doc
	Table.Parent = H1
	Row1.Parent = Table
	Row2.Parent = Table
	Row3.Parent = Table
	Row4.Parent = Table
	Row1Col1.Parent = Row1
	Row1Col2.Parent = Row1
	Row1Col3.Parent = Row1
	Row1Col4.Parent = Row1
	Row2Col1.Parent = Row2
	Row2Col2.Parent = Row2
	Row2Col3.Parent = Row2
	Row2Col4.Parent = Row2
	Row3Col1.Parent = Row3
	Row3Col2.Parent = Row3
	Row3Col3.Parent = Row3
	Row3Col4.Parent = Row3
	Row4Col1.Parent = Row4
	Row4Col2.Parent = Row4
	Row4Col3.Parent = Row4
	Row4Col4.Parent = Row4

	expected := &Document{
		Element: Doc,
	}

	result := Parse(content)

	assert.Equal(t, expected, result)
}

func TestParseUnorderedList(t *testing.T) {
	content := "* item 1\n* item 2\n* item 3"

	ListItem1 := &Element{
		Type:     "list-item",
		Elements: []*Element{},
		Text:     "item 1",
		Parent:   nil,
	}

	ListItem2 := &Element{
		Type:     "list-item",
		Elements: []*Element{},
		Text:     "item 2",
		Parent:   nil,
	}

	ListItem3 := &Element{
		Type:     "list-item",
		Elements: []*Element{},
		Text:     "item 3",
		Parent:   nil,
	}

	List := &Element{
		Type: "unordered-list",
		Text: "",
		Elements: []*Element{
			ListItem1,
			ListItem2,
			ListItem3,
		},
		Parent: nil,
	}

	Doc := &Element{
		Type: "doc",
		Elements: []*Element{
			List,
		},
	}

	List.Parent = Doc
	ListItem1.Parent = List
	ListItem2.Parent = List
	ListItem3.Parent = List

	expected := &Document{
		Element: Doc,
	}

	result := Parse(content)

	assert.Equal(t, expected, result)
}

func TestParseOrderedList(t *testing.T) {
	content := "12. item 1\n1. item 2\n728123123121234511. item 3"

	ListItem1 := &Element{
		Type:     "list-item",
		Elements: []*Element{},
		Text:     "item 1",
		Parent:   nil,
	}

	ListItem2 := &Element{
		Type:     "list-item",
		Elements: []*Element{},
		Text:     "item 2",
		Parent:   nil,
	}

	ListItem3 := &Element{
		Type:     "list-item",
		Elements: []*Element{},
		Text:     "item 3",
		Parent:   nil,
	}

	List := &Element{
		Type: "ordered-list",
		Text: "",
		Elements: []*Element{
			ListItem1,
			ListItem2,
			ListItem3,
		},
		Parent: nil,
	}

	Doc := &Element{
		Type: "doc",
		Elements: []*Element{
			List,
		},
	}

	List.Parent = Doc
	ListItem1.Parent = List
	ListItem2.Parent = List
	ListItem3.Parent = List

	expected := &Document{
		Element: Doc,
	}

	result := Parse(content)

	assert.Equal(t, expected, result)
}
