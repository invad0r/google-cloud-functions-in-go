package image_resizer

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"strconv"

	"github.com/disintegration/imaging"
)

func ResizeImage(w http.ResponseWriter, r *http.Request) {
	// parse the url query string into ResizeParams
	p, err := ParseQuery(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// fetch input image and resize
	img, err := FetchAndResizeImage(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//encode output image to jpeg buffer
	encoded, err := EncodeImageToJpg(img)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// set Content-Tye and Content-Length headers
	w.Header().Set("Content-Type", "image/jpgeg")
	w.Header().Set("Content-Lenght", strconv.Itoa(encoded.Len()))

	// write output image to http response body
	_, err = io.Copy(w, encoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type ResizeParams struct {
	url    string
	height int
	width  int
}

func ParseQuery(r *http.Request) (*ResizeParams, error) {
	var p ResizeParams
	query := r.URL.Query()
	url := query.Get("url")
	if url == "" {
		return &p, errors.New("Url Param 'url' is missing")
	}

	width, _ := strconv.Atoi(query.Get("width"))
	height, _ := strconv.Atoi(query.Get("height"))
	if width == 0 && height == 0 {
		return &p, errors.New("Url Param 'height' or 'width' must be set")
	}

	p = NewResizeParams(url, height, width)

	return &p, nil
}

func NewResizeParams(url string, height int, width int) ResizeParams {
	return ResizeParams{url, height, width}
}

func FetchAndResizeImage(p *ResizeParams) (*image.Image, error) {
	var dst image.Image

	// fetch input data
	response, err := http.Get(p.url)
	if err != nil {
		return &dst, err
	}
	defer response.Body.Close()

	//decode input data to image
	src, _, err := image.Decode(response.Body)
	if err != nil {
		return &dst, err
	}

	// resize input image
	dst = imaging.Resize(src, p.width, p.height, imaging.Lanczos)

	return &dst, nil
}

func EncodeImageToJpg(image *image.Image) (*byte.Buffer, error) {
	encoded := &bytes.Buffer{}
	err := jpeg.Encode(encoded, *img, nil)
	return encoded, err
}
