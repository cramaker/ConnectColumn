package game

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("failed to open in-memory SQLite database: %v", err)
	}

	err = createTables(db, "../internal/database/schema.sql")
	if err != nil {
		t.Fatalf("failed to create tables: %v", err)
	}

	return db
}

func TestCreateAndGetGame(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Test creating a new game
	players := []string{"player1", "player2", "player3"}
	game, err := CreateGame(players, db)
	if err != nil {
		t.Fatalf("failed to create game: %v", err)
	}

	// Test getting the created game
	retrievedGame, err := GetGame(game.ID, db)
	if err != nil {
		t.Fatalf("failed to get game: %v", err)
	}

	// Check if the game data is correct
	if len(retrievedGame.Players) != len(players) {
		t.Errorf("expected %d players, got %d", len(players), len(retrievedGame.Players))
	}

	for i, player := range retrievedGame.Players {
		if player != players[i] {
			t.Errorf("expected player %s, got %s", players[i], player)
		}
	}

	if retrievedGame.State != "IN_PROGRESS" {
		t.Errorf("expected state to be 'IN_PROGRESS', got %s", retrievedGame.State)
	}
}

func createTables(db *sql.DB, schemaFile string) error {
	schema, err := os.ReadFile(schemaFile)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(schema))
	return err
}
