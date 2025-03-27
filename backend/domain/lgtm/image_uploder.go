package lgtm

type ImageUploader interface {
	Upload(LGTMImage) (imageURL string, err error)
}
