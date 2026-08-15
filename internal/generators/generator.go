package generators

import (
	"fmt"

	"github.com/fatih/structtag"
	pgs "github.com/timmonfette1/protoc-gen-star/v2"
	pgsgo "github.com/timmonfette1/protoc-gen-star/v2/lang/go"
)

type Generator interface {
	pgs.DebuggerCommon
	pgsgo.Context

	GetName() string
	GenerateTag(field pgs.Field) (*structtag.Tag, error)
}

type GeneratorSet struct {
	generators []Generator
}

func (ts *GeneratorSet) GetGenerators() []Generator {
	return ts.generators
}

func (ts *GeneratorSet) Add(t Generator) error {
	if t == nil {
		return fmt.Errorf("generator cannot be nil")
	}
	if t.GetName() == "" {
		return fmt.Errorf("'name' cannot be empty")
	}

	added := false
	for i, tg := range ts.generators {
		if tg != nil && tg.GetName() == t.GetName() {
			added = true
			ts.generators[i] = t
		}
	}

	if !added {
		ts.generators = append(ts.generators, t)
	}

	return nil
}
