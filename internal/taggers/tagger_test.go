package taggers

import (
	"testing"

	"github.com/fatih/structtag"
	"github.com/stretchr/testify/assert"
	pgs "github.com/timmonfette1/protoc-gen-star/v2"
	pgsgo "github.com/timmonfette1/protoc-gen-star/v2/lang/go"
)

var _ Tagger = (*dummyTagger)(nil)
var _ Tagger = (*validTagger)(nil)

type dummyTagger struct {
	pgs.DebuggerCommon
	pgsgo.Context
}
type validTagger struct {
	pgs.DebuggerCommon
	pgsgo.Context
}

func (d *dummyTagger) GetName() string {
	return ""
}
func (d *dummyTagger) GenerateTag(field pgs.Field) (*structtag.Tag, error) {
	return nil, nil
}
func (v *validTagger) GetName() string {
	return "UnitTesting"
}
func (v *validTagger) GenerateTag(field pgs.Field) (*structtag.Tag, error) {
	return nil, nil
}

func TestTaggerSet(t *testing.T) {
	ts := TaggerSet{}
	err := ts.Add(nil)
	assert.Error(t, err)

	err = ts.Add(&dummyTagger{})
	assert.Error(t, err)

	err = ts.Add(&validTagger{})
	assert.NoError(t, err)

	// Add again, shouldn't error, should only replace existing Tagger.
	err = ts.Add(&validTagger{})
	assert.NoError(t, err)
}
