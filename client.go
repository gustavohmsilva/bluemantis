package bluemantis

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/asaskevich/govalidator"
)

// Client represents the basic structure that contains the configuration needed
// for a valid communication with any given Mantis Bug Tracker instance.
// It should not be created directly, please make use of the function NewClient
// to create a valid and most importantly, tested, Client.
type Client struct {
	*http.Client
	URL   string
	Token string
}

// NewClient creates and test (including the connection) with a MantisBT server
// and returns a validated pointer to a Client struct for usage in your
// application. In case of a invalid URL, invalid Token, or communication
// failure with the MantisBT desired installation, it will return nil and a
// descriptive error with the problem.
func NewClient(url, token string) (*Client, error) {
	if !govalidator.IsURL(url) {
		return nil, errors.New("invalid URL")
	}

	isValidToken, err := regexp.MatchString(`^[a-zzA-Z0-9\-\_]{32}$`, token)
	if err != nil {
		return nil, errors.New(
			"an error happened when trying to validate token",
		)
	}
	if !isValidToken {
		return nil, errors.New("token provided is invalid")
	}

	newClient := &Client{
		Client: &http.Client{},
		URL:    url,
		Token:  token,
	}
	err = testServerConnection(newClient)
	if err != nil {
		return nil, err
	}

	return newClient, nil
}

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

// NewIssue will currently still being implemented, at the moment it creates
// only a basic pointer to a Issue struct, without any of the actual necessary
// data <- TODO
func (c *Client) NewIssue() *Issue {
	issue := new(Issue)
	issue.Client = c
	var err error
	issue.requestResponse.request, err = http.NewRequest(
		"POST",
		fmt.Sprintf("%s%s", c.URL, newIssue),
		nil,
	)
	if err != nil {
		log.Print(err.Error())
		return &Issue{}
	}
	issue.requestResponse.request.Header.Add(
		"Authorization",
		c.Token,
	)
	issue.requestResponse.request.Header.Add(
		"Content-Type",
		"application/json",
	)
	return issue
}
