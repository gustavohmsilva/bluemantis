package bluemantis

import "net/http"

type requestResponse struct {
	request  *http.Request
	response *http.Response
}
