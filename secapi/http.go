package secapi

import "fmt"

type httpResponse interface {
	StatusCode() int
}

func checkStatusCode(resp httpResponse, alloweds ...int) error {
	code := resp.StatusCode()
	for _, allowed := range alloweds {
		if code == allowed {
			return nil
		}
	}
	return fmt.Errorf("unexpected status code %d", code)
}
