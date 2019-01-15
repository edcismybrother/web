package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"pylearn/web/logic"
	"sync"
)

type st struct {
	Name string
}

var MusicPath = "music"

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		defer wg.Done()
		http.HandleFunc("/", index)
		http.HandleFunc("/music/", musicIndex)
		http.HandleFunc("/music/src/", musicSource)
		http.ListenAndServe(":8080", nil)
	}()
	go func() {
		// 游戏服
		defer wg.Done()
		lis, err := net.Listen("tcp", "127.0.0.1:8123")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("游戏服开启")
		for {
			conn, err := lis.Accept()
			if err != nil {
				break
			}
			fmt.Println("游戏连接成功")
			go logic.NewPlayer(conn)
		}
	}()
	wg.Wait()
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
