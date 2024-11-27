package main

import (
	"fmt"
	"log"

	"fileck/config"
	"fileck/slack"
)

func main() {
	client := slack.CreateSlackClient()
	channels := config.LoadConfig().ChannelIDs
	files := []string{"./payment.jpeg"}

	for _, file := range files {
		size, err := slack.CheckFileSize(file)
		if err != nil {
			log.Printf("Failed to check file size for %s: %v\n", file, err)
			continue
		}

		for _, channel := range channels {
			err := slack.UploadFile(client, file, channel, size)
			if err != nil {
				log.Printf("Error uploading file %s to channel %s: %v\n", file, channel, err)
				continue
			}

			fmt.Printf("Successfully uploaded file %s to channel %s\n", file, channel)
		}
	}
}
