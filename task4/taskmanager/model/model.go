package model
import("time"
)
type Task struct {
	ID          string   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate	    time.Time `json:"due_date"`
	}	