package util

import (
	"io"
	"net/http"
	"os"
	"path"
)

// FileUpload for POST multipart file uploads
func FileUpload(r *http.Request, syncFolder string) (string, error) {
	r.ParseMultipartForm(32 << 20)

	var fileName string

	file, handler, err := r.FormFile("file")
	if err != nil {
		return fileName, err
	}
	defer file.Close()

	err = os.MkdirAll(syncFolder, os.ModePerm)
	if err != nil {
		return fileName, err
	}

	f, err := os.OpenFile(path.Join(syncFolder, handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fileName, err
	}
	defer f.Close()

	fileName = handler.Filename

	io.Copy(f, file)

	return fileName, err
}
