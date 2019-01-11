package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type st struct {
	Name string
}

var MusicPath = "music"

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/music/", musicIndex)
	http.HandleFunc("/music/src", musicSource)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/index.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}
func musicIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/music.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func musicSource(w http.ResponseWriter, r *http.Request) {
	p := MusicPath + r.URL.Path[len("/music/src"):]
	file, _ := os.Stat(p)
	if file.IsDir() {
		f, _ := os.Open(p)
		files, _ := f.Readdir(0)
		list := ""
		for _, fi := range files {
			list = list + fi.Name() + ","
		}
		w.Write([]byte(list))
		return
	}
	http.ServeFile(w, r, p)
}
