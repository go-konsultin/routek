package routek

// Code represents API response codes
type Code string

// Common response codes
const (
	CodeOK            Code = "OK"
	CodeCreated       Code = "CREATED"
	CodeBadRequest    Code = "BAD_REQUEST"
	CodeUnauthorized  Code = "UNAUTHORIZED"
	CodeForbidden     Code = "FORBIDDEN"
	CodeNotFound      Code = "NOT_FOUND"
	CodeConflict      Code = "CONFLICT"
	CodeInternalError Code = "INTERNAL_ERROR"
)

// Response is the standard API response structure
type Response[T any] struct {
	Message   string `json:"message"`
	Code      Code   `json:"code"`
	Data      T      `json:"data"`
	Timestamp int64  `json:"timestamp"`
}
