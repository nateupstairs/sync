package util

import (
	"io"
	"net/http"
	"os"
)

func FileUpload(r *http.Request) (string, error) {
	r.ParseMultipartForm(32 << 20)

	var fileName string

	file, handler, err := r.FormFile("file")

	if err != nil {
		return fileName, err
	}

	defer file.Close()

	f, err := os.OpenFile("./uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return fileName, err
	}

	fileName = handler.Filename
	defer f.Close()

	io.Copy(f, file)

	return fileName, err
}
