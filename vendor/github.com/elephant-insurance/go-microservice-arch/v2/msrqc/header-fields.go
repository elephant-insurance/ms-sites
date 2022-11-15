package msrqc

import (
	"time"

	grq "github.com/parnurzeal/gorequest"
)

const defaultRequestTimeoutMS = 30000 * time.Millisecond

// PostRequestWithHeaders returns a POST request with pre-filled headers
func PostRequestWithHeaders(c Context, url string) *grq.SuperAgent {
	rtn := grq.New().
		Post(url).
		Timeout(defaultRequestTimeoutMS)

	return setHeaders(c, rtn)
}

// GetRequestWithHeaders returns a GET request with pre-filled headers
func GetRequestWithHeaders(c Context, url string) *grq.SuperAgent {
	rtn := grq.New().
		Get(url).
		Timeout(defaultRequestTimeoutMS)

	return setHeaders(c, rtn)
}

// PutRequestWithHeaders returns a PUT request with pre-filled headers
func PutRequestWithHeaders(c Context, url string) *grq.SuperAgent {
	rtn := grq.New().
		Put(url).
		Timeout(defaultRequestTimeoutMS)

	return setHeaders(c, rtn)
}

// PatchRequestWithHeaders returns a PATCH request with pre-filled headers
func PatchRequestWithHeaders(c Context, url string) *grq.SuperAgent {
	rtn := grq.New().
		Patch(url).
		Timeout(defaultRequestTimeoutMS)

	return setHeaders(c, rtn)
}

// DeleteRequestWithHeaders returns a DELETE request with pre-filled headers
func DeleteRequestWithHeaders(c Context, url string) *grq.SuperAgent {
	rtn := grq.New().
		Delete(url).
		Timeout(defaultRequestTimeoutMS)

	return setHeaders(c, rtn)
}

// OptionsRequestWithHeaders returns an OPTIONS request with pre-filled headers
func OptionsRequestWithHeaders(c Context, url string) *grq.SuperAgent {
	rtn := grq.New().
		Options(url).
		Timeout(defaultRequestTimeoutMS)

	return setHeaders(c, rtn)
}

func setHeaders(c Context, sa *grq.SuperAgent) *grq.SuperAgent {
	if c != nil {
		hmap := HeadersFromContext(c)

		for k, v := range hmap {
			if len(v) > 0 && v[0] != `` {
				sa.Set(k, v[0])
			}
		}
	}

	return sa
}
