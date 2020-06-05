package client

import (
	"fmt"
	"github.com/gorilla/websocket"
	"goroutine/domain"
)

type Client struct {
	conn   *websocket.Conn
	status bool
}

func NewClient(conn *websocket.Conn) Client {
	return Client{conn, true}

}

func (c *Client) GetConn() *websocket.Conn {
	return c.conn
}

func (c *Client) GetStatus() bool {
	return c.status
}

func (c *Client) Read(mCh chan domain.ReadMessage, stopCH chan bool) {
	for {
		m := domain.ReadMessage{}
		if err := c.conn.ReadJSON(&m); err != nil {
			stopCH <- true
			fmt.Println(m)
			return
		}
		fmt.Println("READ")
		fmt.Println(m)
		mCh <- m
	}
}
