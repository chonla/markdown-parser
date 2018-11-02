package parser

import (
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
