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

func checkSuccessPutStatusCodes(resp httpResponse) error {
	if err := checkStatusCode(resp, 200, 201); err != nil {
		return err
	}
	return nil
}

func checkSuccessPostStatusCodes(resp httpResponse) error {
	if err := checkStatusCode(resp, 202); err != nil {
		return err
	}
	return nil
}

func checkSuccessDeleteStatusCodes(resp httpResponse) error {
	if err := checkStatusCode(resp, 202, 404); err != nil {
		return err
	}
	return nil
}
