package plugin

import (
	"github.com/fatih/structtag"
	"github.com/timmonfette1/protago/internal/genproto/bson"
	"github.com/timmonfette1/protago/internal/genproto/validate"
)

func convertBsonToTag(b *bson.BsonFieldOptions) (*structtag.Tags, error) {
	opts := []string{}
	if b.GetOmitEmpty() {
		opts = append(opts, "omitempty")
	}
	if b.GetInline() {
		opts = append(opts, "inline")
	}
	if b.GetMinsize() {
		opts = append(opts, "minsize")
	}
	if b.GetTruncate() {
		opts = append(opts, "truncate")
	}

	st := &structtag.Tag{
		Key:     "bson",
		Name:    b.GetFieldName(),
		Options: opts,
	}

	t := structtag.Tags{}
	t.Set(st)
	return &t, nil
}

func convertValidateToTag(v *validate.ValidateFieldOptions) (*structtag.Tags, error) {
	return structtag.Parse(v.GetRule())
}
