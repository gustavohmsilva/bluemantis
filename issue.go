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
	var err error
	i.response, err = i.Client.Do(i.request)
	if err != nil || i.response.StatusCode != 201 {
		return false
	}
	return true
}
