package protago

import (
	"github.com/timmonfette1/protago/internal/genproto/protago/bson/v1"
	"github.com/timmonfette1/protago/internal/genproto/protago/validate/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type BsonOptions struct {
	FieldName string
	Omitempty bool
	Minsize   bool
	Truncate  bool
	Inline    bool
}

type ValidateOptions struct {
	Rule string
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

func openValidateOptions(vOpt *validate.ValidateFieldOptions) *ValidateOptions {
	if vOpt == nil {
		return nil
	}

	return &ValidateOptions{Rule: vOpt.GetRule()}
}

func GetAllBsonOptions(msg proto.Message) map[string]*BsonOptions {
	bOpts := make(map[string]*BsonOptions)
	m := msg.ProtoReflect()
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		bOpt := GetBsonOptions(fd)
		bOpts[string(fd.Name())] = bOpt
		return true
	})

	return bOpts
}

func GetBsonOptions(fd protoreflect.FieldDescriptor) *BsonOptions {
	opts := fd.Options().(*descriptorpb.FieldOptions)
	bOpts := proto.GetExtension(opts, bson.E_Options).(*bson.BsonFieldOptions)
	return openBsonFieldOptions(bOpts)
}

func GetAllValidateOptions(msg proto.Message) map[string]*ValidateOptions {
	vOpts := make(map[string]*ValidateOptions)
	m := msg.ProtoReflect()
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		vOpt := GetValidateOptions(fd)
		vOpts[string(fd.Name())] = vOpt
		return true
	})

	return vOpts
}

func GetValidateOptions(fd protoreflect.FieldDescriptor) *ValidateOptions {
	opts := fd.Options().(*descriptorpb.FieldOptions)
	vOpts := proto.GetExtension(opts, validate.E_Options).(*validate.ValidateFieldOptions)
	return openValidateOptions(vOpts)
}
