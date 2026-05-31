package ps3

import (
	"fmt"
)

const subjectBucketName = "subject"

type SubjectImageUploadInput struct {
	ImageID      string
	ContentType  string
	CreationHour int32
}

func CreateSubjectImageUploadURLs(folder SubjectR2Folder, inputs []SubjectImageUploadInput) ([]string, error) {
	urls := make([]string, len(inputs))
	for i, input := range inputs {
		imageKey := fmt.Sprintf("%s/%d/%s", folder, input.CreationHour, input.ImageID)
		url, err := presignUpload(imageKey, input.ContentType, subjectBucketName)
		if err != nil {
			return nil, err
		}
		urls[i] = url
	}
	return urls, nil
}
