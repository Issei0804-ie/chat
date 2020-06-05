package domain

type ReadMessage struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type SendMessage struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

type User struct {
	UserID int `json:"userid"`
}
