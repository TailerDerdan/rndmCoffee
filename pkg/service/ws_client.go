package service

import (
	"fmt"
	"log"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/gorilla/websocket"
)

type Client struct {
	*chat.Client
}

func (c *Client) writeChatItem() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		message, ok := <-c.Message
		if !ok {
			return
		}
		fmt.Println(message.User_id, "айдишниккк111")
		fmt.Println(message.User_id, "айдишниккк222")

		c.Conn.WriteJSON(message)

	}
}

func (c *Client) ReadChatItem(hub *Hub, clientId string) {
	defer func() {
		hub.Unregister <- c.Client
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &chat.ChatItem{
			Description: string(m),
			Chatlist_id: c.RoomId,
			Username:    c.Username,
			User_id:     clientId,
		}

		fmt.Println(msg)

		hub.Broadcast <- msg
	}
}
