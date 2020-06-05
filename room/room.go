package room

import (
	"fmt"
	"github.com/gorilla/websocket"
	"goroutine/client"
	"goroutine/domain"
)

const (
	MaxConnection = 100
)

type Room struct {
	counter int                     //クライアントのコネクションを数える
	mCH     chan domain.ReadMessage //メッセージを送受信するためのチャンネル
	clientCH chan client.Client
	status  bool
}

func NewRoom() Room {
	mCH := make(chan domain.ReadMessage, 10)
	clientCH := make(chan client.Client, 10)
	r := Room{counter: 0, mCH: mCH, clientCH:clientCH, status: true}
	go r.Write()
	return r
}

func (r *Room) GetStatus() bool {
	return r.status
}

func (r *Room) Write() {
	var clients []client.Client
	//誰かがReadするまで待機
	for {
		select{
		case m := <-r.mCH:
			fmt.Println("Write")

			for i := 0; i < len(clients); i++ {
				if clients[i].GetStatus() == true {
					fmt.Println("WriteNow")
					if err := clients[i].GetConn().WriteJSON(&m); err != nil {
					}
				}
			}

		case c := <- r.clientCH:
			clients = append(clients, c)

		default:
		}
	}
}

//各クライアントのReadに並列処理を実行させる
func (r *Room) Read(conn *websocket.Conn) {
	stopCH := make(chan bool)
	client := client.NewClient(conn)
	fmt.Println("READ INFO:")
	fmt.Println(r.counter)
	go client.Read(r.mCH, stopCH)
	r.counter++
	r.clientCH <- client
	
	for {
		select {
		case <-stopCH:
			fmt.Println("FinishREAD")
			return
		default:
		}
	}


}
