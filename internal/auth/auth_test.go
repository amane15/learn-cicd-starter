package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestApiKey(t *testing.T) {
	tests := []struct {
		key       string
		value     string
		expect    string
		expectErr string
	}{
		{expectErr: "no authorization header"},
		{key: "Authorization", expectErr: "no authorization header"},
		{key: "Authorization", value: "-", expectErr: "malformed authorization header"},
		{key: "Authorization", value: "Bearer xxxxx", expect: "malformed authorization header"},
		{key: "Authorization", value: "ApiKey xxxxx", expect: "xxxxx", expectErr: "not expecting error"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetApiKey Case #%v:", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(test.key, test.value)

			output, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
				return
			}

			if output != test.expect {
				t.Errorf("Unexpected: TestGetAPIKey:%s", output)
				return
			}
		})
	}
}
