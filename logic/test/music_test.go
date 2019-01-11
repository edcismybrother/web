package test

import (
	"pylearn/web/logic"
	"testing"
	"time"
)

func TestMusic(t *testing.T) {
	logic.PutMusicInDownLoadChannel("hah1", "1111", 2)
	logic.PutMusicInDownLoadChannel("hah2", "1111", 3)
	logic.PutMusicInDownLoadChannel("hah3", "1111", 4)
	logic.PutMusicInDownLoadChannel("hah4", "1111", 3)
	time.Sleep(10 * time.Second)
}
