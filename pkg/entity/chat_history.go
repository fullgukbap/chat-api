package entity

import "time"

type ChatHistory struct {
	UserID   int
	Content  string
	WritedAt time.Time
}
