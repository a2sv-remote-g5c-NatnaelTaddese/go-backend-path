package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func NewTask(id int, title string, description string, completed bool) *Task {
	return &Task{
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}

func (t *Task) Update(title string, description string, completed bool) {
	t.Title = title
	t.Description = description
	t.Completed = completed
}
