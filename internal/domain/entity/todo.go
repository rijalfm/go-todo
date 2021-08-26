package entity

import "time"

// Todo Structure
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
}

// Validate Not Null Todo Title
func (todo *Todo) Validate() map[string]string {
	var err = make(map[string]string)

	if todo.Title == "" || todo.Title == "null" {
		err["message"] = "title is required"
		return err
	}

	return nil

}
