package builders

import "github.com/go-playground/validator/v10"

func validateRequired(validate *validator.Validate, fields ...any) error {
	for _, f := range fields {
		if err := validate.Var(f, "required"); err != nil {
			return err
		}
	}
	return nil
}

// TODO Find a better name for this function
func validateOneRequired(validate *validator.Validate, fields ...any) error {
	errors := []error{}
	for _, f := range fields {
		if err := validate.Var(f, "required"); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) == len(fields) {
		return errors[0]
	}
	return nil
}
