package bluemantis

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		url   string
		token string
	}

	mockServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth != "7-EtgZGHhpONO7shfeZXxKEX66WXuE9-" {
				w.WriteHeader(http.StatusBadRequest)
			}
		}),
	)
	defer mockServer.Close()

	connectionTest := &Client{
		Client: &http.Client{},
		URL:    mockServer.URL,
		Token:  "7-EtgZGHhpONO7shfeZXxKEX66WXuE9-",
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
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
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
		Token:  "SAMPLETOKEN",
	}

	baseRequest, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s%s", validIssueClient.URL, newIssue),
		nil,
	)
	baseRequest.Header.Add("Content-Type", "application/json")
	baseRequest.Header.Add("Authorization", validIssueClient.Token)

	validIssueBasicInfo := args{
		bascInfo: &BaseIssue{
			Summary:     "This is the summary",
			Description: "This is the description",
			Category: &Category{
				Name: "This is the Category Name",
			},
			Project: &Project{
				Name: "This is the Project Name",
			},
		},
	}

	MissingSummaryBasicInfo := args{
		bascInfo: &BaseIssue{
			Description: "This is the description",
			Category: &Category{
				Name: "This is the Category Name",
			},
			Project: &Project{
				Name: "This is the Project Name",
			},
		},
	}

	MissingDescriptionBasicInfo := args{
		bascInfo: &BaseIssue{
			Summary: "This is the summary",
			Category: &Category{
				Name: "This is the Category Name",
			},
			Project: &Project{
				Name: "This is the Project Name",
			},
		},
	}

	MissingCategoryBasicInfo := args{
		bascInfo: &BaseIssue{
			Summary:     "This is the summary",
			Description: "This is the description",
			Project: &Project{
				Name: "This is the Project Name",
			},
		},
	}

	MissingProjectBasicInfo := args{
		bascInfo: &BaseIssue{
			Summary:     "This is the summary",
			Description: "This is the description",
			Category: &Category{
				Name: "This is the Category Name",
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
					Token:  validIssueClient.Token,
				},
				requestResponse: requestResponse{
					request: baseRequest,
				},
				BaseIssue: &BaseIssue{
					Summary:     "This is the summary",
					Description: "This is the description",
					Category: &Category{
						Name: "This is the Category Name",
					},
					Project: &Project{
						Name: "This is the Project Name",
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
			if got := c.NewIssue(tt.args.bascInfo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.NewIssue() = %v, want %v", got, tt.want)
			}
		})
	}
}
