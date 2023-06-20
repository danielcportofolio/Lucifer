package models

import "time"

type Chat struct {
	ID       int        `json:"id"`
	UserID1  int        `json:"user_id_1"`
	UserID2  int        `json:"user_id_2"`
	Messages []Message  `json:"messages"`
}

type Message struct {
	ID          int         `json:"id"`
	UserID      int         `json:"user_id"`
	Message     string      `json:"message"`
	Attachments []Attachment `json:"attachments"`
	CreatedAt   time.Time   `json:"created_at"`
}

type Attachment struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}
