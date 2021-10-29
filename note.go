package todo

import "time"

// Note ...
type Note struct {
	ID        int
	CreatedAt time.Time
	IsDone    bool
	Text      string
}
