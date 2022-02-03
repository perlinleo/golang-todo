package model

type Task struct {
	ID             int32    `json:"id,omitempty"`
	Name  		   string   `json:"name,omitempty"`
	Done 		   bool  	`json:"done"`
	Description    string 	`json:"about"`
}