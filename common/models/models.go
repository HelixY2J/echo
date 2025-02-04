package models

import (
	"time"
)

// type Notification struct {
// 	ID         int             `db:"id"`
// 	Recipients json.RawMessage `db:"recipients"`
// 	Message    string          `db:"message"`
// 	Metadata   json.RawMessage `db:"metadata"`
// 	Status     string          `db:"status"`
// 	CreatedAt  time.Time       `db:"created_at"`
// }

// type Role struct {
// 	ID   int    `db:"id"`
// 	Name string `db:"name"`
// }

// type UserRole struct {
// 	UserID string `db:"user_id"`
// 	RoleID int    `db:"role_id"`
// }

type Notification struct {
	ID         int         `db:"id" json:"ID"`
	Recipients []Recipient `db:"recipients" json:"Recipients"`
	Message    string      `db:"message" json:"Message"`
	Metadata   Metadata    `db:"metadata" json:"Metadata"`
	Status     string      `db:"status" json:"Status"`
	CreatedAt  time.Time   `db:"created_at" json:"CreatedAt"`
}

type Recipient struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Metadata struct {
	Type     string `json:"type"`
	Event    string `json:"event"`
	Employee struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"employee"`
	Timestamp string `json:"timestamp"`
}
