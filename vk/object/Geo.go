package object

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Place struct {
	ID        float64 `json:"id"`
	Title     string  `json:"title"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	CreatedAt float64 `json:"created"`
	Icon      string  `json:"icon"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
}

type Geo struct {
	Type        string       `json:"type"`
	Coordinates *Coordinates `json:"coordinates"`
	Place       *Place       `json:"place"`
}
