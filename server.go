package bluemantis

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/asaskevich/govalidator"
)

type Server struct {
	URL   string
	Token string
}

func NewServer(url, token string) (*Server, error) {
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
	err = testServerConnection(&Server{URL: url, Token: token})
	if err != nil {
		return nil, err
	}
	return &Server{URL: url, Token: token}, nil
}

func testServerConnection(s *Server) error {
	client := &http.Client{}
	testRequest, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s%s", s.URL, getMyUserInfo),
		nil,
	)
	if err != nil {
		return errors.New("problem contacting MantisBT server")
	}
	testRequest.Header.Add("Authorization", s.Token)
	testResponse, err := client.Do(testRequest)
	if err != nil {
		return errors.New("problem contacting MantisBT server")

	}
	if testResponse.StatusCode != 200 {
		return errors.New("the token isn't valid for this server")
	}
	return nil
}
