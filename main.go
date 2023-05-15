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

	// Create a new gameInstance
	players := []string{"player1", "player2", "player3"}
	gameInstance, err := game.CreateGame(players, db)
	if err != nil {
		log.Fatalf("Failed to create gameInstance: %v", err)
	}
	fmt.Printf("Created gameInstance with ID: %d\n", gameInstance.ID)

	// Get gameInstance data
	gameData, err := game.GetGame(gameInstance.ID, db)
	if err != nil {
		log.Fatalf("Failed to get gameInstance: %v", err)
	}

	fmt.Printf("Game ID: %d\n", gameData.ID)
	fmt.Printf("Number of players: %d\n", gameData.NumberOfPlayers)
	fmt.Printf("Current player: %d\n", gameData.CurrentPlayer)
	fmt.Printf("Game state: %s\n", gameData.State)
	if gameData.WinnerID.Valid {
		fmt.Printf("Winner ID: %s\n", gameData.WinnerID.String)
	} else {
		fmt.Println("Winner ID: No Current Winner")
	}
	fmt.Printf("Columns: %d\n", gameData.Columns)
	fmt.Printf("Rows: %d\n", gameData.Rows)
	fmt.Printf("Players: %v\n", gameData.Players)
}
