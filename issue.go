package bluemantis

import "errors"

// Issue represents the basic structure that contains a client, the necessary
// data to create the issue in MantisBT and a request-response pair, that will
// hold the request sent/to be send to MantisBT and the response received/ to be
// received by the MantisBT application. It should not be manually created. It
// is preferred to be created with the NewIssue(), that receive the BaseIssue
// and the ExtendedIssue structs.
type Issue struct {
	Client *Client
	requestResponse
	*BaseIssue
}

// Category represents a specific type of Rel (relation). This should probably
// be deprecated soon as the package increase size, because the relations will
// differ a little bit more than just ID, Name and Label.
type Category Rel

// BaseIssue represent the minimal information required to submit a valid issue
// report to MantisBT using the package BlueMantis. Please notice that all of
// the data must be set, including the minimal data from Category and Project.
type BaseIssue struct {
	Summary     string    `valid:"required"`
	Description string    `valid:"required"`
	Category    *Category `valid:"required"`
	Project     *Project  `valid:"required"`
}

// Send do a immediate request to the MantisBT server using the information in
// Issue.Request and returns a confirmation if such request was successful or
// not. If anything other than 201 is returned, or if the connection fail for
// some reason, it will return an error.
func (i *Issue) Send() error {
	if i == nil {
		return errors.New("can't to submit, invalid issue")
	}
	var err error
	i.response, err = i.Client.Do(i.request)
	if err != nil || i.response.StatusCode != 201 {
		return errors.New("can't to submit, error contacting server")
	}
	return nil
}
