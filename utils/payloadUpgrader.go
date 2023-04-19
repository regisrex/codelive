package utils

import (
	"os"
	"strconv"

	"github.com/gorilla/websocket"
)

var Upgrader websocket.Upgrader ;

func PayloadUpgrader(){

	bufferSize, _ := strconv.Atoi(os.Getenv("BUFFER_SIZE"))

	 Upgrader = websocket.Upgrader{
		ReadBufferSize:  bufferSize,
		WriteBufferSize: bufferSize,
	}

}
