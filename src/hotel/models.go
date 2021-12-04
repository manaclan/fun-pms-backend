package hotel

import (
	"html"
	"strings"
)

type Location struct {
	Address        string
	Street         string
	Ward           string
	District       string
	CityOrProvince string
	Country        string
}

type Room struct {
	LayoutFile    string
	PricePerNight int32
	Type          int32 //1 is single, 2 is double, 3 is family
}

type Floor struct {
	LayoutFile  string
	FloorNumber string
	Rooms       []Room
}

type Hotel struct {
	Name     string
	Location Location
	Floors   []Floor
}

func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}
