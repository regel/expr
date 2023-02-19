package ast

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseFunction(t *testing.T) {
	var buf bytes.Buffer
	expected := `add
  ,
    1
    2
`
	code := `add((1) , (2))`
	ast, err := ParseExpr(code)
	require.NoError(t, err, "ParseExpr returned an error")
	if err != nil {
		t.Error("got error")
	}
	PrettyPrint(&buf, ast, "")
	actual := buf.String()
	assert.Equal(t, expected, actual)
}
