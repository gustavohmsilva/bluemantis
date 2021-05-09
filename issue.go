package bluemantis

import (
	"errors"
	"reflect"
)

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
	*ExtendedIssue
}

// BaseIssue represent the minimal information required to submit a valid issue
// report to MantisBT using the package BlueMantis. Please notice that all of
// the data must be set, including the minimal data from Category and Project.
type BaseIssue struct {
	// Summary: We recommend putting your err.Error() here
	Summary string `valid:"required" json:"summary"`
	// Description: We recommend you put the object transited at the moment
	// of the error here in a serialized format (JSON)
	Description string   `valid:"required" json:"description"`
	Category    *Rel     `valid:"required" json:"category"`
	Project     *Project `valid:"required" json:"project"`
}

// ExtendedIssue represent the complete (not including history atm) information
// recommended to submit a valid issue report to MantisBT using the package
// BlueMantis. Please notice that all of this data is optional, and Mantis will
// normally assign default values for all of this relations. You should use name
// if you don't know a object given ID. Do not use label for this.
type ExtendedIssue struct {
	Status          *Status   `valid:"-" json:"status"`
	Reporter        *Reporter `valid:"-" json:"reporter"`
	Resolution      *Rel      `valid:"-" json:"resolution"`
	ViewState       *Rel      `valid:"-" json:"view_state"`
	Priority        *Rel      `valid:"-" json:"priority"`
	Severity        *Rel      `valid:"-" json:"severity"`
	Reproducibility *Rel      `valid:"-" json:"reproducibility"`
	Sticky          bool      `valid:"-" json:"sticky"`
	Meta
}

// Send do a immediate request to the MantisBT server using the information in
// Issue.Request and returns a confirmation if such request was successful or
// not. If anything other than 201 is returned, or if the connection fail for
// some reason, it will return an error.
func (i *Issue) Send() error {
	if reflect.DeepEqual(i, &Issue{}) {
		return errors.New("can't submit, invalid issue")
	}

	var err error
	i.response, err = i.Client.Do(i.request)
	if err != nil || i.response.StatusCode != 201 {
		return errors.New("can't submit, error contacting server")
	}
	return nil
}

// Retry do a second attempt of a immediate request to the MantisBT server using
// the same information used before. Best used together with Send, like:
// newIssue.Send().Retry().RetryLater()
func (i *Issue) Retry(err error) error {
	if err == nil {
		return nil
	}
	return i.Send()
}

// RetryLater will schedule
func (i *Issue) RetryLater(at string) {
	i.Client.Scheduler.AddFunc(i.Client.SchedulerInterval, func() {
		i.Send()
	})
}
