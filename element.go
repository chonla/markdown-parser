package parser

import (
	"regexp"
)

// ElementHierarchy provides hierachical structure
var ElementHierarchy = map[string]int{
	"doc":  0,
	"h1":   10,
	"h2":   20,
	"h3":   30,
	"h4":   40,
	"h5":   50,
	"h6":   60,
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
func NewElement(elType, text string) *Element {
	return &Element{
		Parent:   nil,
		Text:     text,
		Type:     elType,
		Elements: []*Element{},
	}
}

// Append element to current element
func (e *Element) Append(el *Element) {
	e.Elements = append(e.Elements, el)
}

func createElement(block string) *Element {
	if text, ok := tryH1(block); ok {
		return NewElement("h1", text)
	}
	if text, ok := tryH2(block); ok {
		return NewElement("h2", text)
	}
	return NewElement("text", block)
}

func testLinePattern(pat, text string) (string, bool) {
	re := regexp.MustCompile(pat)
	m := re.FindAllStringSubmatch(text, -1)
	if len(m) > 0 {
		return m[0][1], true
	}
	return "", false
}

func tryH1(block string) (string, bool) {
	if text, ok := testLinePattern("^# (.+)$", block); ok {
		return text, ok
	}
	if text, ok := testLinePattern("^(.+)\n==+$", block); ok {
		return text, ok
	}
	return "", false
}

func tryH2(block string) (string, bool) {
	if text, ok := testLinePattern("^## (.+)$", block); ok {
		return text, ok
	}
	if text, ok := testLinePattern("^(.+)\n--+$", block); ok {
		return text, ok
	}
	return "", false
}

func tryH3(block string) (string, bool) {
	if text, ok := testLinePattern("^### (.+)$", block); ok {
		return text, ok
	}
	return "", false
}

func tryH4(block string) (string, bool) {
	if text, ok := testLinePattern("^#### (.+)$", block); ok {
		return text, ok
	}
	return "", false
}

func tryH5(block string) (string, bool) {
	if text, ok := testLinePattern("^##### (.+)$", block); ok {
		return text, ok
	}
	return "", false
}

func tryH6(block string) (string, bool) {
	if text, ok := testLinePattern("^###### (.+)$", block); ok {
		return text, ok
	}
	return "", false
}

func tryCode(block string) (string, bool) {
	if text, ok := testLinePattern("(?s)^```[^\n]*\n(.+)\n```$", block); ok {
		return text, ok
	}
	if text, ok := testLinePattern("(?s)^~~~[^\n]*\n(.+)\n~~~$", block); ok {
		return text, ok
	}
	return "", false
}
