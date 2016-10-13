package model

type TodoList []*Todo

type Todo struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
