package parser

import (
	"regexp"
	"strings"
)

// Parse markdown text to document
func Parse(content string) *Document {
	doc := NewDocument()
	var cursor = doc.Element

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if line != "" {
			element := createElement(line)
			for ElementHierarchy[cursor.Type] >= ElementHierarchy[element.Type] {
				cursor = cursor.Parent
			}

			element.Parent = cursor
			cursor.Append(element)
			cursor = element
		}
	}

	return doc
}

func createElement(line string) *Element {
	if text, ok := tryH1(line); ok {
		return &Element{
			Text:     text,
			Type:     "h1",
			Parent:   nil,
			Elements: []*Element{},
		}
	}
	if text, ok := tryH2(line); ok {
		return &Element{
			Text:     text,
			Type:     "h2",
			Parent:   nil,
			Elements: []*Element{},
		}
	}
	return &Element{
		Text:     line,
		Type:     "text",
		Parent:   nil,
		Elements: []*Element{},
	}
}

func tryH1(line string) (string, bool) {
	re := regexp.MustCompile("^# (.+)$")
	m := re.FindAllStringSubmatch(line, -1)
	if len(m) > 0 {
		return m[0][1], true
	}
	return "", false
}

func tryH2(line string) (string, bool) {
	re := regexp.MustCompile("^## (.+)$")
	m := re.FindAllStringSubmatch(line, -1)
	if len(m) > 0 {
		return m[0][1], true
	}
	return "", false
}
