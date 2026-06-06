package ps3

type ImageSize string

const (
	ImageSizeSm ImageSize = "sm"
	ImageSizeMd ImageSize = "md"
	ImageSizeLg ImageSize = "lg"
)

var imageSizes []ImageSize = []ImageSize{
	ImageSizeSm,
	ImageSizeMd,
}

const IMAGES_LEN = 2
