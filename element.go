package parser

// Element represents element in markdown document
type Element struct {
	Text     string
	Elements []Element
}

// NewElement creates a new element
func NewElement() Element {
	return Element{
		Text:     "",
		Elements: []Element{},
	}
}
