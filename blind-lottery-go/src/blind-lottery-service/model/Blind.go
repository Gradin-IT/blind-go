package model

type Blind struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Score       float32 `json:"score"`
}

var blinds = []Blind{
	{1, "Volmarna", 9.6},
	{2, "Volmarna", 9.6},
	{3, "Volmarna", 9.6},
	{4, "Volmarna", 9.6},
	{5, "Volmarna", 9.6},
	{6, "Volmarna", 9.6},
	{7, "Volmarna", 9.6},
	{8, "Volmarna", 9.6},
	{9, "Volmarna", 9.6},
	{10, "Volmarna", 9.6},
	{11, "Volmarna", 9.6},
	{12, "Volmarna", 9.6},
	{13, "Volmarna", 9.6},
	{14, "Volmarna", 9.6},
	{15, "Volmarna", 9.6},
	{16, "Volmarna", 9.6},
	{17, "Volmarna", 9.6},
	{18, "Volmarna", 9.6},
	{19, "Volmarna", 9.6},
	{20, "Volmarna", 9.6},
	{21, "Volmarna", 9.6},
	{22, "Volmarna", 9.6},
	{23, "Volmarna", 9.6},
	{24, "Volmarna", 9.6},
	{25, "Volmarna", 9.6},
	{26, "Volmarna", 9.6},
	{27, "Volmarna", 9.6},
	{28, "Volmarna", 9.6},
	{29, "Volmarna", 9.6},
	{30, "Volmarna", 9.6},
	{31, "Volmarna", 9.6},
	{32, "Volmarna", 9.6},
	{33, "Volmarna", 9.6},
	{34, "Volmarna", 9.6},
	{35, "Volmarna", 9.6},
	{36, "Volmarna", 9.6},
	{37, "Volmarna", 9.6},
	{38, "Volmarna", 9.6},
	{39, "Volmarna", 9.6},
}

func GetBlinds() []Blind {
	return blinds
}
