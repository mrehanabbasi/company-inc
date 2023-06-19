package domainErr

// const will be used for the status of api handlers.
const (
	NotFound            = "NOT_FOUND"
	BadRequest          = "BAD_REQUEST"
	InternalServerError = "INTERNAL_SERVER_ERROR"
)

// APIError struct contains the code and message of error.
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type APIErrors struct {
	Code    string   `json:"code"`
	Message []string `json:"message"`
}

func (a *APIError) Error() string {
	return a.Message
}

func (a *APIErrors) Error() []string {
	return a.Message
}

// IsError will return whether error exists or not.
func (a *APIError) IsError(errType string) bool {
	return a.Code == errType
}

// NewAPIError returns the error type and error message.
func NewAPIError(errType string, message string) *APIError {
	return &APIError{
		Code:    errType,
		Message: message,
	}
}

func NewAPIErrors(errType string, message []string) *APIErrors {
	return &APIErrors{
		Code:    errType,
		Message: message,
	}
}
