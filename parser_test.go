package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEmptyContent(t *testing.T) {
	content := ""

	result := Parse(content)

	assert.Equal(t, Document{}, result)
}
