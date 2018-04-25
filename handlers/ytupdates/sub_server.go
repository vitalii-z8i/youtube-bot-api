package ytupdates

import (
	"encoding/xml"
	"log"
)

// YTFeed is a yt feed representation
type YTFeed struct {
	Video struct {
		ID         string `xml:"yt:videoId"`
		ChannelID  int64  `xml:"yt:channelId"`
		VideoTitle string `xml:"title"`
		URL        string `xml:"link"`
	} `xml:"entry"`
}

// ProcessUpdate - YouTube push-notifications handler func
func ProcessUpdate(contentType string, body []byte) {
	var feed YTFeed
	xmlError := xml.Unmarshal(body, &feed)

	if xmlError != nil {
		log.Printf("XML Parse Error %v", xmlError)

	} else {
		log.Printf("%+v\n", feed)
	}

}
