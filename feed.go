package golang_rss_reader_package

import "time"

//date type for date-time parsing as it is not straight-forward with go
type date string

//parse (date function) and returns Time, error
func (d date) parse() (time.Time, error) {

	t, err := d.parseWithFormat(time.RFC822) // RSS 2.0 spec
	if err != nil {

		t, err = d.parseWithFormat(time.RFC3339) // Atom
	}

	return t, err
}

//parseWithFormat (date function), takes a string and returns Time, error
func (d date) parseWithFormat(format string) (time.Time, error) {

	return time.Parse(format, string(d))
}

//channel struct for RSS input (Used for Input RSS)
type channel struct {

	Title         string `xml:"title"`
	Link          string `xml:"link"`
	Description   string `xml:"description"`
	Language      string `xml:"language"`
	LastBuildDate date   `xml:"lastBuildDate"`
	Item          []item `xml:"item"`
}

//itemEnclosure struct for each Item Enclosure (Used for Input RSS)
type itemEnclosure struct {

	URL  string `xml:"url,attr"`
	Type string `xml:"type,attr"`
}

type source struct {
	Title string `xml:"title,omitempty"`
	URL   string `xml:"url,omitempty"`
}

//item struct for each Item in the Channel (Used for Input RSS)
type item struct {

	Title       string          `xml:"title"`
	Link        string          `xml:"link"`
	Comments    string          `xml:"comments"`
	PubDate     date            `xml:"pubDate"`
	GUID        string          `xml:"guid"`
	Category    []string        `xml:"category"`
	Enclosure   []itemEnclosure `xml:"enclosure"`
	Source 		source 			`xml:"source"`
	Description string          `xml:"description"`
	Author      string          `xml:"author"`
	Content     string          `xml:"content"`
	FullText    string          `xml:"full-text"`
}
