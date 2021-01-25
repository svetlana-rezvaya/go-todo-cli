package main

import "time"

type note struct {
	ID        int
	CreatedAt time.Time
	IsDone    bool
	Text      string
}
