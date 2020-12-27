package main

import (
	"github.com/gin-gonic/gin"
	"goroutine/handl"
	"io"
	"log"
	"os"
)

// WebSocket 更新用
func main() {
	logfile := initLog()
	logfile.Close()

	h := handl.NewHandler()

	router := gin.Default()
	router.LoadHTMLGlob("static/*.html")
	router.GET("/room", h.ChatRoom)
	router.GET("/room/:roomid/message", h.Connect)
	router.GET("/user", h.GetUserID)

	router.Run(":80")
}

func initLog() *os.File {
	logfile, err := os.OpenFile("./test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannot open test.log:" + err.Error())
	}

	// io.MultiWriteで、
	// 標準出力とファイルの両方を束ねて、
	// logの出力先に設定する
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))

	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("Log!!")
	return logfile
}
