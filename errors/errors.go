package errors

// ApplicationError struct
type ApplicationError struct {
	Status  uint32 `json:"status,omitempty"`
	Source  string `json:"source"`
	Message string `json:"message"`
}

func (e ApplicationError) Error() string {
	return e.Message
}
