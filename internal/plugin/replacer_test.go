package plugin

import (
	"bytes"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"testing"

	"github.com/fatih/structtag"
	"github.com/stretchr/testify/assert"
)

func TestReplaceTags(t *testing.T) {
	fs := token.NewFileSet()

	input, err := parser.ParseFile(fs, "../testdata/input.txt", nil, parser.ParseComments)
	assert.NoError(t, err)

	singleTags, err := structtag.Parse(`sql:"-,omitempty"`)
	assert.NoError(t, err)
	multipleTags, err := structtag.Parse(`xml:"-,omitempty" sql:"ke,op" bson:"ke,op"`)
	assert.NoError(t, err)
	noneTags, err := structtag.Parse(`json:"none,omitempty"`)
	assert.NoError(t, err)

	err = replaceTags(input, StructTags{
		"Simple": {
			"Single":   singleTags,
			"Multiple": multipleTags,
			"None":     noneTags,
		},
	})

	var inputBytes bytes.Buffer
	err = printer.Fprint(&inputBytes, fs, input)
	assert.NoError(t, err)

	outputBytes, err := os.ReadFile("../testdata/output.txt")
	assert.NoError(t, err)
	assert.Equal(t, outputBytes, inputBytes.Bytes())
}
