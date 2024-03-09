package model

type Task struct {
	ID          int
	Title       string
	Description string
	Status      int //0 : To do, 1 : In progress, 2 : Done
}
