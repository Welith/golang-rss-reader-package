package main

import (
	"encoding/xml"
	"fmt"
	"github.com/paulrosania/go-charset/charset"
	"net/http"
	"time"
)

//date type for date-time parsing as it is not straight-forward with go
type date string

//Parse (date function) and returns Time, error
func (d date) Parse() (time.Time, error) {

	t, err := d.parseWithFormat(time.RFC822) // RSS 2.0 spec
	if err != nil {
		t, err = d.parseWithFormat(time.RFC3339) // Atom
	}

	return t, err
}

//ParseWithFormat (date function), takes a string and returns Time, error
func (d date) parseWithFormat(format string) (time.Time, error) {
	return time.Parse(format, string(d))
}

//Channel struct for RSS input (Used for Input RSS)
type channel struct {
	Title         string `xml:"title"`
	Link          string `xml:"link"`
	Description   string `xml:"description"`
	Language      string `xml:"language"`
	LastBuildDate date   `xml:"lastBuildDate"`
	Item          []item `xml:"item"`
}

//ItemEnclosure struct for each Item Enclosure (Used for Input RSS)
type itemEnclosure struct {
	URL  string `xml:"url,attr"`
	Type string `xml:"type,attr"`
}

//Item struct for each Item in the Channel (Used for Input RSS)
type item struct {
	Title       string          `xml:"title"`
	Link        string          `xml:"link"`
	Comments    string          `xml:"comments"`
	PubDate     date            `xml:"pubDate"`
	GUID        string          `xml:"guid"`
	Category    []string        `xml:"category"`
	Enclosure   []itemEnclosure `xml:"enclosure"`
	Description string          `xml:"description"`
	Author      string          `xml:"author"`
	Content     string          `xml:"content"`
	FullText    string          `xml:"full-text"`
}


//RssItem (Used for required output)
type RssItem struct {

	Title string
	Source string
	SourceUrl string
	Link string
	PublishDate date
	Description string
}

func main()  {

	urls := []string{
		"https://helpdeskgeek.com/feed/",
	}

	c := make(chan RssItem)

	for _, url := range urls {
		go Parse(url, c)
	}

	result := make([]RssItem, len(urls))

	for i, _ := range result {
		result[i] = <-c
		fmt.Printf("Error GET: %v\n", result[i])
	}
}

func Parse(url string, c chan RssItem) {

	resp, err := http.Get(url)

	if err != nil {

		fmt.Printf("Error GET: %v\n", err)
		return
	}

	defer resp.Body.Close()

	xmlDecoder := xml.NewDecoder(resp.Body)
	xmlDecoder.CharsetReader = charset.NewReader

	var rss struct {

		Channel channel `xml:"channel"`
	}

	if err := xmlDecoder.Decode(&rss); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, item := range rss.Channel.Item {

		c <- RssItem{
			Title:       item.Title,
			Source:      item.Author,
			SourceUrl:   "",
			Link:        item.Link,
			PublishDate: item.PubDate,
			Description: item.Description,
		}
	}
}