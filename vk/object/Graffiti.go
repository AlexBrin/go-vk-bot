package object

type Graffiti struct {
	Source string  `json:"src" map:"src"`
	Width  float64 `json:"width" map:"width"`
	Height float64 `json:"height" map:"height"`
}
