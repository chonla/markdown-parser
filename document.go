package parser

// Document holds markdown document
type Document struct {
	Elements []Element
}

// NewDocument creates a document
func NewDocument() Document {
	return Document{
		Elements: []Element{},
	}
}

// Append element to element list
func (d *Document) Append(el Element) {
	d.Elements = append(d.Elements, el)
}
