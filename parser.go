package parser

import (
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
	return Element{
		Text:     line,
		Elements: []Element{},
	}
}

// func tryH1(line string) (string, bool) {
// 	re := regexp.MustCompile("^# (.+)$")
// 	m := re.FindAllString(line, -1)
// 	if len(m) > 0 {
// 		return m[1], true
// 	}
// 	return "", false
// }
