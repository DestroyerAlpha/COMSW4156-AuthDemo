package main

import (
	"log"
	"net/http"

	"github.com/DestroyerAlpha/COMSW4156-AuthDemo/router"
)

func main() {
	log.Fatal(http.ListenAndServe(":10000", router.SetupRoutes()))
}
