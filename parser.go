package parser

import (
	"regexp"
	"strings"
)

func Parse(content string) Document {
	doc := NewDocument()

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if line != "" {
			element := createElement(line)
			doc.Append(element)
		}
	}
	return doc
}

func createElement(line string) Element {
	if text, ok := tryH1(line); ok {
		return Element{
			Text:     text,
			Type:     "H1",
			Elements: []Element{},
		}
	}
	if text, ok := tryH2(line); ok {
		return Element{
			Text:     text,
			Type:     "H2",
			Elements: []Element{},
		}
	}
	return Element{
		Text:     line,
		Type:     "Text",
		Elements: []Element{},
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
