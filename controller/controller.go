package controller

import (
	"fmt"
	"net/http"
	"path"

	"github.com/davecgh/go-spew/spew"
	"github.com/nateupstairs/sync/config"
	"github.com/nateupstairs/sync/models"
	"github.com/nateupstairs/sync/util"
)

// PostVideo post a video
func PostVideo(w http.ResponseWriter, r *http.Request) {
	config := config.Get()

	asset := models.NewAsset()
	asset.Save()

	filename, err := util.FileUpload(r, path.Join(
		config.HomeDir,
		"Desktop",
		"SYNC",
		fmt.Sprintf("%05d", asset.ID)))

	if err != nil {
		spew.Dump(err)
	}

	asset.Filename = filename
	asset.Save()

	w.Write([]byte("Success!\n"))
}
