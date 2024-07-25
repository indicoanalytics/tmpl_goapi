package errs

import "fmt"

type RequestError struct {
	LogErr error
	Err    error
}

func (requestError *RequestError) Error() string {
	if requestError.LogErr != nil {
		return fmt.Sprintf("%s: %s", requestError.LogErr, requestError.Err)
	}

	return requestError.Err.Error()
}

func (requestError *RequestError) Unwrap() error {
	return requestError.Err
}
