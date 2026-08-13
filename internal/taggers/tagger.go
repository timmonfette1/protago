package taggers

import (
	"fmt"

	"github.com/fatih/structtag"
	pgs "github.com/timmonfette1/protoc-gen-star/v2"
	pgsgo "github.com/timmonfette1/protoc-gen-star/v2/lang/go"
)

type Tagger interface {
	pgs.DebuggerCommon
	pgsgo.Context

	GetName() string
	GenerateTag(field pgs.Field) (*structtag.Tag, error)
}

type TaggerSet struct {
	taggers []Tagger
}

func (ts *TaggerSet) GetTaggers() []Tagger {
	return ts.taggers
}

func (ts *TaggerSet) Add(t Tagger) error {
	if t == nil {
		return fmt.Errorf("tagger cannot be nil")
	}
	if t.GetName() == "" {
		return fmt.Errorf("'name' cannot be empty")
	}

	added := false
	for i, tg := range ts.taggers {
		if tg != nil && tg.GetName() == t.GetName() {
			added = true
			ts.taggers[i] = t
		}
	}

	if !added {
		ts.taggers = append(ts.taggers, t)
	}

	return nil
}
