package dotenv

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

func readDotEnv() *map[string]string {
	file, err := os.Open(".env")
	defer file.Close()
	if err != nil {
		log.Panicln("open file with error", err)
	}
	stat, err := file.Stat()
	if err != nil {
		log.Panicln("open file with error", err)
	}
	buffer := make([]byte, stat.Size())
	file.Read(buffer)
	data := string(buffer)
	result := make(map[string]string)
	for _, row := range strings.Split(data, "\n") {
		keyVal := strings.Split(row, "=")
		if len(keyVal) != 2 {
			// fmt.Printf("error in line %d, cannot parse: %s\n", line+1, row)
			continue
		}
		if row[0] == '#' {
			continue
		}
		result[keyVal[0]] = strings.Trim(keyVal[1], "\n\r\t")
	}
	fmt.Println(result)
	return &result
}
func watchDotEnv(onChange chan interface{}) {
	file(onChange, ".env")
}
func setEnv(data *map[string]string) {
	for k, v := range *data {
		os.Setenv(k, v)
	}
}

// Load .env file using os.SetEnv method to be accessible across the application
func Load() {
	setEnv(readDotEnv())
}

// Watch .env file for changes and then reloads the .env file
func Watch() chan interface{} {
	notifier := make(chan interface{})
	exportedNotifier := make(chan interface{})
	go watchDotEnv(notifier)
	go func() {
		for range notifier {
			time.Sleep(time.Microsecond * 10)
			Load()
			exportedNotifier <- 0
		}
	}()
	return exportedNotifier
}

func file(notifier chan interface{}, files ...string) {
	if len(files) < 1 {
		panic("must specify at least one file to watch")
	}
	w, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("creating a new watcher: %s", err)
	}
	defer w.Close()
	go fileLoop(notifier, w, files)
	for _, p := range files {
		st, err := os.Lstat(p)
		if err != nil {
			fmt.Printf("%s", err)
		}
		if st.IsDir() {
			fmt.Printf("%q is a directory, not a file", p)
		}
		err = w.Add(filepath.Dir(p))
		if err != nil {
			fmt.Printf("%q: %s", p, err)
		}
	}
	<-make(chan struct{})
}

func fileLoop(notifier chan interface{}, w *fsnotify.Watcher, files []string) {
	i := 0
	for {
		select {
		case err, ok := <-w.Errors:
			if !ok {
				return
			}
			fmt.Printf("ERROR: %s", err)
		case e, ok := <-w.Events:
			if !ok {
				return
			}
			var found bool
			for _, f := range files {
				if f == e.Name {
					found = true
				}
			}
			if !found {
				continue
			}
			notifier <- 0
			i++
		}
	}
}
