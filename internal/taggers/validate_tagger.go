package taggers

import (
	"fmt"

	"github.com/fatih/structtag"
	"github.com/timmonfette1/protago/internal/genproto/validate"
	pgs "github.com/timmonfette1/protoc-gen-star/v2"
	pgsgo "github.com/timmonfette1/protoc-gen-star/v2/lang/go"
)

type ValidateTagger struct {
	pgs.DebuggerCommon
	pgsgo.Context
}

var _ Tagger = (*ValidateTagger)(nil)

func (vt *ValidateTagger) GetName() string {
	return "ValidateTagger"
}

func (vt *ValidateTagger) GenerateTag(field pgs.Field) (*structtag.Tag, error) {
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
