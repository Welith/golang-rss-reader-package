# golang-rss-reader-package
A golang rss parser which parses an array of urls asynchronously.

# Specification
1. Use the latest stable Go version.
2. Take an array of urls and parse asynchronously their RSS feeds.
3. Return an array of items in the following structure:
```
type RssItem struct {

   Title string
   Source string
   SourceUrl string
   Link string
   PublishDate date
   Description string
}
```

## Installation
Installation is done with `go get`: <br>
`go get -u github.com/Welith/golang-rss-reader-package`

## Usage
The Parse() method accepts an array of URL feeds and parses them (the feeds need to be valid url strings), returning an array of RssItem structs. Erroneous URLs are disregarded. Each URL is handled asynchronous using goroutines and channels.

## Tests
Generic unit tests have been provided (I am not that good at golang testing and mocking, so I did my best)

##Notes
As per the requirement the PubDate should be a time.Time. As golang's serilialization/deserialization is not straight-forward with dates, I have created my own parser for the date format, together with a custom time string.
