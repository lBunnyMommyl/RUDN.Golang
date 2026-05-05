package models

type Service struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Price       string `json:"price"`
	Duration    string `json:"duration"`
	Description string `json:"description"`
}