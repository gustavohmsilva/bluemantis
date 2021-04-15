package bluemantis

// Issue represents the basic structure that contains a client, the necessary
// data to create the issue in MantisBT and a request-response pair, that will
// hold the request sent/to be send to MantisBT and the response received/ to be
// received by the MantisBT application. It should not be manually created. It
// is preferred to be created with the NewIssue(), that receive the BaseIssue
// and the ExtendedIssue structs.
type Issue struct {
	Client *Client
	requestResponse
	BaseIssue
}

type BaseIssue struct {
	Summary     string
	Description string
	Category    string
	Project     Project
}

func (i *Issue) Send() bool {
	var err error
	i.response, err = i.Client.Do(i.request)
	if err != nil || i.response.StatusCode != 201 {
		return false
	}
	return true
}
