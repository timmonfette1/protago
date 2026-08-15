package generators

import (
	"testing"

	"github.com/fatih/structtag"
	"github.com/stretchr/testify/assert"
	pgs "github.com/timmonfette1/protoc-gen-star/v2"
	pgsgo "github.com/timmonfette1/protoc-gen-star/v2/lang/go"
)

var _ Generator = (*dummyGenerator)(nil)
var _ Generator = (*validGenerator)(nil)

type dummyGenerator struct {
	pgs.DebuggerCommon
	pgsgo.Context
}
type validGenerator struct {
	pgs.DebuggerCommon
	pgsgo.Context
}

func (d *dummyGenerator) GetName() string {
	return ""
}
func (d *dummyGenerator) GenerateTag(field pgs.Field) (*structtag.Tag, error) {
	return nil, nil
}
func (v *validGenerator) GetName() string {
	return "UnitTesting"
}
func (v *validGenerator) GenerateTag(field pgs.Field) (*structtag.Tag, error) {
	return nil, nil
}

func TestGeneratorSet(t *testing.T) {
	ts := GeneratorSet{}
	err := ts.Add(nil)
	assert.Error(t, err)

	err = ts.Add(&dummyGenerator{})
	assert.Error(t, err)

	err = ts.Add(&validGenerator{})
	assert.NoError(t, err)

	// Add again, shouldn't error, should only replace existing Tagger.
	err = ts.Add(&validGenerator{})
	assert.NoError(t, err)
}
