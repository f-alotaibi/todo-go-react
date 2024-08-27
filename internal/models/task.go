package models

type TodoTask struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}
