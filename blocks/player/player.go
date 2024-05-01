package player

import (
	"errors"
	"games/pkg/database"
)

type Player struct {
	ID    string
	Funds int
}

func ExistsById(db database.DB, playerID string) error {
	var exists bool
	if err := db.
		QueryRow("SELECT EXISTS (SELECT 1 FROM player WHERE id = $1)", playerID).
		Scan(&exists); err != nil {
		return err
	}

	if !exists {
		return errors.New("player does not exist")
	}

	return nil
}

func SufficientFunds(db database.DB, playerID string, betAmount int) error {
	var balance int
	if err := db.
		QueryRow("SELECT balance FROM player where id = $1", playerID).
		Scan(&balance); err != nil {
		return err
	}

	if betAmount > balance {
		return errors.New("insufficient funds")
	}

	return nil
}
