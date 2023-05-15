package main

import (
	"database/sql"
	"fmt"
	"github.com/cramaker/ConnectColumn/game"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "connectcolumn.db")
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}
	defer db.Close()

	// Create a new game with two players
	newGame, err := game.CreateGame([]string{"player1", "player2"}, db)
	if err != nil {
		log.Fatalf("Failed to create game: %v", err)
	}

	// Print the game information
	fmt.Printf("Created game: %+v\n", newGame)
}
