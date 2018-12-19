package utils

type (
	// HTTPError custom http error to handle custom errors
	HTTPError struct {
		error      `json:"-"`
		StatusCode int                 `json:"-"`
		Message    string              `json:"message"`
		Errors     map[string][]string `json:"errors,omitempty"`
	}
)

// GetNotFoundError returns not error associated with HTTP NotFound error
func GetNotFoundError(message string) error {
	if message == "" {
		message = "Not Found Error"
	}
	err := HTTPError{}
	err.StatusCode = 404
	err.Message = message
	return err
}

// GetUnAuthorizedError returns not error associated with HTTP Unauthorized error
func GetUnAuthorizedError(message string) error {
	if message == "" {
		message = "You do not have access to this section"
	}
	err := HTTPError{}
	err.StatusCode = 401
	err.Message = message
	return err
}

// GetForbiddenError returns not error associated with HTTP Forbidden error
func GetForbiddenError(message string) error {
	if message == "" {
		message = "You do not have access to this section"
	}
	err := HTTPError{}
	err.StatusCode = 403
	err.Message = message
	return err
}

// GetValidationError returns error associated with HTTP Validation error
func GetValidationError(message string, errors ...string) error {
	if message == "" {
		message = "The submitted information is not valid"
	}
	err := HTTPError{}
	err.StatusCode = 400
	err.Message = message
	if len(errors) > 0 {
		err.Errors = make(map[string][]string)
		for i := 0; i < len(errors); i += 2 {
			err.Errors[errors[i]] = []string{errors[i+1]}
		}
	}
	return err
}
