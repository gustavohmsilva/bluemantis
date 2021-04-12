package bluemantis

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
	res, err := i.Client.Do(i.request)
	if err != nil || res.StatusCode != 201 {
		return false
	}
	return true
}
