package service

import (
	"fmt"

	chat "github.com/MerBasNik/rndmCoffee"
)

type Hub chat.Hub

func NewHub() *Hub {
	return &Hub{
		Chats:      make(map[string]*chat.ChatList),
		Register:   make(chan *chat.Client),
		Unregister: make(chan *chat.Client),
		Broadcast:  make(chan *chat.ChatItem, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			fmt.Println(h.Chats[cl.RoomId], ":::", h.Chats, ":::", cl.RoomId)
			if _, ok := h.Chats[cl.RoomId]; ok {
				r := h.Chats[cl.RoomId]
				fmt.Println(h.Chats[cl.RoomId], ":::", h.Chats, ":::", cl.RoomId)
				if _, ok := r.UsersId[cl.Id]; !ok {
					r.UsersId[cl.Id] = cl
				}
			}
		case cl := <-h.Unregister:
			fmt.Println(h.Chats[cl.RoomId], "ЗАКРЫТО")
			if _, ok := h.Chats[cl.RoomId]; ok {
				if _, ok := h.Chats[cl.RoomId].UsersId[cl.Id]; ok {
					if len(h.Chats[cl.RoomId].UsersId) != 0 {

						h.Broadcast <- &chat.ChatItem{
							Description: "Пользователь покинул чат",
							Chatlist_id: cl.RoomId,
							Username:    cl.Username,
						}
					}
					delete(h.Chats[cl.RoomId].UsersId, cl.Id)
					close(cl.Message)
				}
			}
		case m := <-h.Broadcast:
			fmt.Println("ЯЯЯЯЯЯЯЯЯЯЯЯЯЯЯЯЯЯ")
			fmt.Println(h.Chats, "рука")
			fmt.Println(m.Chatlist_id, "ДЕНИС")
			fmt.Println(h.Chats[m.Chatlist_id], "нога")
			if _, ok := h.Chats[m.Chatlist_id]; ok {
				fmt.Println("ююююююююю")
				for _, cl := range h.Chats[m.Chatlist_id].UsersId {
					fmt.Println("ъъъъъъъъъ")
					cl.Message <- m
				}
			}
		}
	}
}
