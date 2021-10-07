package models

type Todo struct {
	ID uint `grom:"primarykey" json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}