package httpclient_test

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestHowToConfigureHttpRequest(t *testing.T) {
	t.Skip("Skipping the test because is integrating with external services. Should be mocked")
	cases := []struct {
		name string
	}{
		{"case 1"},
	}

	// Configure the client
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		// TODO here configure the rest of transport
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   30 * time.Second,
		//CheckRedirect: redirectPolicyFunc,
		// TODO here configure the rest of client
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			// Create and configure the request
			req, err := http.NewRequest("GET", "http://google.com", nil)
			if err != nil {
				t.Fatalf("error creating new request: %v", err)
			}
			req.Header.Add("If-None-Match", `W/"wyzzy"`)
			// TODO here add more headers

			// Process the response
			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("error executing request: %v", err)
			}
			defer resp.Body.Close()

			// Here read the body. Test assertions should be performed
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			t.Logf("resp: %v", bodyBytes)
		})
	}
}
