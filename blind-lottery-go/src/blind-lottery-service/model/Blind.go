package model

type Blind struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Score       float32 `json:"score"`
}

var blinds = []Blind{
	{1, "Volmarna", 9.6},
	{1, "Volmarna", 9.6},
	{1, "Volmarna", 9.6},
	{1, "Volmarna", 9.6},
	{1, "Volmarna", 9.6},
	{1, "Volmarna", 9.6},
	{1, "Volmarna", 9.6},
	{1, "Volmarna", 9.6},
	{1, "Volmarna", 9.6},
	{1, "Volmarna", 9.6},
	{1, "Volmarna", 9.6},
	{1, "Volmarna", 9.6},
	{1, "Volmarna", 9.6},
}

func GetBlinds() []Blind {
	return blinds
}
