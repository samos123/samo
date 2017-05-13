package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/radovskyb/watcher"
	"log"
	"net/http"
	"os/exec"
	"time"
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
	watchDir(path)
	out, err := exec.Command("find", path, "-name \"*.mp4\" -or -name \"*.mkv\"").Output()
	fmt.Fprintln(w, "path:", path)
	fmt.Fprintln(w, "out:", out)
	fmt.Fprintln(w, "err:", err)
}

func watchDir(dir string) {
	w := watcher.New()
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Create, watcher.Rename)
	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Println(event)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch test_folder recursively for changes.
	if err := w.AddRecursive(dir); err != nil {
		log.Fatalln(err)
	}
	// Print a list of all of the files and folders currently
	// being watched and their paths.
	for path, f := range w.WatchedFiles() {
		fmt.Printf("%s: %s\n", path, f.Name())
	}
	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
