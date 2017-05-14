package samo

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Video struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Duration int    `json:"duration"`
	At       int    `json:"at"`
}

type Videos []Video

type Dir struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Directories []Dir

func VideosIndex(w http.ResponseWriter, r *http.Request) {
	videos := Videos{
		Video{Name: "Test", Path: "/home"},
		Video{Name: "Test", Path: "/home"},
	}
	log.Println("Serving videos")

	json.NewEncoder(w).Encode(videos)
}

func VideosServe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	videoId := vars["videoId"]
	fmt.Fprintln(w, "Video ID:", videoId)
}
