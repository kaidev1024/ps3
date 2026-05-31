package ps3

import "fmt"

type ImageUploadInput struct {
	ImageID     string
	PageID      string
	ContentType string
}

func getImageKey(pageID, imageID string) string {
	return fmt.Sprintf("%s/%s", pageID, imageID)
}
