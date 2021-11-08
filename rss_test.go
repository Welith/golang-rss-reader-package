package golang_rss_reader_package

import (
	"reflect"
	"testing"
)

func TestParseEmptyUrl(t *testing.T) {

	emptyArray, err := Parse([]string{})

	if reflect.DeepEqual(emptyArray, []RssItem{}) && err == nil {
		t.Logf("Parse() succeeded, expected %v, got %v", []RssItem{}, emptyArray)
	} else {
		t.Errorf("Parse() failed, expected %v, got %v", []RssItem{}, emptyArray)
		if err != nil {
			t.Errorf("Actual error %v: ", err)
		}
	}
}

func TestParseInvalidUrlScheme(t *testing.T) {

	invalidUrl, err := Parse([]string{"test"})

	t.Logf("%v", invalidUrl)

	if err != nil {
		t.Logf("Expected error: %v, got: %v ", "Get \"test\": unsupported protocol scheme \"\"", err)

	} else {

		t.Errorf("Parse() failed to produce error, expected %v, got %v", "Get \"test\": unsupported protocol scheme \"\"", invalidUrl)
	}
}



func TestParseInvalidRSSFeed(t *testing.T) {

	invalidUrl, err := Parse([]string{"https://facebook.com"})

	if err != nil {
		t.Logf("Expected error: %v, got: %v ", "XML syntax error on line 3: invalid character entity & (no semicolon)", err)

	} else {

		t.Errorf("Parse() failed to produce error, expected %v, got %v", "XML syntax error on line 3: invalid character entity & (no semicolon)", invalidUrl)
	}
}

func TestParse(t *testing.T) {

	feedItems, err := Parse([]string{"https://www.theboltonnews.co.uk/news/rss/"})

	if err == nil {
		if feedItems[0].Title != "" {
			t.Logf("Parse succeeded.")
		}
	} else {
		t.Errorf("Parse failed with error %v:", err)
	}
}
