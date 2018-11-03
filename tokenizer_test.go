package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizeEmptyData(t *testing.T) {
	content := ""
	expected := []string{}
	tokenizer := NewTokenizer()

	result := tokenizer.Tokenize(content)

	assert.Equal(t, expected, result)
}

func TestTokenizeSimpleData(t *testing.T) {
	content := "test"
	expected := []string{"test"}
	tokenizer := NewTokenizer()

	result := tokenizer.Tokenize(content)

	assert.Equal(t, expected, result)
}

func TestTokenizeSimpleDataWith2Lines(t *testing.T) {
	content := "test\ntest2"
	expected := []string{"test\ntest2"}
	tokenizer := NewTokenizer()

	result := tokenizer.Tokenize(content)

	assert.Equal(t, expected, result)
}

func TestTokenizeSimpleDataWith2Blocks(t *testing.T) {
	content := "test\n\ntest2"
	expected := []string{"test", "test2"}
	tokenizer := NewTokenizer()

	result := tokenizer.Tokenize(content)

	assert.Equal(t, expected, result)
}

func TestTokenizeSimpleDataWithVeryLongBlockSeparator(t *testing.T) {
	content := "test\n\n\n\n\n\n\n\n\n\n\n\n\n\n\ntest2\n\n\n\n\n\n\n\n\n\n\n\n\n\n\ntest3\ntest4"
	expected := []string{"test", "test2", "test3\ntest4"}
	tokenizer := NewTokenizer()

	result := tokenizer.Tokenize(content)

	assert.Equal(t, expected, result)
}

func TestTokenizeSimpleDataWithEmptyContentBeforeSeparator(t *testing.T) {
	content := "\n\n\n\n\n\n\n\n\n\n\n\n\n\n\ntest2\n\n\n\n\n\n\n\n\n\n\n\n\n\n\ntest3\ntest4"
	expected := []string{"test2", "test3\ntest4"}
	tokenizer := NewTokenizer()

	result := tokenizer.Tokenize(content)

	assert.Equal(t, expected, result)
}

func TestTokenizeSimpleDataWithEmptyContentAfterSeparator(t *testing.T) {
	content := "test\n\n\n\n\n\n\n\n\n\n\n\n\n\n\ntest2\n\n\n\n\n\n\n\n\n\n\n\n\n\n\ntest3\ntest4\n\n\n\n\n\n"
	expected := []string{"test", "test2", "test3\ntest4"}
	tokenizer := NewTokenizer()

	result := tokenizer.Tokenize(content)

	assert.Equal(t, expected, result)
}

func TestTokenizeSimpleDataWith1CodeBlock(t *testing.T) {
	content := "```\ntest\n\ntest2\n```"
	expected := []string{"```\ntest\n\ntest2\n```"}
	tokenizer := NewTokenizer()

	result := tokenizer.Tokenize(content)

	assert.Equal(t, expected, result)
}

func TestTokenizeSimpleDataWith1CodeBlockAnd1Line(t *testing.T) {
	content := "```\ntest\n\ntest2\n```\n\nanother line"
	expected := []string{"```\ntest\n\ntest2\n```", "another line"}
	tokenizer := NewTokenizer()

	result := tokenizer.Tokenize(content)

	assert.Equal(t, expected, result)
}

func TestTokenizeSimpleDataWith2CodeBlocks(t *testing.T) {
	content := "```\ntest\n\ntest2\n```\n\n```\nanother block\n\n\nline in block\n```"
	expected := []string{"```\ntest\n\ntest2\n```", "```\nanother block\n\n\nline in block\n```"}
	tokenizer := NewTokenizer()

	result := tokenizer.Tokenize(content)

	assert.Equal(t, expected, result)
}

func TestTokenizeSimpleDataWith1NamedCodeBlock(t *testing.T) {
	content := "```javascript\ntest\n\ntest2\n```"
	expected := []string{"```javascript\ntest\n\ntest2\n```"}
	tokenizer := NewTokenizer()

	result := tokenizer.Tokenize(content)

	assert.Equal(t, expected, result)
}

func TestTokenizeSimpleDataWith1AlternativeCodeBlock(t *testing.T) {
	content := "~~~\ntest\n\ntest2\n~~~"
	expected := []string{"~~~\ntest\n\ntest2\n~~~"}
	tokenizer := NewTokenizer()

	result := tokenizer.Tokenize(content)

	assert.Equal(t, expected, result)
}
