package main

import (
	"log"
	"net/http"

	"github.com/nateupstairs/sync/router"
)

func main() {
	log.Fatal(http.ListenAndServe(":3000", router.Get()))
}
