package validator

import (
	"context"
	"reflect"
	"strings"

	"github.com/byvinesse/vinance-backend/pkg/errors"
	"github.com/go-playground/validator/v10"
)

var (
	validate    *validator.Validate
	tagRequired = "required"
)

func Init() {
	validate = validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// ValidateStruct validates a struct based on the tags on the fields.
// It will return the first error.
func ValidateStruct(ctx context.Context, s interface{}) error {
	err := validate.StructCtx(ctx, s)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}
		for _, err := range errs {
			if err.Tag() == tagRequired {
				return errors.ErrMissingField(err.Field())
			} else {
				return errors.ErrInvalidValue(err.Field())
			}
		}
	}

	return nil
}
