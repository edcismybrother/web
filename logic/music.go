package logic

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
)

type musicInfo struct {
	name     string
	url      string
	formData map[string]string
}

var cdn = "https://music.163.com/weapi/song/enhance/player/url?csrf_token="

var musicChannel = make(chan musicInfo, 1000)

// PutMusicInDownLoadChannel 将音乐放入下载
func PutMusicInDownLoadChannel(name, url string, formData map[string]string) {
	fmt.Println("将" + name + "放入下载队列")
	musicChannel <- musicInfo{name: name, url: url, formData: formData}
}

func init() {
	go process()
}

func process() {
	for {
		m := <-musicChannel
		fmt.Printf("开始下载:%+v\n", m)
		go download(m)
	}
}

func download(m musicInfo) {
	c := http.Client{}
	req, err := http.NewRequest("POST", cdn, nil)
	if err != nil {
		// 错误处理
	}
	for k, v := range m.formData {
		req.Form.Set(k, v)
	}
	resp, err := c.Do(req)
	if err != nil {
		// 错误处理
	}
	read, err := gzip.NewReader(resp.Body)
	if err != nil {
		// 错误处理
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(read)
	data := buf.Bytes()
	// 从data中拿到music的实际储存地址，然后将地址中的音乐放到服务器本地
	fmt.Printf("资源：%v下载完毕\n", data)
}
