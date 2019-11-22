package defs

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Srcurl string `json:"srcurl"`
}
