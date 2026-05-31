package ps3

import "fmt"

const pageBucketName = "page"

type PageImageUploadInput struct {
	ImageID     string
	PageID      string
	ContentType string
}

func CreatePageImageUploadURLs(folder PageR2Folder, inputs []PageImageUploadInput) ([]string, error) {
	urls := make([]string, len(inputs))
	for i, input := range inputs {
		imageKey := fmt.Sprintf("%s/%s/%s", folder, input.PageID, input.ImageID)
		url, err := presignUpload(imageKey, input.ContentType, pageBucketName)
		if err != nil {
			return nil, err
		}
		urls[i] = url
	}
	return urls, nil
}
