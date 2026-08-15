package generators

import (
	"fmt"

	"github.com/fatih/structtag"
	"github.com/timmonfette1/protago/internal/genproto/protago/validate/v1"
	pgs "github.com/timmonfette1/protoc-gen-star/v2"
	pgsgo "github.com/timmonfette1/protoc-gen-star/v2/lang/go"
)

type ValidateGenerator struct {
	pgs.DebuggerCommon
	pgsgo.Context
}

var _ Generator = (*ValidateGenerator)(nil)

func (vt *ValidateGenerator) GetName() string {
	return "ValidateTagger"
}

func (vt *ValidateGenerator) GenerateTag(field pgs.Field) (*structtag.Tag, error) {
	vt.Debug(fmt.Sprintf("parsing 'validate' tags for field '%s'", field.Name().String()))

	var validateFieldOptions *validate.ValidateFieldOptions
	_, err := field.Extension(validate.E_Options, &validateFieldOptions)
	if err != nil {
		return nil, err
	}

	if validateFieldOptions == nil {
		return nil, nil
	}

	tags, err := structtag.Parse(validateFieldOptions.GetRule())
	if err != nil {
		return nil, err
	}
	if tags == nil {
		return nil, fmt.Errorf("unable to parse 'validate' rule into a valid struct tag")
	}
	if tags.Len() != 1 {
		return nil, fmt.Errorf("invalid 'validate' rule. Expected only a single 'validate' tag, found %d tags instead", tags.Len())
	}

	return tags.Get("validate")
}
