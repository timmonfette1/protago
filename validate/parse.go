package validate

import (
	"github.com/timmonfette1/protago/internal/genproto/protago/validate/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

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
