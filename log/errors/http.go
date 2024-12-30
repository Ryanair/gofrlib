package errors

import (
	"bitbucket.org/ryanair/gofrlib/log"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

var defaultBlackListHeaders = []string{"x-api-key", "Authorization"}

type HttpApiError struct {
	error
	req *http.Request
	res *http.Response
}

func (e *HttpApiError) Error() string {
	return e.error.Error()
}

func (e *HttpApiError) Unwrap() error {
	return e.error
}

func (e *HttpApiError) Is(tgt error) bool {
	_, ok := tgt.(*HttpApiError)
	return ok
}

func NewHttpApiError(error error, req *http.Request, res *http.Response) *HttpApiError {
	return &HttpApiError{error: error, req: req, res: res}
}

func (e *HttpApiError) LogErrorWithMessage(msg string, withRequestBody, withResponseBody bool) {
	for _, h := range getBlackListHeaders() {
		e.req.Header.Del(h)
	}
	dumpedRequest, _ := httputil.DumpRequest(e.req, withRequestBody)
	dumpedResponse := make([]byte, 0)
	if e.res != nil {
		dumpedResponse, _ = httputil.DumpResponse(e.res, withResponseBody)
	}

	log.ErrorW(msg,
		log.ErrorKey, e.Error(),
		log.RequestDumpKey, string(dumpedRequest),
		log.ResponseDumpKey, string(dumpedResponse))
}

func getBlackListHeaders() []string {
	if blackListHeadersEnv, exists := os.LookupEnv("BLACK_LIST_HEADERS"); exists {
		return strings.Split(blackListHeadersEnv, ",")
	}
	return defaultBlackListHeaders
}

// LogError this is a generic method which is handling various types of errors and logs what's the most important
func LogError(err error, withRequestBody, withResponseBody bool, msg string, args ...any) {
	var httpApiError *HttpApiError
	if errors.As(err, &httpApiError) {
		httpApiError.LogErrorWithMessage(fmt.Sprintf(msg, args...), withRequestBody, withResponseBody)
	} else {
		log.ErrorW(fmt.Sprintf(msg, args...), log.ErrorKey, err)
	}
}
