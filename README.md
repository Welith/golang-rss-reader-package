# golang-rss-reader-package
A golang rss parser which parses an array of urls asynchronously.

## Installation

Installation is done with `go get`

`go get -u github.com/Welith/golang-rss-reader-package`

## Usage

The Parse() method accepts an array of URL feeds and parses them, returning an array of RssItem structs. Erroneous URLs are disregarded. Each URL is handled asynchronous using goroutines and channels.

## Tests

Generic unit tests have been provided (I am not that good at golang testing and mocking, so I did my best)


##Notes

As per the requirement the PubDate shoudl be a time.Time. As golang's serilialization/deserialization is not straight-forward with dates, I have created my own parser for the date format.
