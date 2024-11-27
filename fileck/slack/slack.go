package slack

import (
	"fileck/config"
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func CreateSlackClient() *slack.Client {
	cfg := config.LoadConfig()
	return slack.New(cfg.SlackBotToken)
}

func CheckFileSize(file string) (int64, error) {
	info, err := os.Stat(file)
	if err != nil {
		return 0, fmt.Errorf("failed to stat file %s: %v", file, err)
	}

	size := info.Size()
	if size == 0 {
		return 0, fmt.Errorf("file %s has size 0, skipping upload", file)
	}

	return size, nil
}

func UploadFile(client *slack.Client, file string, channel string, size int64) error {
	params := slack.UploadFileV2Parameters{
		Channel:  channel,
		File:     file,
		Filename: file,
		FileSize: int(size),
	}

	upload, err := client.UploadFileV2(params)
	if err != nil {
		return fmt.Errorf("error uploading file %s: %v", file, err)
	}

	fmt.Printf("Uploaded file: %s\n", upload.Title)
	return nil
}
