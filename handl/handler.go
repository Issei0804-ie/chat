package handl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"goroutine/domain"
	"goroutine/room"
	"log"
	"strconv"
	"time"
)

const MAXROOMS = 100

type handler struct {
	rooms  [MAXROOMS]room.Room
	userid int
}

func NewHandler() *handler {
	userid := 0
	h := &handler{userid: userid}
	return h
}

func (h *handler) ChatRoom(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

//connect
func (h *handler) Connect(c *gin.Context) {
	fmt.Println("ReadMessage")
	param := c.Param("roomid")
	roomID, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(200, domain.ReadMessage{0, "Error:roomid", ""})
	}

	var upgrader = websocket.Upgrader{HandshakeTimeout: time.Second * 100, EnableCompression: true}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal("error upgrading GET request to a websocket::", err)
	}
	defer conn.Close()
	//新規ルームの作成

	if h.rooms[roomID].GetStatus() != true {
		h.rooms[roomID] = room.NewRoom()
	}
	h.rooms[roomID].Read(conn)
}

func (h *handler) GetUserID(c *gin.Context) {
	fmt.Println("UserID")
	u := domain.User{h.userid}
	h.userid++
	c.JSON(200, u)
}
