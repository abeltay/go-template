package code

// ErrorCode is used to indicate server responses
type ErrorCode int

// Constants representing different code response types
const (
	Ok ErrorCode = iota
	ServerErr
	ParamErr
	ReqBodyErr
)

// codeToErrorString maps a code to its error string, if the code is for an error.
// The assumption is that for a non error ErrorCode, the error string should be empty.
var codeToErrorString = map[ErrorCode]string{
	ServerErr:  "Server error",
	ParamErr:   "Error on HTTP params",
	ReqBodyErr: "Error reading body",
}
