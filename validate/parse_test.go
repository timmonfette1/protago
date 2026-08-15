package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timmonfette1/protago/internal/testdata"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestGetAllValidateOptions(t *testing.T) {
	e := "mock@unittest.com"
	s := true
	input := &testdata.TestValidate{
		Id:    "id",
		Email: &e,
		Skip:  &s,
	}

	validateOpts := GetAllValidateOptions(input)
	assert.Len(t, validateOpts, 3)

	type args struct {
		name  string
		isNil bool
		opts  *ValidateOptions
	}
	expected := []args{
		{
			name: "id",
			opts: &ValidateOptions{
				Rule: "validate:\"required,omitempty,max=255\"",
			},
		},
		{
			name: "email",
			opts: &ValidateOptions{
				Rule: "validate:\"required,hostname_rfc1123\"",
			},
		},
		{
			name:  "skip",
			isNil: true,
			opts:  nil,
		},
	}

	for _, arg := range expected {
		t.Run(arg.name, func(t *testing.T) {
			v, ok := validateOpts[arg.name]
			assert.True(t, ok)

			if arg.isNil {
				assert.Nil(t, v)
			} else {
				assert.NotNil(t, v)
				assert.Equal(t, *arg.opts, *v)
			}
		})
	}
}

func TestGetValidateOptions(t *testing.T) {
	input := &testdata.TestValidate{
		Id: "id",
	}

	var idField protoreflect.FieldDescriptor
	m := input.ProtoReflect()
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if fd.Name() == "id" {
			idField = fd
			return false
		}
		return true
	})
	assert.NotNil(t, idField)

	validateOpts := GetValidateOptions(idField)
	assert.NotNil(t, validateOpts)
	assert.Equal(t, "validate:\"required,omitempty,max=255\"", validateOpts.Rule)
}
