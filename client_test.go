package bluemantis

import (
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
