package model

type Area struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	BlindIds []int  `json:"blinds"`
}

var areas = []Area{
	{1, "Brattåsberget", []int{1, 2, 3}},
	{2, "Högåsen", []int{4, 2, 5}},
	{3, "Gammelkullen", []int{7, 8, 9}},
	{4, "Rödmyrberget", []int{8, 10, 11}},
}

func GetAreas() []Area {
	return areas
}
