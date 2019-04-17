package controller

import (
	"fmt"
	"net/http"
	"path"

	"github.com/davecgh/go-spew/spew"
	"github.com/nateupstairs/sync/config"
	"github.com/nateupstairs/sync/db"
	"github.com/nateupstairs/sync/util"
)

// PostVideo post a video
func PostVideo(w http.ResponseWriter, r *http.Request) {
	id := db.CreateAsset()
	config := config.Get()

	fileName, err := util.FileUpload(r, path.Join(config.HomeDir, fmt.Sprintf("%05d", id)))

	if err != nil {
		spew.Dump(fileName)
		spew.Dump(err)
	}

	w.Write([]byte("Gorilla!\n"))
}
