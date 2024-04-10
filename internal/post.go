package internal

type TaskModel struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}
