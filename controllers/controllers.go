package controllers

import (
	"log"
	"net/http"
	"sharecode/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func MainRoute(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "index.html")
}

type Client struct {
	Conn *websocket.Conn
	Id   string
	Address string ;
}

var clients []Client

func NewSnippet(c *gin.Context) {
	id  := c.Param("id")
	if strings.TrimSpace(id) == "" {
		log.Fatal("No Id provided")
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
	}
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
		Address: conn.RemoteAddr().String(),
	}

	clients = append(clients, newCli)

	for {
		mt, m, err := conn.ReadMessage()
		if err != nil {
			return
		}
		var destinationclients []Client ;
		
		for _, client := range clients {
			if client.Address != conn.RemoteAddr().String() {
				destinationclients = append(destinationclients, client)
			}
		}
		for _, client := range destinationclients {
			if client.Id == id {
				client.Conn.WriteMessage(mt, m)
			}
		}
	}

}
