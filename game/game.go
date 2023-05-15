package game

import (
	"database/sql"
	"fmt"
)

// Game represents a ConnectColumn game
type Game struct {
	ID              int64
	NumberOfPlayers int
	CurrentPlayer   int
	State           string
	WinnerID        string
	Columns         int
	Rows            int
	Players         []string
}

// CreateGame creates a new ConnectColumn game with two or more players
func CreateGame(players []string, db *sql.DB) (*Game, error) {
	state := "IN_PROGRESS"
	numberOfPlayers := len(players)
	currentPlayer := 1

	result, err := db.Exec(`INSERT INTO games (number_of_players, current_player, state, winner_id, columns, rows) VALUES (?, ?, ?, NULL, 7, 6)`, numberOfPlayers, currentPlayer, state)
	if err != nil {
		return nil, fmt.Errorf("failed to insert game: %w", err)
	}

	gameID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get the last inserted game ID: %w", err)
	}

	for _, playerID := range players {
		_, err := db.Exec(`INSERT INTO players (game_id, player_id) VALUES (?, ?)`, gameID, playerID)
		if err != nil {
			return nil, fmt.Errorf("failed to insert player: %w", err)
		}
	}

	return &Game{
		ID:              gameID,
		NumberOfPlayers: numberOfPlayers,
		CurrentPlayer:   currentPlayer,
		State:           state,
		WinnerID:        "",
		Columns:         7,
		Rows:            6,
		Players:         players,
	}, nil
}
