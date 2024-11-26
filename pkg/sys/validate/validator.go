package validate

import "context"

// Condition checks condition in function
type Condition func(ctx context.Context) error

// Validate checks error from conditions is validation error
func Validate(ctx context.Context, conds ...Condition) error {
	ve := NewValidationErrors()

	for _, c := range conds {
		err := c(ctx)
		if err != nil {
			if IsValidationError(err) {
				ve.addError(err.Error())
				continue
			}

			return err
		}
	}

	if ve.Messages == nil {
		return nil
	}

	return ve
}
