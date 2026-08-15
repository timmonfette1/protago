package protago

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timmonfette1/protago/internal/testdata"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestGetAllBsonOptions(t *testing.T) {
	ks := "cool sink"
	a := int32(30)
	s := true
	input := &testdata.TestBson{
		Id:          "id",
		Age:         &a,
		Skip:        &s,
		KitchenSink: &ks,
	}

	bsonOpts := GetAllBsonOptions(input)
	assert.Len(t, bsonOpts, 4)

	type args struct {
		name  string
		isNil bool
		opts  *BsonOptions
	}
	expected := []args{
		{
			name: "id",
			opts: &BsonOptions{
				FieldName: "_id",
			},
		},
		{
			name: "age",
			opts: &BsonOptions{
				FieldName: "age",
				Omitempty: true,
			},
		},
		{
			name:  "skip",
			isNil: true,
			opts:  nil,
		},
		{
			name: "kitchen_sink",
			opts: &BsonOptions{
				FieldName: "kitchen_sink",
				Omitempty: true,
				Minsize:   true,
				Truncate:  true,
				Inline:    true,
			},
		},
	}

	for _, arg := range expected {
		t.Run(arg.name, func(t *testing.T) {
			b, ok := bsonOpts[arg.name]
			assert.True(t, ok)

			if arg.isNil {
				assert.Nil(t, b)
			} else {
				assert.NotNil(t, b)
				assert.Equal(t, *arg.opts, *b)
			}
		})
	}
}

func TestGetBsonOptions(t *testing.T) {
	ks := "cool sink"
	input := &testdata.TestBson{
		Id:          "id",
		KitchenSink: &ks,
	}

	var kitchenSinkField protoreflect.FieldDescriptor
	m := input.ProtoReflect()
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if fd.Name() == "kitchen_sink" {
			kitchenSinkField = fd
			return false
		}
		return true
	})
	assert.NotNil(t, kitchenSinkField)

	bsonOpts := GetBsonOptions(kitchenSinkField)
	assert.NotNil(t, bsonOpts)
	assert.Equal(t, "kitchen_sink", bsonOpts.FieldName)
	assert.True(t, bsonOpts.Omitempty)
	assert.True(t, bsonOpts.Minsize)
	assert.True(t, bsonOpts.Truncate)
	assert.True(t, bsonOpts.Inline)
}

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
