package notify

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMessageBody(t *testing.T) {
	type msg struct {
		Username string
		Text     string
	}
	m := &msg{}
	b := messageBody("yo", "yodawg")
	if err := json.Unmarshal([]byte(b), &m); err != nil {
		t.Errorf("expected message body to be valid JSON: %s", err)
	}
	if m.Username != "yo" || m.Text != "yodawg" {
		t.Error("expected JSON to be filled out with provided arguments")
	}
}

func TestSlack(t *testing.T) {
	if err := Slack("", "bot", "yo"); err != nil {
		t.Errorf("expected nil when no slack URL is passed but got %s", err)
	}

	cases := []struct {
		ok bool
	}{
		{ok: true},
		{ok: false},
	}

	for _, test := range cases {
		ts := slackServer(test.ok)
		err := Slack(ts.URL, "bot", "yo")
		if test.ok && err != nil {
			t.Errorf("expected nil but got %s", err)
		}
		if !test.ok && err == nil {
			t.Error("expected err but got nil")
		}
	}
}

func slackServer(ok bool) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ok {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
}
