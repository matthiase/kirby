package errors

// AuthenticationError struct
type AuthenticationError struct {
	Message string
}

func (e AuthenticationError) Error() string {
	return e.Message
}

// ValidationError struct
type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

// NotFoundError struct
type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "Record not found"
}
