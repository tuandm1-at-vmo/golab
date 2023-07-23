package lab01

import (
	"reflect"
	"testing"
)

func TestCrawlAnchorsFrom(t *testing.T) {
	testCrawlAnchorsFrom(t, "https://tuanm.dev", []string{"tel:+84949500296", "https://instagram.com/tuanm_", "https://messenger.com/t/Teemoing", "https://github.com/Tuanm", "https://tuanm.dev"})
}

func testCrawlAnchorsFrom(t *testing.T, url string, expect []string) {
	actual, err := CrawlAnchorsFrom(url)
	if err != nil {
		t.Errorf("failed: %s", err.Error())
	}
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("failed: expect=%#v actual=%#v", expect, actual)
	}
}
