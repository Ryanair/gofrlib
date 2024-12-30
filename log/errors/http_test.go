package errors_test

import (
	"bitbucket.org/ryanair/gofrlib/log"
	gifrlibErrors "bitbucket.org/ryanair/gofrlib/log/errors"
	"github.com/pkg/errors"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHttpApiError_Error(t *testing.T) {
	err := errors.New("test error")
	httpErr := gifrlibErrors.NewHttpApiError(err, nil, nil)
	if httpErr.Error() != "test error" {
		t.Errorf("expected 'test error', got '%s'", httpErr.Error())
	}
}

func TestHttpApiError_Unwrap(t *testing.T) {
	err := errors.New("test error")
	httpErr := gifrlibErrors.NewHttpApiError(err, nil, nil)
	if !errors.Is(httpErr, err) {
		t.Errorf("expected true, got false")
	}
}

func TestHttpApiError_Is(t *testing.T) {
	err := errors.New("test error")
	httpErr := gifrlibErrors.NewHttpApiError(err, nil, nil)
	if !httpErr.Is(httpErr) {
		t.Errorf("expected true, got false")
	}
}

// doesn't assert anything because we have no method output, it's only to check if log format is valid
func TestHttpApiError_LogErrorWithMessage_OneHeaderHidden(t *testing.T) {
	log.Init(log.NewConfiguration("DEBUG", "TEST-APPLICATION", "TEST-PROJECT", "TEST-PROJECT-GROUP", "1.0.0", "testPrefix", "testEnv"))
	req := httptest.NewRequest("GET", "http://example.com", nil)
	req.Header.Set("Authorization", "Bearer abcd")
	req.Header.Set("x-api-key", "xcxcxcxc") // should be present
	res := httptest.NewRecorder().Result()
	err := errors.New("test error")
	httpErr := gifrlibErrors.NewHttpApiError(err, req, res)

	// Set environment variable for headers to hide
	os.Setenv("BLACK_LIST_HEADERS", "Authorization")
	defer os.Unsetenv("BLACK_LIST_HEADERS")

	httpErr.LogErrorWithMessage("test message", true, true)
}

// doesn't assert anything because we have no method output, it's only to check if log format is valid
func TestHttpApiError_LogErrorWithMessage_BothHeaderHidden(t *testing.T) {
	log.Init(log.NewConfiguration("DEBUG", "TEST-APPLICATION", "TEST-PROJECT", "TEST-PROJECT-GROUP", "1.0.0", "testPrefix", "testEnv"))
	req := httptest.NewRequest("GET", "http://example.com", nil)
	req.Header.Set("Authorization", "Bearer abcd")
	req.Header.Set("x-api-key", "xcxcxcxc") // should be present
	res := httptest.NewRecorder().Result()
	err := errors.New("test error")
	httpErr := gifrlibErrors.NewHttpApiError(err, req, res)

	// Set environment variable for headers to hide
	os.Setenv("BLACK_LIST_HEADERS", "Authorization,x-api-key")
	defer os.Unsetenv("BLACK_LIST_HEADERS")

	httpErr.LogErrorWithMessage("test message", true, true)
}

// doesn't assert anything because we have no method output, it's only to check if log format is valid
func TestHttpApiError_LogErrorWithMessage_UseDefaultSettings(t *testing.T) {
	log.Init(log.NewConfiguration("DEBUG", "TEST-APPLICATION", "TEST-PROJECT", "TEST-PROJECT-GROUP", "1.0.0", "testPrefix", "testEnv"))
	req := httptest.NewRequest("GET", "http://example.com", nil)
	req.Header.Set("Authorization", "Bearer abcd")
	req.Header.Set("x-api-key", "xcxcxcxc") // should be present
	res := httptest.NewRecorder().Result()
	err := errors.New("test error")
	httpErr := gifrlibErrors.NewHttpApiError(err, req, res)

	httpErr.LogErrorWithMessage("test message", true, true)
}

// doesn't assert anything because we have no method output, it's only to check if log format is valid
func TestLogError(t *testing.T) {
	log.Init(log.NewConfiguration("DEBUG", "TEST-APPLICATION", "TEST-PROJECT", "TEST-PROJECT-GROUP", "1.0.0", "testPrefix", "testEnv"))
	err := errors.New("test error")
	gifrlibErrors.LogError(err, true, true, "test message")
}
