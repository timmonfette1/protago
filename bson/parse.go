package bson

import (
	"github.com/timmonfette1/protago/internal/genproto/protago/bson/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

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
