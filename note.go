package main

import "time"

type Note struct {
	ID        int
	CreatedAt time.Time
	IsDone    bool
	Text      string
}
