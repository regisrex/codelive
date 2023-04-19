package controllers

import (
	"fmt"
	"log"
	"net/http"
	"sharecode/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func MainRoute(c *gin.Context) {
	http.ServeFile(c.Writer , c.Request , "index.html")
}



type Client struct {
	Conn *websocket.Conn
	Id   string
}

var clients []Client

func NewSnippet(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	conn, err := utils.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal("An error occured connecting to the server")
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
	}
	newCli := Client{
		Conn: conn,
		Id:   id,
	}

	clients = append(clients, newCli)

	for {
		mt, m, err := conn.ReadMessage()
		if err != nil {
			return
		}

		for _, client := range clients {
			if client.Id == id {
				client.Conn.WriteMessage(mt, m)
			}
		}
	}

}
