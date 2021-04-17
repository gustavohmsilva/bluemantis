package bluemantis

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIssue_Send(t *testing.T) {
	type fields struct {
		Client          *Client
		requestResponse requestResponse
		BaseIssue       *BaseIssue
	}

	mockServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.RequestURI {
			case newIssue:
				auth := r.Header.Get("Authorization")
				if auth != SAMPLETOKEN {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				if r.Method != "POST" {
					w.WriteHeader(
						http.StatusMethodNotAllowed,
					)
					return
				}
				if r.RequestURI != newIssue {
					w.WriteHeader(http.StatusNotFound)
					return
				}
				ct := r.Header.Get("Content-Type")
				if ct != "application/json" {
					w.WriteHeader(
						http.StatusUnsupportedMediaType,
					)
					return
				}
				w.WriteHeader(http.StatusCreated)
			default:
				auth := r.Header.Get("Authorization")
				if auth != "7-EtgZGHhpONO7shfeZXxKEX66WXuE9-" {
					w.WriteHeader(http.StatusBadRequest)
				}
			}
		}),
	)

	c, _ := NewClient(mockServer.URL, SAMPLETOKEN)
	i := c.NewIssue(&BaseIssue{
		Summary:     SAMPLESUMMARY,
		Description: SAMPLEDESCRIPTION,
		Category:    &Category{Name: SAMPLECATEGORYNAME},
		Project:     &Project{Name: SAMPLEPROJECTNAME},
	})

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"Invalid Issue Test",
			fields{},
			true,
		}, {
			"Valid Issue Test",
			fields{
				BaseIssue:       i.BaseIssue,
				requestResponse: i.requestResponse,
				Client:          i.Client,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Issue{
				Client:          tt.fields.Client,
				requestResponse: tt.fields.requestResponse,
				BaseIssue:       tt.fields.BaseIssue,
			}
			if err := i.Send(); (err != nil) != tt.wantErr {
				t.Errorf(
					"Issue.Send() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}
