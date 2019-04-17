package router

import (
	"github.com/gorilla/mux"
	"github.com/nateupstairs/sync/controller"
)

var r *mux.Router

func init() {
	r = mux.NewRouter()

	r.HandleFunc("/", controller.PostVideo)
}

// Get router
func Get() *mux.Router {
	return r
}
