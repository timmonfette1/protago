package taggers

import (
	"fmt"

	"github.com/fatih/structtag"
	"github.com/timmonfette1/protago/internal/genproto/bson"
	pgs "github.com/timmonfette1/protoc-gen-star/v2"
	pgsgo "github.com/timmonfette1/protoc-gen-star/v2/lang/go"
)

type BsonTagger struct {
	pgs.DebuggerCommon
	pgsgo.Context
}

var _ Tagger = (*BsonTagger)(nil)

func (bt *BsonTagger) GetName() string {
	return "BsonTagger"
}

func (bt *BsonTagger) GenerateTag(field pgs.Field) (*structtag.Tag, error) {
	bt.Debug(fmt.Sprintf("parsing 'bson' tags for field '%s'", field.Name().String()))

	var bsonFieldOptions *bson.BsonFieldOptions
	_, err := field.Extension(bson.E_Options, &bsonFieldOptions)
	if err != nil {
		return nil, err
	}

	if bsonFieldOptions == nil {
		return nil, nil
	}

	opts := []string{}
	if bsonFieldOptions.GetOmitempty() {
		opts = append(opts, "omitempty")
	}
	if bsonFieldOptions.GetInline() {
		opts = append(opts, "inline")
	}
	if bsonFieldOptions.GetMinsize() {
		opts = append(opts, "minsize")
	}
	if bsonFieldOptions.GetTruncate() {
		opts = append(opts, "truncate")
	}

	return &structtag.Tag{
		Key:     "bson",
		Name:    bsonFieldOptions.GetFieldName(),
		Options: opts,
	}, nil
}
