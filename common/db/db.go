package db

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/HelixY2J/echo/common/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	db *sqlx.DB
}

func InitDB(connStr string) (*sqlx.DB, error) {

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to Postgres: %v", err)
		return nil, err
	}
	store := NewStore(db)
	store.createSchema()
	return db, nil

}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s *Store) createSchema() {
	schema := `
	CREATE TABLE IF NOT EXISTS notifications (
	  id SERIAL PRIMARY KEY,
	  recipients JSONB NOT NULL,
	  message TEXT NOT NULL,
	  metadata JSONB NOT NULL,
	  status VARCHAR(50) NOT NULL DEFAULT 'pending',
	  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS roles (
	  id SERIAL PRIMARY KEY,
	  name VARCHAR(50) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS user_roles (
	  user_id VARCHAR(50) NOT NULL,
	  role_id INT NOT NULL,
	  PRIMARY KEY (user_id, role_id)
	);`

	_, err := s.db.Exec(schema)
	if err != nil {
		log.Fatalf(" Failed to create tables: %v", err)
	}

	log.Println(" Database schema is set up!")

}

func (s *Store) InsertNotification(notification models.Notification) error {
	recJSON, err := json.Marshal(notification.Recipients)
	if err != nil {
		return fmt.Errorf("failed to marshal recipients: %w", err)
	}

	metaJSON, err := json.Marshal(notification.Metadata)
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %w", err)
	}

	_, err = s.db.Exec(`
		INSERT INTO notifications (recipients, message, metadata, status)
		VALUES ($1, $2, $3, $4)
	`, recJSON, notification.Message, metaJSON, "pending")
	return err
}
