package logic

import (
	"encoding/json"
	"fmt"
	"net"
)

// 交互

// Player 玩家
type Player struct {
	conn   net.Conn
	events chan Event
}

// Event 事件
type Event struct {
	Action  uint32
	Message []byte
}

const (
	Action_Move uint32 = iota
)

// map[name]

type Move struct {
	Direction int `json:"direction"`
	Step      int `json:"step"`
}

// ExecuteEvent 处理事件
func (p *Player) ExecuteEvent() {
	for {
		event := <-p.events
		fmt.Println("event:", event)
		switch event.Action {
		case Action_Move:
			am := &Move{}
			err := json.Unmarshal(event.Message, am)
			if err == nil {
				p.conn.Write([]byte("hello,move"))
			}
		}
	}
}

func NewPlayer(conn net.Conn) {
	p := &Player{
		conn:   conn,
		events: make(chan Event, 10),
	}
	go p.ExecuteEvent()
	for {
		b := make([]byte, 1024)
		_, err := conn.Read(b)
		if err != nil {
			return
		}
		fmt.Println(string(b))
		p.events <- Event{Action: 1, Message: b}
	}
}
