package bson

import (
	"github.com/timmonfette1/protago/internal/genproto/protago/bson/v1"
)

type BsonOptions struct {
	FieldName string
	Omitempty bool
	Minsize   bool
	Truncate  bool
	Inline    bool
}

func openBsonFieldOptions(bOpt *bson.BsonFieldOptions) *BsonOptions {
	if bOpt == nil {
		return nil
	}

	return &BsonOptions{
		FieldName: bOpt.GetFieldName(),
		Omitempty: bOpt.GetOmitempty(),
		Minsize:   bOpt.GetMinsize(),
		Truncate:  bOpt.GetTruncate(),
		Inline:    bOpt.GetInline(),
	}
}
