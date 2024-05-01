package blocks

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-games/blocks/dice"
	"go-games/blocks/message"
	"go-games/blocks/player"
	"go-games/pkg/cache"
	"go-games/pkg/database"
	"log"
	"math/rand/v2"
)

func connectionID() int32 {
	return rand.Int32()
}

func Play(memo cache.Cache, db database.DB) fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
		connID := connectionID()
		games := make(map[string]*Game) // Map to store games for the connection

		defer closeConnection(conn, connID, db, memo, games)

		const maxErrors = 3
		var errorCount int

		errorCounter := func() bool {
			errorCount++
			return errorCount > maxErrors
		}

		for {
			var msg message.Message

			if err := conn.ReadJSON(&msg); err != nil {
				message.SendError(conn, err)
				if errorCounter() {
					return
				}
				continue
			}

			log.Printf("msg: %v", msg)

			if msg.Type == "connect" {
				playerID, err := msg.PlayerID()
				if err != nil {
					message.SendError(conn, err)
					if errorCounter() {
						return
					}
					continue
				}

				if err = player.ExistsById(db, playerID); err != nil {
					message.SendError(conn, err)
					if errorCounter() {
						return
					}
					continue
				}

				message.SendStatus(conn, "ok")
			}

			if msg.Type == "bet" {
				playerID, err := msg.PlayerID()
				if err != nil {
					message.SendError(conn, err)
					if errorCounter() {
						return
					}
					continue
				}

				amount, err := msg.BetAmount()
				if err != nil {
					message.SendError(conn, err)
					if errorCounter() {
						return
					}
					continue
				}

				risk, err := msg.Risk()
				if err != nil {
					message.SendError(conn, err)
					if errorCounter() {
						return
					}
					continue
				}

				params := gameParams{
					gameID:    uuid.New().String(),
					playerID:  playerID,
					betAmount: amount,
					risk:      risk,
				}

				game, err := betAction(params, memo, db)
				if err != nil {
					message.SendError(conn, err)
					if errorCounter() {
						return
					}
					continue
				}

				// Store the game in the map under the connection ID
				games[params.gameID] = game

				log.Println("Game: ", game)

				if err = conn.WriteJSON(map[string]interface{}{
					"type": "game_data",
					"data": game,
				}); err != nil {
					message.SendError(conn, err)
					if errorCounter() {
						return
					}
					continue
				}
			}
		}
	})
}

func closeConnection(conn *websocket.Conn, connID int32, db database.DB, memo cache.Cache, games map[string]*Game) {
	log.Println("Closing connection")

	for gameID, game := range games {
		query := `
            INSERT INTO player_history (player_id, game_id, risk, bet_amount, win_amount)
            VALUES ($1, $2, $3, $4, $5)
        `

		if err := db.
			Exec(query, game.PlayerID, gameID, game.Risk, game.BetAmount, game.Result); err != nil {
			log.Println("Error inserting data into player_history table:", err)
			continue
		}
	}

	if err := conn.Close(); err != nil {
		log.Println("Error closing WebSocket connection:", err)
		return
	}
}

func betAction(params gameParams, memo cache.Cache, db database.DB) (*Game, error) {
	if err := player.SufficientFunds(db, params.playerID, params.betAmount); err != nil {
		return nil, err
	}

	d := dice.NewDicer()

	game := &Game{
		ID:        params.gameID,
		PlayerID:  params.playerID,
		BetAmount: params.betAmount,
		Dice:      d,
		Risk:      params.risk,
	}

	game.Dice.Roll()

	premiumMultiplier := dice.CalculatePremiumMultiplier(game.Risk, game.Dice.Values())
	game.Result = premiumMultiplier * game.BetAmount

	if err := memo.Set(game.ID, game); err != nil {
		return nil, err
	}

	return game, nil
}
