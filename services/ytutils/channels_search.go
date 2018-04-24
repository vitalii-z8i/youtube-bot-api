package ytutils

import (
	"log"
	"net/http"

	"github.com/vtl-pol/youtube-bot-api/config"
	"github.com/vtl-pol/youtube-bot-api/entities"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
	// "google.golang.org/api/youtube/v3"
)

var (
	query            = "Google"
	maxResults int64 = 5
)

// ChannelsSearch loactes one or many YT channels by a given keyword
func ChannelsSearch(keyword string) (finalResult [][]entities.YTChannel, err error) {
	client := &http.Client{
		Transport: &transport.APIKey{Key: config.YT.DeveloperKey},
	}

	ytClient, err := youtube.New(client)
	if err != nil {
		log.Printf("Error creating new YouTube client: %v\n", err)
		return nil, err
	}
	call := ytClient.Search.List("snippet").
		Q(keyword).
		Type("channel").
		MaxResults(maxResults)

	response, err := call.Do()
	if err != nil {
		log.Printf("Error making search API call: %v", err)
		return nil, err
	}

	for _, item := range response.Items {
		res := append([]entities.YTChannel{}, entities.YTChannel{
			ChannelID:   item.Snippet.ChannelId,
			ChannelName: item.Snippet.ChannelTitle,
			ChannelInfo: item.Snippet.Description})
		finalResult = append(finalResult, res)

	}

	return finalResult, err

}
