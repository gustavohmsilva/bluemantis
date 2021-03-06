package bluemantis

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/robfig/cron"
)

// Client represents the basic structure that contains the configuration needed
// for a valid communication with any given Mantis Bug Tracker instance.
// It should not be created directly, please make use of the function NewClient
// to create a valid and most importantly, tested, Client.
type Client struct {
	*http.Client
	URL               string
	Token             string
	Scheduler         *cron.Cron
	SchedulerInterval string
}

// NewClient creates and test (including the connection) with a MantisBT server
// and returns a validated pointer to a Client struct for usage in your
// application. In case of a invalid URL, invalid Token, or communication
// failure with the MantisBT desired installation, it will return nil and a
// descriptive error with the problem.
func NewClient(url, token string, delayInterval string) (*Client, error) {
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
		Client:            &http.Client{},
		URL:               url,
		Token:             token,
		Scheduler:         cron.New(),
		SchedulerInterval: delayInterval,
	}
	err = testServerConnection(newClient)
	if err != nil {
		return nil, err
	}
	newClient.Scheduler.Start()
	return newClient, nil
}

// NewIssue will currently still being implemented, at the moment it creates
// only a basic pointer to a Issue struct, without any of the actual necessary
// data <- TODO
func (c *Client) NewIssue(bascInfo *BaseIssue) *Issue {
	_, err := govalidator.ValidateStruct(bascInfo)
	if err != nil {
		return nil
	}

	issue := new(Issue)
	issue.Client = c
	issue.BaseIssue = bascInfo

	body, err := json.Marshal(issue.BaseIssue)
	if err != nil {
		return nil
	}
	issue.request, _ = http.NewRequest(
		"POST",
		fmt.Sprintf("%s%s", c.URL, newIssue),
		strings.NewReader(string(body)),
	)
	issue.requestResponse.request.Header.Add(
		"Content-Type",
		"application/json",
	)
	issue.requestResponse.request.Header.Add("Authorization", c.Token)

	return issue
}
