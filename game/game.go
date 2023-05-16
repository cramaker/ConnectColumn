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
	WinnerID        sql.NullString
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
		WinnerID:        sql.NullString{},
		Columns:         7,
		Rows:            6,
		Players:         players,
	}, nil
}

// GetGame retrieves an existing ConnectColumn game record
func GetGame(gameID int64, db *sql.DB) (*Game, error) {
	gameData := &Game{}
	err := db.QueryRow("SELECT * FROM games WHERE id = ?", gameID).Scan(
		&gameData.ID,
		&gameData.NumberOfPlayers,
		&gameData.CurrentPlayer,
		&gameData.State,
		&gameData.WinnerID,
		&gameData.Columns,
		&gameData.Rows,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to query game: %w", err)
	}

	rows, err := db.Query("SELECT player_id FROM players WHERE game_id = ?", gameID)
	if err != nil {
		return nil, fmt.Errorf("failed to query players: %w", err)
	}
	defer rows.Close()

	var playerID string
	for rows.Next() {
		err := rows.Scan(&playerID)
		if err != nil {
			return nil, fmt.Errorf("failed to read player row: %w", err)
		}
		gameData.Players = append(gameData.Players, playerID)
	}

	return gameData, nil
}

// ListGames fetches all games from the database that are still in progress
func ListGames(db *sql.DB) ([]*Game, error) {
	rows, err := db.Query("SELECT * FROM games WHERE state = 'IN_PROGRESS'")
	if err != nil {
		return nil, fmt.Errorf("failed to query games: %w", err)
	}
	defer rows.Close()

	games := make([]*Game, 0)
	for rows.Next() {
		game := new(Game)
		err := rows.Scan(&game.ID, &game.NumberOfPlayers, &game.CurrentPlayer, &game.State, &game.WinnerID, &game.Columns, &game.Rows)
		if err != nil {
			return nil, fmt.Errorf("failed to read game row: %w", err)
		}
		games = append(games, game)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to read game row: %w", err)
	}
	return games, nil
}
