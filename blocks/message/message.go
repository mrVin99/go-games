package message

import (
	"errors"
	"github.com/gofiber/contrib/websocket"
	"log"
)

type Message struct {
	Type string
	Data map[string]any
}

func (msg Message) PlayerID() (string, error) {
	playerID, ok := msg.Data["playerID"].(string)
	if !ok {
		return "", errors.New("player id is empty")
	}
	return playerID, nil
}

func (msg Message) BetAmount() (int, error) {
	betAmount, ok := msg.Data["betAmount"].(float64)
	if !ok {
		return 0, errors.New("bet amount is zero")
	}

	return int(betAmount), nil
}

func (msg Message) Risk() (string, error) {
	risk, ok := msg.Data["risk"].(string)
	if !ok {
		return "", errors.New("risk is empty")
	}

	return risk, nil
}

func Send(conn *websocket.Conn, msg Message) {
	if err := conn.WriteJSON(msg); err != nil {
		if conn == nil {
			log.Println("Cannot write to an closed connection")
		} else {
			log.Println("Error writing error message: ", err.Error())
		}
		return
	}
}

func SendError(conn *websocket.Conn, err error) {
	if errJson := conn.WriteJSON(map[string]any{
		"type": "error",
		"data": err.Error(),
	}); errJson != nil {
		if conn == nil {
			log.Println("Cannot write to an closed connection")
		} else {
			log.Println("Error writing error message: ", errJson.Error())
		}
	}
}

func SendStatus(conn *websocket.Conn, status string) {
	if errJson := conn.WriteJSON(map[string]any{
		"type": "status",
		"data": status,
	}); errJson != nil {
		if conn == nil {
			log.Println("Cannot write to an closed connection")
		} else {
			log.Println("Error writing error message: ", errJson.Error())
		}
	}
}
