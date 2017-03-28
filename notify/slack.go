package notify

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Slack notifies the slack URL with the title and message.
// If an empty string is passed for URL the notification is skipped without error,
// this allows for code to conditionally send notifications. For example an
// application could use a `SLACK_WEBHOOK_URL` environment variable to toggle on/off
// notifications without producing errors for the application.
//
// Example:
//
//    // `SLACK_WEBHOOK_URL` behaves like a toggle for enabling/disabling notifications
//    if err := notify.Slack(os.Getenv("SLACK_WEBHOOK_URL"), "heeeeey", "yoou guyyyszz <ANNOYING_GIF>"); err != nil {
//      // do something w/ err
//    }
//
// An error will only be returned if there is an issue with the network or Slack API
func Slack(url, title, message string) error {
	if url == "" {
		fmt.Println("Slack URL not set, skipping notification...")
		return nil
	}

	msg := messageBody(title, message)
	resp, err := http.Post(url, "application/json", strings.NewReader(msg))
	if err != nil {
		return fmt.Errorf("slack notification failed: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("slack notification failed: %s", b)
	}

	return nil
}

func messageBody(title, message string) string {
	return fmt.Sprintf(`{
  "username": "%s",
  "text": "%s"
}`, title, message)
}
