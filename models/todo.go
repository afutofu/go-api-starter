package models

/*
Todo represents a task that needs to be done.

Fields:
  - ID int: The unique identifier of the todo item.
  - Text string: The text or description of the todo item.
  - Completed boolean: Indicates whether the todo item is completed.
*/
type Todo struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}
