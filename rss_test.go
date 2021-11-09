package golang_rss_reader_package

import (
	"reflect"
	"testing"
)

//TestParseEmptyUrl should return an empty array
func TestParseEmptyUrl(t *testing.T) {

	emptyArray := Parse([]string{})

	if reflect.DeepEqual(emptyArray, []RssItem{}) {
		t.Logf("Parse() succeeded, expected %v, got %v", []RssItem{}, emptyArray)
	} else {
		t.Errorf("Parse() failed, expected %v, got %v", []RssItem{}, emptyArray)
	}
}

// TestParseInvalidUrlScheme should return an array of the valid urls, disregarding the erroneous one
// https://www.online-tech-tips.com/feed/ has 15 feeds
// test.com has a missing scheme
func TestParseInvalidUrlScheme(t *testing.T) {

	feeds := Parse([]string{
		"test.com",
		"https://www.online-tech-tips.com/feed/"})

	if len(feeds) == 15 {
		t.Logf("Expected result length given. Expected %v, got %v", 15, len(feeds))

	} else {

		t.Errorf("Parse() failed where the expected feed length is %v, got %v", 15, len(feeds))
	}
}


// TestParseInvalidRSSFeed same as above test, where the error here is invalid rss feed given
func TestParseInvalidRSSFeed(t *testing.T) {

	feeds := Parse([]string{
		"https://facebook.com",
		"https://www.online-tech-tips.com/feed/",
	})

	if len(feeds) == 15 {
		t.Logf("Expected result length given. Expected %v, got %v", 15, len(feeds))

	} else {

		t.Errorf("Parse() failed where the expected feed length is %v, got %v", 15, len(feeds))
	}
}

// TestParse shows the correct operation of parse with 2 feeds
// https://www.theboltonnews.co.uk/news/rss/ has 50 feeds
// https://www.online-tech-tips.com/feed/ has 15
func TestParse(t *testing.T) {

	feeds := Parse([]string{
		"https://www.theboltonnews.co.uk/news/rss/",
		"https://www.online-tech-tips.com/feed/",
	})

	if len(feeds) == 65 {
		t.Logf("Expected result length given. Expected %v, got %v", 65, len(feeds))

	} else {

		t.Errorf("Parse() failed where the expected feed length is %v, got %v", 65, len(feeds))
	}
}
