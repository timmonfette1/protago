package plugin

import (
	"fmt"

	"github.com/fatih/structtag"
	"github.com/timmonfette1/protago/internal/generators"
	pgs "github.com/timmonfette1/protoc-gen-star/v2"
	pgsgo "github.com/timmonfette1/protoc-gen-star/v2/lang/go"
)

type StructTags map[string]map[string]*structtag.Tags

type tagExtractor struct {
	pgs.Visitor
	pgs.DebuggerCommon
	pgsgo.Context
	generators *generators.GeneratorSet
	tags       StructTags
}

func newTagExtactor(debug pgs.DebuggerCommon, ctx pgsgo.Context) *tagExtractor {
	te := &tagExtractor{
		DebuggerCommon: debug,
		Context:        ctx,
		generators:     &generators.GeneratorSet{},
	}
	te.Visitor = pgs.PassThroughVisitor(te)

	return te
}

func (te *tagExtractor) registerGenerator(t generators.Generator) *tagExtractor {
	te.generators.Add(t)
	return te
}

func (te *tagExtractor) VisitField(field pgs.Field) (pgs.Visitor, error) {
	messageName := te.Context.Name(field.Message()).String()
	if te.tags[messageName] == nil {
		te.tags[messageName] = map[string]*structtag.Tags{}
	}

	tags := structtag.Tags{}
	for _, t := range te.generators.GetGenerators() {
		tag, err := t.GenerateTag(field)
		te.CheckErr(err)
		if tag == nil {
			continue
		}

		te.Debug("adding tag:", tag.String())

		err = tags.Set(tag)
		if err != nil {
			te.DebuggerCommon.Fail(fmt.Sprintf("Error with generating '%s' tags: %v", t.GetName(), err))
		}
	}

	te.tags[messageName][te.Context.Name(field).String()] = &tags

	return te, nil
}

func (te *tagExtractor) Extract(file pgs.File) StructTags {
	te.tags = StructTags{}
	err := pgs.Walk(te, file)
	te.CheckErr(err)
	return te.tags
}
