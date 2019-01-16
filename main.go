package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type st struct {
	Name string
}

var MusicPath = "music"

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/music/", musicIndex)
	http.HandleFunc("/music/src/", musicSource)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
		if err != nil {
			fmt.Println(err)
			return
		}
		go func() {
			for {
				_, data, err := conn.ReadMessage()
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(string(data))
				err = conn.WriteMessage(1, []byte("hello,client"))
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
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
