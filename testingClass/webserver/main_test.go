package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDoubleHandler(t *testing.T) {
	tt := []struct {
		name   string
		value  string
		double int
		err    string
	}{
		{
			name:  "missing number",
			value: "",
			err:   "missing value",
		},
		{
			name:  "not a number",
			value: "abc",
			err:   "not a number: abc",
		},
		{
			name:   "double of two",
			value:  "2",
			double: 4,
			err:    "",
		},
	}
	// curl localhost:8080/user?user_id=123
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "localhost:8080/double?v="+tc.value, nil)
			if err != nil {
				t.Fatalf("could not create a request %v", err)
			}
			rec := httptest.NewRecorder() // it gives w (Response Writer)
			doubleHandler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response %v", err)
			}

			if tc.err != "" {

				if res.StatusCode != http.StatusBadRequest {
					t.Errorf("exepected status Bad Request; got %v", res.StatusCode)
				}
				msg := string(bytes.TrimSpace(b))
				//errMsg := errors.New(msg)
				if tc.err != msg {
					t.Errorf("exepected message %q; got %q", tc.err, msg)
				}
				return
			}

			if res.StatusCode != http.StatusOK {
				t.Errorf("exepected status OK; got %v", res.StatusCode)
			}
			d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
			if err != nil {
				t.Fatalf("expected an int; got %s", b)
			}
			if d != tc.double {
				t.Errorf("expected double to be %v; got %v", tc.double, d)
			}

		})
	}

}
