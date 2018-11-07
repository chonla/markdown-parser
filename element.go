package parser

import (
	"fmt"
	"regexp"
	"strings"
)

// ElementHierarchy provides hierachical structure
var ElementHierarchy = map[string]int{
	"doc":            0,
	"h1":             10,
	"h2":             20,
	"h3":             30,
	"h4":             40,
	"h5":             50,
	"h6":             60,
	"code":           100,
	"table":          100,
	"text":           100,
	"unordered-list": 100,
	"ordered-list":   100,
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

// NewTable creates a new table
func NewTable(table [][]string) *Element {
	tableElement := &Element{
		Parent:   nil,
		Text:     "",
		Type:     "table",
		Elements: []*Element{},
	}

	for _, row := range table {
		rowElement := &Element{
			Parent:   tableElement,
			Text:     "",
			Type:     "row",
			Elements: []*Element{},
		}
		for _, cell := range row {
			cellElement := &Element{
				Parent:   rowElement,
				Text:     cell,
				Type:     "cell",
				Elements: []*Element{},
			}
			rowElement.Append(cellElement)
		}
		tableElement.Append(rowElement)
	}

	return tableElement
}

// NewUnorderedList creates unordered list
func NewUnorderedList(list []string) *Element {
	listElement := &Element{
		Parent:   nil,
		Text:     "",
		Type:     "unordered-list",
		Elements: []*Element{},
	}

	for _, item := range list {
		itemElement := &Element{
			Parent:   listElement,
			Text:     item,
			Type:     "list-item",
			Elements: []*Element{},
		}

		listElement.Append(itemElement)
	}

	return listElement
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
	if text, ok := tryH3(block); ok {
		return NewElement("h3", text)
	}
	if text, ok := tryH4(block); ok {
		return NewElement("h4", text)
	}
	if text, ok := tryH5(block); ok {
		return NewElement("h5", text)
	}
	if text, ok := tryH6(block); ok {
		return NewElement("h6", text)
	}
	if text, ok := tryCode(block); ok {
		return NewElement("code", text)
	}
	if table, ok := tryTable(block); ok {
		return NewTable(table)
	}
	if list, ok := tryUnorderedList(block); ok {
		return NewUnorderedList(list)
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

func tryUnorderedList(block string) ([]string, bool) {
	output := []string{}
	lines := strings.Split(block, "\n")

	for _, line := range lines {
		if text, ok := testLinePattern("^\\* (.+)$", line); ok {
			output = append(output, text)
		} else {
			return nil, false
		}
	}

	return output, true
}

func tryTable(block string) ([][]string, bool) {
	output := [][]string{}

	lines := strings.Split(block, "\n")
	if len(lines) < 3 {
		fmt.Println("not a table")
		return nil, false
	}

	// check header separator
	// header separator for one column
	patSep := "^\\|( :?---+? \\|)+$"
	reSep := regexp.MustCompile(patSep)
	if !reSep.MatchString(lines[1]) {
		// header separator does not present
		fmt.Println("header separator does not present")
		return nil, false
	}

	colCount := columnCount(lines[1])

	if colCount != columnCount(lines[0]) {
		// header and column count does not match
		fmt.Println("header and column count does not match")
		return nil, false
	}

	row := []string{}
	for i := 0; i < colCount; i++ {
		row = append(row, getCellValue(i, lines[0]))
	}
	output = append(output, row)

	for i, n := 2, len(lines); i < n; i++ {
		row = []string{}
		for j := 0; j < colCount; j++ {
			row = append(row, getCellValue(j, lines[i]))
		}
		output = append(output, row)
	}

	return output, true
}

func getCellValue(index int, line string) string {
	cols := strings.Split(line, "|")
	if cols[0] == "" { // detect left boundary pipe
		cols = cols[1:]
	}
	if cols[len(cols)-1] == "" { // detect right boundary pipe
		cols = cols[0 : len(cols)-1]
	}

	return strings.TrimSpace(cols[index])
}

func columnCount(line string) int {
	colCount := strings.Count(line, "|") - 1
	if line[0] != '|' || line[len(line)-1] != '|' {
		colCount++
	}
	return colCount
}
