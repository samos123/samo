package main

import (
	"github.com/gorilla/mux"
	"github.com/samos123/samo"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	w := samo.SetUpVideoWatcher()
	router.HandleFunc("/", samo.VideosIndex)
	router.HandleFunc("/view/{videoId}", samo.VideosServe)
	router.Handle("/dir/add", samo.DirAddHTTPHandler(w))
	log.Fatal(http.ListenAndServe(":8080", router))
}
