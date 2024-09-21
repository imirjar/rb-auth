package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/imirjar/rb-auth/internal/models"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		want    *HTTPServer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogIn(t *testing.T) {
	type expected struct {
		token  string
		status uint
	}

	tests := []struct {
		name     string
		user     models.User
		expected expected
	}{
		{
			name: "ok",
			user: models.User{Login: "login", Password: "password"},
			expected: expected{
				status: http.StatusOK,
				token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjY1NzI1NjksIlVzZXJJRCI6MX0.GAUD2ulqg-UIXsomcc6B9vFD5Eqyrg75jwjH39o4BXg",
			},
		},
	}

	// create new server
	srv, err := New()
	if err != nil {
		log.Print(err)
	}

	//create fake service
	ctrl := gomock.NewController(t)
	mockService := NewMockService(ctrl)

	// mockService.EXPECT().Authenticate(gomock.Any(), models.User{Login: "NO", Password: "NO"}).Return(models.User{}, errors.New("not found"))
	first := mockService.EXPECT().
		BuildJWTString(gomock.Any(), models.User{Login: "Wrong", Password: "User"}).
		Return("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjY1NzI1NjksIlVzZXJJRCI6MX0.GAUD2ulqg-UIXsomcc6B9vFD5Eqyrg75jwjH39o4BXg", errors.New("bad")).
		MaxTimes(2)

	second := mockService.EXPECT().
		BuildJWTString(gomock.Any(), gomock.Any()).
		Return("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjY1NzI1NjksIlVzZXJJRCI6MX0.GAUD2ulqg-UIXsomcc6B9vFD5Eqyrg75jwjH39o4BXg", nil).
		MaxTimes(2)

	gomock.InOrder(
		first,
		second,
	)

	// connect fake service with http server
	srv.Service = mockService

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// user to buffer for httptest.NewRequest 3 param
			var buf bytes.Buffer
			err = json.NewEncoder(&buf).Encode(tt.user)
			if err != nil {
				log.Print(err)
			}

			req := httptest.NewRequest(http.MethodPost, "/auth", &buf)
			w := httptest.NewRecorder()

			handler := http.HandlerFunc(srv.LogIn())
			handler(w, req)
			// log.Println(w.Body)
			// t.Errorf("%s", w.Result().Status)

			if tt.expected.status != uint(w.Code) {
				t.Errorf("Status -> %v, want -> %v", uint(w.Code), tt.expected.status)
			}
		})
	}
}
