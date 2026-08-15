package validate

import (
	"github.com/timmonfette1/protago/internal/genproto/protago/validate/v1"
)

type ValidateOptions struct {
	Rule string
}

func openValidateOptions(vOpt *validate.ValidateFieldOptions) *ValidateOptions {
	if vOpt == nil {
		return nil
	}

	return &ValidateOptions{Rule: vOpt.GetRule()}
}
