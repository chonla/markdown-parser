package parser

// Document holds markdown document
type Document struct {
	*Element
}

// NewDocument creates a document
func NewDocument() *Document {
	return &Document{
		Element: &Element{
			Type:     "doc",
			Elements: []*Element{},
		},
	}
}

// Append element to element list
func (d *Document) Append(el *Element) {
	d.Element.Append(el)
}
