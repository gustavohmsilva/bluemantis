package bluemantis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/go-test/deep"
)

func TestNewClient(t *testing.T) {
	type args struct {
		url   string
		token string
	}

	mockServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth != SAMPLETOKEN {
				w.WriteHeader(http.StatusBadRequest)
			}
		}),
	)
	defer mockServer.Close()

	connectionTest := &Client{
		Client: &http.Client{},
		URL:    mockServer.URL,
		Token:  SAMPLETOKEN,
	}

	// Easy to miss, token is slightly off
	InvalidConnectionTest := &Client{
		Client: &http.Client{},
		URL:    mockServer.URL,
		Token:  "7-AvtZGHhpONO7shfeZXwKEX66WXuE9-",
	}

	invalidURL := &Client{
		Client: &http.Client{},
		URL:    "htxp://notavalid.domainname/",
		Token:  "7-EtgZGHhpONO7shfeZXxKEX66WXuE9-",
	}

	// Also slightly off, but because a rune is unnaceptable
	invalidToken := &Client{
		Client: &http.Client{},
		URL:    mockServer.URL,
		Token:  "7-EtgZGHhpONO7shfeZXxKEX66WXu!9-",
	}

	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		{
			"Connection Test",
			args{
				url:   connectionTest.URL,
				token: connectionTest.Token,
			},
			connectionTest,
			false,
		},
		{
			"Invalid Connection Test",
			args{
				url:   InvalidConnectionTest.URL,
				token: InvalidConnectionTest.Token,
			},
			nil,
			true,
		},
		{
			"Invalid URL Test",
			args{
				url:   invalidURL.URL,
				token: invalidURL.Token,
			},
			nil,
			true,
		},
		{
			"Invalid Token Test",
			args{
				url:   invalidToken.URL,
				token: invalidToken.Token,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.url, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"NewClient() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"NewClient() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestClient_NewIssue(t *testing.T) {
	type fields struct {
		Client *http.Client
		URL    string
		Token  string
	}
	type args struct {
		bascInfo *BaseIssue
	}

	validIssueClient := fields{
		Client: &http.Client{},
		URL:    "http://test.local/",
		Token:  SAMPLETOKEN,
	}

	baseIssue := &BaseIssue{
		Summary:     SAMPLESUMMARY,
		Description: SAMPLEDESCRIPTION,
		Category: &Category{
			Name: SAMPLECATEGORYNAME,
		},
		Project: &Project{
			Name: SAMPLEPROJECTNAME,
		},
	}

	validIssueBasicInfo := args{
		bascInfo: baseIssue,
	}
	baseJSON, err := json.Marshal(baseIssue)
	if err != nil {
		panic(err)
	}
	baseRequest, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s%s", validIssueClient.URL, newIssue),
		strings.NewReader(string(baseJSON)),
	)
	baseRequest.Header.Add("Content-Type", "application/json")
	baseRequest.Header.Add("Authorization", SAMPLETOKEN)

	MissingSummaryBasicInfo := args{
		bascInfo: &BaseIssue{
			Description: SAMPLEDESCRIPTION,
			Category: &Category{
				Name: SAMPLECATEGORYNAME,
			},
			Project: &Project{
				Name: SAMPLEPROJECTNAME,
			},
		},
	}

	MissingDescriptionBasicInfo := args{
		bascInfo: &BaseIssue{
			Summary: SAMPLESUMMARY,
			Category: &Category{
				Name: SAMPLECATEGORYNAME,
			},
			Project: &Project{
				Name: SAMPLEPROJECTNAME,
			},
		},
	}

	MissingCategoryBasicInfo := args{
		bascInfo: &BaseIssue{
			Summary:     SAMPLESUMMARY,
			Description: SAMPLEDESCRIPTION,
			Project: &Project{
				Name: SAMPLEPROJECTNAME,
			},
		},
	}

	MissingProjectBasicInfo := args{
		bascInfo: &BaseIssue{
			Summary:     SAMPLESUMMARY,
			Description: SAMPLEDESCRIPTION,
			Category: &Category{
				Name: SAMPLECATEGORYNAME,
			},
		},
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Issue
	}{
		{
			"Valid Issue test",
			validIssueClient,
			validIssueBasicInfo,
			&Issue{
				Client: &Client{
					Client: &http.Client{},
					URL:    validIssueClient.URL,
					Token:  SAMPLETOKEN,
				},
				requestResponse: requestResponse{
					request: baseRequest,
				},
				BaseIssue: &BaseIssue{
					Summary:     SAMPLESUMMARY,
					Description: SAMPLEDESCRIPTION,
					Category: &Category{
						Name: SAMPLECATEGORYNAME,
					},
					Project: &Project{
						Name: SAMPLEPROJECTNAME,
					},
				},
			},
		}, {
			"Missing Summary test",
			validIssueClient,
			MissingSummaryBasicInfo,
			nil,
		}, {
			"Missing Description test",
			validIssueClient,
			MissingDescriptionBasicInfo,
			nil,
		}, {
			"Missing Category test",
			validIssueClient,
			MissingCategoryBasicInfo,
			nil,
		}, {
			"Missing Project test",
			validIssueClient,
			MissingProjectBasicInfo,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Client: tt.fields.Client,
				URL:    tt.fields.URL,
				Token:  tt.fields.Token,
			}
			got := c.NewIssue(tt.args.bascInfo)
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf(
					"Client.NewIssue() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}
