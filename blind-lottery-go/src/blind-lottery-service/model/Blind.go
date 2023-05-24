package model

type Blind struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Score       float32 `json:"score"`
}

var blinds = []Blind{
	{"1", "Volmarna", 9.6},
	{"2", "Volmarna", 9.6},
	{"3", "Volmarna", 9.6},
	{"4", "Volmarna", 9.6},
	{"5", "Volmarna", 9.6},
	{"6", "Volmarna", 9.6},
	{"7", "Volmarna", 9.6},
	{"8", "Volmarna", 9.6},
	{"9", "Volmarna", 9.6},
	{"10", "Volmarna", 9.6},
	{"11", "Volmarna", 9.6},
	{"12", "Volmarna", 9.6},
	{"13", "Volmarna", 9.6},
	{"14", "Volmarna", 9.6},
	{"15", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
	{"16", "Volmarna", 9.6},
}

func GetBlinds() []Blind {
	return blinds
}
