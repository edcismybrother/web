package main

import (
	"fmt"
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
	http.HandleFunc("/music/", music)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/music.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func music(w http.ResponseWriter, r *http.Request) {
	fmt.Println("lenmon走音乐")
	p := MusicPath + r.URL.Path[len("/music"):]
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
