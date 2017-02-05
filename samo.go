package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/exec"
)

type Video struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Videos []Video

type Dir struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Directories []Dir

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", VideosIndex)
	router.HandleFunc("/view/{videoId}", VideosServe)
	router.HandleFunc("/dir/add", DirAdd)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func VideosIndex(w http.ResponseWriter, r *http.Request) {
	videos := Videos{
		Video{Name: "Test", Path: "/home"},
		Video{Name: "Test", Path: "/home"},
	}

	json.NewEncoder(w).Encode(videos)
}

func VideosServe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	videoId := vars["videoId"]
	fmt.Fprintln(w, "Video ID:", videoId)
}

func DirAdd(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	out, err := exec.Command("find", path, "-name \"*.mp4\" -or -name \"*.mkv\"").Output()
	fmt.Fprintln(w, "Dir:", path)
}
