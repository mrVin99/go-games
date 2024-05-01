package blocks

import (
	"go-games/blocks/dice"
)

type Game struct {
	ID        string
	PlayerID  string
	Risk      string
	BetAmount int
	Dice      *dice.Dicer
	Result    int
}

type gameParams struct {
	gameID    string
	playerID  string
	betAmount int
	risk      string
}
