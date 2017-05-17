package samo

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"net/http"
	"os"
	"strings"
)

type VideoWatcher struct {
	*fsnotify.Watcher
	SupportedFormats []string
}

func NewVideoWatcher(dirs ...string) *VideoWatcher {
	watcher, err := fsnotify.NewWatcher()
	vw := &VideoWatcher{watcher, []string{".mkv", ".mp4", ".avi"}}
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

func (watcher *VideoWatcher) ProcessEvents() {
	for {
		select {
		case event := <-watcher.Events:
			log.Println("event:", event)
			if event.Op&fsnotify.Create == fsnotify.Create {
				if isDir(event.Name) {
					log.Println("Adding subdirectory")
					watcher.Add(event.Name)
				} else if watcher.isSupportedFormat(event.Name) {
					log.Println("Adding video")
				}
			}
		case err := <-watcher.Errors:
			log.Println("error:", err)
		}
	}
}

func isDir(path string) bool {
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return true
	} else {
		return false
	}
}

func (watcher *VideoWatcher) isSupportedFormat(path string) bool {
	for _, format := range watcher.SupportedFormats {
		if strings.HasSuffix(path, format) {
			return true
		}
	}
	return false
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
