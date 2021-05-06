

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestIndex(t *testing.T) {
	tests := []struct {
		path           string
		wantStatusCode int
		wantBody       string
	}{
		{ path: "/go-auth/",	wantStatusCode: http.StatusOK,			wantBody: "No Cloud IAP header found.\n"	},
		{ path: "/hello",		wantStatusCode: http.StatusNotFound,	wantBody: "404 page not found\n"			},
	}

	for _, test := range tests {
		req	:= httptest.NewRequest("GET", test.path, nil)
		rr	:= httptest.NewRecorder()
		a	:= &app{}											// Do not use newApp since it uses the metadata server

		a.index(rr, req)

		if got := rr.Result().StatusCode;	got != test.wantStatusCode	{ t.Errorf("index(%s) got status code %d, want %d", test.path, got, test.wantStatusCode) }
		if got := rr.Body.String();			got != test.wantBody		{ t.Errorf("index(%s) got %q, want %q", test.path, got, test.wantBody) }
	}
}
