package plugin

import (
	"github.com/fatih/structtag"
	"github.com/timmonfette1/protago/internal/genproto/bson"
	"github.com/timmonfette1/protago/internal/genproto/validate"
	pgs "github.com/timmonfette1/protoc-gen-star/v2"
	pgsgo "github.com/timmonfette1/protoc-gen-star/v2/lang/go"
)

type StructTags map[string]map[string]*structtag.Tags

type tagExtractor struct {
	pgs.Visitor
	pgs.DebuggerCommon
	pgsgo.Context
	tags StructTags
}

func newTagExtactor(debug pgs.DebuggerCommon, ctx pgsgo.Context) *tagExtractor {
	te := &tagExtractor{
		DebuggerCommon: debug,
		Context:        ctx,
	}
	te.Visitor = pgs.PassThroughVisitor(te)

	return te
}

func (te *tagExtractor) VisitField(field pgs.Field) (pgs.Visitor, error) {
	messageName := te.Context.Name(field.Message()).String()
	if te.tags[messageName] == nil {
		te.tags[messageName] = map[string]*structtag.Tags{}
	}

	tags := structtag.Tags{}

	// BSON tags
	var bsonFieldOptions *bson.BsonFieldOptions
	_, err := field.Extension(bson.E_Options, &bsonFieldOptions)
	if err != nil {
		return nil, err
	}

	bsonTags, err := convertBsonToTag(bsonFieldOptions)
	te.CheckErr(err)
	for _, tag := range bsonTags.Tags() {
		err := tags.Set(tag)
		if err != nil {
			te.DebuggerCommon.Fail("Error with generating BSON tags: ", err)
		}
	}

	// Validate tags
	var validateFieldOptions *validate.ValidateFieldOptions
	_, err = field.Extension(validate.E_Options, &validateFieldOptions)
	if err != nil {
		return nil, err
	}

	validateTags, err := convertValidateToTag(validateFieldOptions)
	te.CheckErr(err)
	for _, tag := range validateTags.Tags() {
		err := tags.Set(tag)
		if err != nil {
			te.DebuggerCommon.Fail("Error with generating BSON tags: ", err)
		}
	}

	te.tags[messageName][te.Context.Name(field).String()] = &tags

	return te, nil
}

func (te *tagExtractor) Extract(file pgs.File) StructTags {
	te.tags = map[string]map[string]*structtag.Tags{}
	err := pgs.Walk(te, file)
	te.CheckErr(err)
	return te.tags
}
