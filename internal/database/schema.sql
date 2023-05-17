CREATE TABLE IF NOT EXISTS games (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    number_of_players INTEGER NOT NULL,
    current_player INTEGER NOT NULL,
    state TEXT NOT NULL,
    winner_id TEXT,
    columns INTEGER DEFAULT 7,
    rows INTEGER DEFAULT 6,
    board TEXT NOT NULL
                                 );

CREATE TABLE IF NOT EXISTS players (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    game_id INTEGER NOT NULL,
    player_id TEXT NOT NULL,
    FOREIGN KEY (game_id) REFERENCES games (id) ON DELETE CASCADE
                                   );

CREATE TABLE IF NOT EXISTS moves (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    game_id INTEGER NOT NULL,
    player_id TEXT NOT NULL,
    move_type TEXT NOT NULL,
    "column" INTEGER,
    FOREIGN KEY (game_id) REFERENCES games (id) ON DELETE CASCADE
                                 );
