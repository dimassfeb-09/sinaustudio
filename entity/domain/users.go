package domain

type Users struct {
	UUID    string `json:uuid`
	Name    string `json:email`
	Email   string `json:email`
	ClassID int    `json:class_id`
}
