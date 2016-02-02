package dnsbl

import (
	"testing"
)

func TestQuery(t *testing.T) {
	r, err := Query("zen.spamhaus.org", "google.com")
	if err != nil {
		t.Fatal(err)
	}

	// Check if google.com is on the list and if it is then something
	// is wrong.
	if r.Listed {
		t.Fatal("google.com is on the zen.spamhaus.org blacklist.")
	}
}
