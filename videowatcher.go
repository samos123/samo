package samo

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"net/http"
)

func SetUpVideoWatcher(dirs ...string) *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	for _, dir := range dirs {
		err = watcher.Add(dir)
		if err != nil {
			log.Fatal(err)
		}
	}
	return watcher
}

func DirAddHTTPHandler(w *fsnotify.Watcher) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		path := r.FormValue("path")
		if err := w.Add(path); err != nil {
			log.Fatalln(err)
		}
		fmt.Fprintln(rw, "Added dir:", path)
	})
}
