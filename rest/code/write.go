package code

import (
	"encoding/json"
	"net/http"
)

// Write writes the code as defined in the constants along with its error string
func Write(w http.ResponseWriter, code ErrorCode) {
	WriteValueCode(w, code, codeToErrorString[code])
}

// WriteOk writes a successful reply to the ResponseWriter
func WriteOk(w http.ResponseWriter) {
	Write(w, Ok)
}

// WriteValue outputs the result as Ok(0) and body=value
func WriteValue(w http.ResponseWriter, value interface{}) {
	WriteValueCode(w, Ok, value)
}

// WriteError outputs the result as Error(1) and body=value
func WriteError(w http.ResponseWriter, value interface{}) {
	WriteValueCode(w, ServerErr, value)
}

// WriteValueCode outputs the result as code=code and body=value
func WriteValueCode(w http.ResponseWriter, code ErrorCode, value interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	v := struct {
		Code ErrorCode   `json:"code"`
		Body interface{} `json:"body"`
	}{
		Code: code,
		Body: value,
	}
	json.NewEncoder(w).Encode(v)
}
