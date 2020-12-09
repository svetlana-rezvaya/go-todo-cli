package main

import "time"

type note struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDone    bool
	Text      string
}
