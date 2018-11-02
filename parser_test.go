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
	content := "Test\nTest2"

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

func TestParseH1H2DocumentWithHierarchy(t *testing.T) {
	content := "# H1 Title\n## H2 Title"

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
