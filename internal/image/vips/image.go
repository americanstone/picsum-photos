package vips

import "github.com/DMarby/picsum-photos/internal/vips"

// resizedImage is a resized image
type resizedImage struct {
	vipsImage vips.Image
}

// resizeImage loads an image from a byte buffer, resizes it and returns an Image object for further use
// Note that it does not use the processor worker queue, use ProcessImage for that
func resizeImage(buffer []byte, width int, height int) (*resizedImage, error) {
	image, err := vips.ResizeImage(buffer, width, height)

	if err != nil {
		return nil, err
	}

	return &resizedImage{
		vipsImage: image,
	}, nil
}

// grayscale turns an image into grayscale
func (i *resizedImage) grayscale() (*resizedImage, error) {
	image, err := vips.Grayscale(i.vipsImage)
	if err != nil {
		return nil, err
	}

	return &resizedImage{
		vipsImage: image,
	}, nil
}

// blur applies gaussian blur to an image
func (i *resizedImage) blur(blur int) (*resizedImage, error) {
	image, err := vips.Blur(i.vipsImage, blur)
	if err != nil {
		return nil, err
	}

	return &resizedImage{
		vipsImage: image,
	}, nil
}

// setUserComment sets the exif usercomment
func (i *resizedImage) setUserComment(comment string) {
	vips.SetUserComment(i.vipsImage, comment)
}

// saveToJpegBuffer returns the image as a JPEG byte buffer
func (i *resizedImage) saveToJpegBuffer() ([]byte, error) {
	imageBuffer, err := vips.SaveToJpegBuffer(i.vipsImage)

	if err != nil {
		return nil, err
	}

	return imageBuffer, nil
}

// saveToWebPBuffer returns the image as a WebP byte buffer
func (i *resizedImage) saveToWebPBuffer() ([]byte, error) {
	imageBuffer, err := vips.SaveToWebPBuffer(i.vipsImage)

	if err != nil {
		return nil, err
	}

	return imageBuffer, nil
}
