package auth

import (
	"testing"
	"net/http"
)

func TestGetAPIKey(t *testing.T) {
	hdrEmptyAuth := http.Header{}
	hdrEmptyAuth.Add("Authorization", "")
	hdrMalformed := http.Header{}
	hdrMalformed.Add("Authorization", "Bearer_malformed")
	hdrAuth := http.Header{}
	hdrAuth.Add("Authorization", "ApiKey token")
	// Fail test
	// hdrAuth.Add("Authorization", "AKey token")

	tests := []struct {
		headers http.Header
		wantError bool
		wanted string
	}{
		{ headers: http.Header{}, wantError: true, },
		{ headers: hdrEmptyAuth, wantError: true, },
		{ headers: hdrMalformed, wantError: true, },
		{ headers: hdrAuth, wantError: false, wanted: "token"},
	}


	for _, test := range tests {
		result, error := GetAPIKey(test.headers)
		if error != nil && !test.wantError {
			t.Fatalf("Did not want Error and got it")
		}
		if error == nil && test.wantError {
			t.Fatalf("Wanted Error and did not get it")
		}

		if error == nil && !test.wantError {
			if test.wanted != result {
				t.Fatalf("Mismatch wanted: %s got: %s", test.wanted, result)
			}
		}

	}

}


