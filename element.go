package parser

// ElementHierarchy provides hierachical structure
var ElementHierarchy = map[string]int{
	"doc":  0,
	"h1":   10,
	"h2":   20,
	"text": 1000,
}

// Element represents element in markdown document
type Element struct {
	Text     string
	Type     string
	Parent   *Element
	Elements []*Element
}

// NewElement creates a new element
func NewElement() *Element {
	return &Element{
		Text:     "",
		Type:     "text",
		Elements: []*Element{},
	}
}

// Append element to current element
func (e *Element) Append(el *Element) {
	e.Elements = append(e.Elements, el)
}
