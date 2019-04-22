package models

import (
	"time"

	"github.com/nateupstairs/sync/db"
)

// Asset model
type Asset struct {
	new      bool
	ID       int64
	Success  int64
	Created  int64
	Updated  int64
	Filename string
}

// NewAsset creation of Asset model
func NewAsset() *Asset {
	x := new(Asset)
	x.Success = 0
	x.new = true
	return x
}

// Save Asset model
func (a *Asset) Save() {
	now := time.Now().UnixNano()

	if a.new == true {
		id := db.CreateAsset(now)
		a.ID = id
		a.Created = now
		a.Updated = now
		a.new = false
	} else {
		db.SaveAsset(now, a.ID, a.Filename)
		a.Updated = now
		a.Success = 1
	}
}
