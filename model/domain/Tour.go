package domain

type Tour struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Duration    string `json:"duration"`
	Description string `json:"description"`
}
