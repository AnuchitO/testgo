package http

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestMakeHTTP(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": 1, "name": "kyle", "description": "novice gopher"}`))
	}))
	want := &Response{
		ID:          1,
		Name:        "kyle",
		Description: "novice gopher",
	}

	t.Run("Happy server response", func(t *testing.T) {
		defer server.Close()
		fmt.Println(server.URL)
		time.Sleep(10 * time.Second)
		resp, err := MakeHTTPCall(server.URL)

		if !reflect.DeepEqual(resp, want) {
			t.Errorf("expected (%v), got (%v)", want, resp)
		}

		if !errors.Is(err, nil) {
			t.Errorf("expected (%v), got (%v)", nil, err)
		}
	})
}
