package bluemantis

import (
	"errors"
	"fmt"
	"net/http"
)

func testServerConnection(c *Client) error {
	testRequest, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s%s", c.URL, getMyUserInfo),
		nil,
	)
	if err != nil {
		return errors.New("problem contacting MantisBT server")
	}
	testRequest.Header.Add("Authorization", c.Token)
	testResponse, err := c.Do(testRequest)
	if err != nil {
		return errors.New("problem contacting MantisBT server")

	}
	if testResponse.StatusCode != 200 {
		return errors.New("the token isn't valid for this server")
	}
	return nil
}
