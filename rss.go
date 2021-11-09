package golang_rss_reader_package

import (
	"encoding/xml"
	"fmt"
	"github.com/paulrosania/go-charset/charset"
	"net/http"
)

//RssItem (Used for required output)
type RssItem struct {

	Title string 		`json:"title,omitempty"`
	Source string 		`json:"source,omitempty"`
	SourceUrl string 	`json:"source_url,omitempty"`
	Link string 		`json:"link,omitempty"`
	PublishDate date 	`json:"publish_date,omitempty"`
	Description string	`json:"description,omitempty"`
}

//Parse exported method used as a package function
func Parse(urls []string) []RssItem {

	if len(urls) == 0 {

		return []RssItem{}
	}

	c := make(chan []RssItem)
	defer close(c)

	var result []RssItem

	for _, url := range urls {

		go parseUrl(url, c)

		result = append(result, <- c...)
	}

	return result
}

func parseUrl(url string, c chan []RssItem) {

	resp, err := http.Get(url)

	if err != nil {

		fmt.Printf("RSS Package Error: %v \n", err)
		c <- []RssItem{}
	} else {

		defer resp.Body.Close()
		xmlDecoder := xml.NewDecoder(resp.Body)
		xmlDecoder.CharsetReader = charset.NewReader

		var rss struct {

			Channel channel `xml:"channel"`
		}

		if err = xmlDecoder.Decode(&rss); err != nil {

			fmt.Printf("RSS Package Error: %v \n", err)
			c <- []RssItem{}
		} else {

			var items []RssItem

			for _, item := range rss.Channel.Item {

				items = append(items, RssItem{

					Title:       item.Title,
					Source:      item.Source.Title,
					SourceUrl:   item.Source.URL,
					Link:        item.Link,
					PublishDate: item.PubDate,
					Description: item.Description,
				})
			}

			c <- items
		}
	}
}