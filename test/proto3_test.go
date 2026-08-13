package test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBsonTags(t *testing.T) {
	b := TestBson{}
	te := reflect.TypeOf(&b).Elem()

	id, ok := te.FieldByName("Id")
	assert.True(t, ok)
	idBsonTag := id.Tag.Get("bson")
	assert.Equal(t, "_id", idBsonTag)

	displayName, ok := te.FieldByName("DisplayName")
	assert.True(t, ok)
	displayNameBsonTag := displayName.Tag.Get("bson")
	assert.Equal(t, "display_namex", displayNameBsonTag)

	age, ok := te.FieldByName("Age")
	assert.True(t, ok)
	ageBsonTag := age.Tag.Get("bson")
	assert.Equal(t, "age,omitempty", ageBsonTag)

	kitchenSink, ok := te.FieldByName("KitchenSink")
	assert.True(t, ok)
	kitchenSinkBsonTag := kitchenSink.Tag.Get("bson")
	assert.Equal(t, "kitchen_sink,omitempty,inline,minsize,truncate", kitchenSinkBsonTag)

	skip, ok := te.FieldByName("Skip")
	assert.True(t, ok)
	skipBsonTag := skip.Tag.Get("bson")
	assert.Empty(t, skipBsonTag)

	email, ok := te.FieldByName("Email")
	assert.True(t, ok)
	emailBsonTag := email.Tag.Get("bson")
	assert.Equal(t, "email,omitempty", emailBsonTag)
}

func TestValidateTags(t *testing.T) {
	v := TestValidate{}
	te := reflect.TypeOf(&v).Elem()

	id, ok := te.FieldByName("Id")
	assert.True(t, ok)
	idValidateTag := id.Tag.Get("validate")
	assert.Equal(t, "required,omitempty,max=255", idValidateTag)

	email, ok := te.FieldByName("Email")
	assert.True(t, ok)
	emailValidateTag := email.Tag.Get("validate")
	assert.Equal(t, "required,hostname_rfc1123", emailValidateTag)

	skip, ok := te.FieldByName("Skip")
	assert.True(t, ok)
	skipValidateTag := skip.Tag.Get("validate")
	assert.Empty(t, skipValidateTag)
}
