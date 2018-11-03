package parser

import (
	"regexp"
	"strings"
)

// linesBlock defines multilines block open/close pattern
var linesBlock = map[string]string{
	"^```.*$": "^```$",
}

// Tokenizer is markdown block tokenizer
type Tokenizer struct {
	Output []string
	Block  []string
}

// NewTokenizer creates a new tokenizer
func NewTokenizer() *Tokenizer {
	return &Tokenizer{}
}

// Tokenize creates tokens from markdown content
func (t *Tokenizer) Tokenize(content string) []string {
	t.Output = []string{}
	t.Block = []string{}
	blockType := ""

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if blockType != "" {
			t.Block = append(t.Block, line)
			if t.isEndOfLinesBlock(line, blockType) {
				blockType = ""
			}
		} else {
			if line == "" {
				t.flushBlock()
			} else {
				t.Block = append(t.Block, line)
				if detectedType, ok := t.isBeginOfLinesBlock(line); ok {
					blockType = detectedType
				}
			}
		}
	}
	t.flushBlock()

	return t.Output
}

// flushBlock flushes content remained in block to output
func (t *Tokenizer) flushBlock() {
	if len(t.Block) > 0 {
		t.Output = append(t.Output, strings.Join(t.Block, "\n"))
		t.Block = []string{}
	}
}

func (t *Tokenizer) isBeginOfLinesBlock(line string) (string, bool) {
	for pattern := range linesBlock {
		re := regexp.MustCompile(pattern)
		if re.MatchString(line) {
			return pattern, true
		}
	}
	return "", false
}

func (t *Tokenizer) isEndOfLinesBlock(line, blockType string) bool {
	if pattern, ok := linesBlock[blockType]; ok {
		re := regexp.MustCompile(pattern)
		return re.MatchString(line)
	}
	return false
}
