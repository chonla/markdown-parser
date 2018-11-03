package parser

import (
	"github.com/kr/pretty"
)

// Parse markdown text to document
func Parse(content string) *Document {
	doc := NewDocument()
	var cursor = doc.Element
	tokenizer := NewTokenizer()

	blocks := tokenizer.Tokenize(content)

	pretty.Println(blocks)

	for _, block := range blocks {
		if block != "" {
			element := createElement(block)
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
