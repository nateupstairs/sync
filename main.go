package main

import (
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/nateupstairs/sync/db"
	"github.com/nateupstairs/sync/util"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	imageName, err := util.FileUpload(r)

	if err != nil {
		spew.Dump(imageName)
		spew.Dump(err)
	}

	w.Write([]byte("Gorilla!\n"))
}

func main() {
	db.Test()
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":3000", r))
}
