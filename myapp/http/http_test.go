package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/benbjohnson/myapp"
	"github.com/benbjohnson/myapp/mock"

	myhttp "github.com/benbjohnson/myapp/http"
)

func TestHandler(t *testing.T) {
	var us mock.UserService
	var h myhttp.Handler
	h.UserService = &us

	us.UserFn = func(id int) (*myapp.User, error) {
		if id != 100 {
			t.Fatalf("unexpected id: %d", id)
		}
		return &myapp.User{ID: 100, Name: "susy"}, nil
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/users/100", nil)
	h.ServeHTTP(w, r)

	if !us.UserInvoked {
		t.Fatalf("expected User() to be invoked")
	}
}
