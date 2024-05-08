package service

import (
	"fmt"
	"net/http"
	"strconv"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type HandlerWS struct {
	hub *Hub
}

func NewHandlerWS(h *Hub) *HandlerWS {
	return &HandlerWS{
		hub: h,
	}
}

type CreateRoomReq struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (h *HandlerWS) CreateRoom(c *gin.Context) {
	var req CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.hub.Chats[req.Id] = &chat.ChatList{
		Id:      req.Id,
		Title:   req.Name,
		UsersId: make(map[string]*chat.Client),
	}
	fmt.Println("УХХХХХХ")
	fmt.Println(h.hub.Chats, "::", req.Id, "::", h.hub.Chats[req.Id])
	c.JSON(http.StatusOK, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *HandlerWS) JoinRoom(c *gin.Context, CreateItem func(userId, listId int, item chat.ChatItem) (int, error)) {
	fmt.Println("1")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("2")

	roomId := c.Param("roomId")
	clientId := c.Query("userId")
	username := c.Query("username")

	cl := &Client{
		&chat.Client{
			Conn:     conn,
			Message:  make(chan *chat.ChatItem, 10),
			Id:       clientId,
			RoomId:   roomId,
			Username: username,
		},
	}

	var description = "Пользователь зашел в комнату"
	clientIdNum, err := strconv.Atoi(clientId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("3", clientIdNum)

	roomIdNum, err := strconv.Atoi(roomId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("4")

	idChatItem, err := CreateItem(clientIdNum, roomIdNum, chat.ChatItem{
		Chatlist_id: roomId,
		User_id:     clientId,
		Username:    username,
		Description: description,
	})
	if err != nil {
		fmt.Println("123123")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("5")

	m := &chat.ChatItem{
		Description: description,
		Id:          idChatItem,
		User_id:     clientId,
		Username:    username,
		Chatlist_id: roomId,
	}

	h.hub.Register <- cl.Client
	fmt.Println("6")
	h.hub.Broadcast <- m
	fmt.Println("7")

	go cl.writeChatItem(clientId)
	cl.ReadChatItem(h.hub, clientId)
	fmt.Println("8")
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *HandlerWS) GetRooms(c *gin.Context) {
	rooms := make([]RoomRes, 0)

	for _, r := range h.hub.Chats {
		rooms = append(rooms, RoomRes{
			ID:   r.Id,
			Name: r.Title,
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"rooms": rooms,
	})
}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *HandlerWS) GetClients(c *gin.Context) {
	var clients []ClientRes
	roomId := c.Param("roomId")

	if _, ok := h.hub.Chats[roomId]; !ok {
		clients = make([]ClientRes, 0)
		c.JSON(http.StatusOK, clients)
	}

	for _, c := range h.hub.Chats[roomId].UsersId {
		clients = append(clients, ClientRes{
			ID:       c.Id,
			Username: c.Username,
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"clients": clients,
	})
}
