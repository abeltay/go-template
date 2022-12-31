package code

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	tests := []struct {
		name     string
		input    ErrorCode
		expected string
	}{
		{
			name:     "Ok",
			input:    Ok,
			expected: fmt.Sprintf(`{"code":%d,"body":""}`, Ok),
		},
		{
			name:     "Server error code",
			input:    ServerErr,
			expected: fmt.Sprintf(`{"code":%d,"body":"Server error"}`, ServerErr),
		},
	}
	for _, tt := range tests {
		w := httptest.NewRecorder()
		Write(w, tt.input)
		assert.Equal(t, tt.expected, strings.TrimSpace(w.Body.String()), tt.name)
	}
}

func TestWriteOk(t *testing.T) {
	expected := fmt.Sprintf(`{"code":%d,"body":""}`, Ok)

	w := httptest.NewRecorder()
	WriteOk(w)
	assert.Equal(t, expected, strings.TrimSpace(w.Body.String()))
}

func TestWriteValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: fmt.Sprintf(`{"code":%d,"body":""}`, Ok),
		},
		{
			name:     "With message",
			input:    "the error message",
			expected: fmt.Sprintf(`{"code":%d,"body":"the error message"}`, Ok),
		},
	}
	for _, tt := range tests {
		w := httptest.NewRecorder()
		WriteValue(w, tt.input)
		assert.Equal(t, tt.expected, strings.TrimSpace(w.Body.String()), tt.name)
	}
}

func TestWriteError(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: fmt.Sprintf(`{"code":%d,"body":""}`, ServerErr),
		},
		{
			name:     "With message",
			input:    "the error message",
			expected: fmt.Sprintf(`{"code":%d,"body":"the error message"}`, ServerErr),
		},
	}
	for _, tt := range tests {
		w := httptest.NewRecorder()
		WriteError(w, tt.input)
		assert.Equal(t, tt.expected, strings.TrimSpace(w.Body.String()), tt.name)
	}
}

func TestWriteValueCode(t *testing.T) {
	tests := []struct {
		name     string
		code     ErrorCode
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			code:     0,
			input:    "",
			expected: fmt.Sprintf(`{"code":%d,"body":""}`, Ok),
		},
		{
			name:     "With message",
			code:     1,
			input:    "the error message",
			expected: fmt.Sprintf(`{"code":%d,"body":"the error message"}`, ServerErr),
		},
	}
	for _, tt := range tests {
		w := httptest.NewRecorder()
		WriteValueCode(w, tt.code, tt.input)
		assert.Equal(t, tt.expected, strings.TrimSpace(w.Body.String()), tt.name)
	}
}
