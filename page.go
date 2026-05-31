package ps3

const pageBucketName = "page"

func CreatePageImageUploadURLs(folder PageR2Folder, inputs []ImageUploadInput) ([]string, error) {
	urls := make([]string, len(inputs))
	for i, input := range inputs {
		url, err := presignUpload(getImageKey(input.PageID, input.ImageID), input.ContentType, pageBucketName)
		if err != nil {
			return nil, err
		}
		urls[i] = url
	}
	return urls, nil
}
