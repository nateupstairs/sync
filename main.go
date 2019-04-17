package main

import (
	"log"
	"net/http"

	"github.com/nateupstairs/sync/db"
	"github.com/nateupstairs/sync/router"
)

func main() {
	//db.Test()
	db.CreateAsset()

	log.Fatal(http.ListenAndServe(":3000", router.Get()))
}
