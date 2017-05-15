package samo

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"net/http"
)

type VideoWatcher struct {
	*fsnotify.Watcher
}

func (watcher *VideoWatcher) ProcessEvents() {
	for {
		select {
		case event := <-watcher.Events:
			log.Println("event:", event)
		case err := <-watcher.Errors:
			log.Println("error:", err)
		}
	}
}

func NewVideoWatcher(dirs ...string) *VideoWatcher {
	watcher, err := fsnotify.NewWatcher()
	vw := &VideoWatcher{watcher}
	if err != nil {
		log.Fatal(err)
	}
	go vw.ProcessEvents()

	for _, dir := range dirs {
		err = vw.Add(dir)
		if err != nil {
			log.Fatal(err)
		}
	}
	return vw
}

func DirAddHTTPHandler(w *VideoWatcher) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		path := r.FormValue("path")
		if err := w.Add(path); err != nil {
			log.Fatalln(err)
		}
		fmt.Fprintln(rw, "Added dir:", path)
	})
}
