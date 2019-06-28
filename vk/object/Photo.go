package object

type PhotoSize struct {
	Type   string  `json:"type" map:"type"`
	Url    string  `json:"url" map:"url"`
	Width  float64 `json:"width" map:"width"`
	Height float64 `json:"height" map:"height"`
}

type Photo struct {
	ID      float64     `json:"id" map:"id"`
	AlbumID float64     `json:"album_id" map:"album_id"`
	OwnerID float64     `json:"owner_id" map:"owner_id"`
	Text    string      `json:"text" map:"text"`
	Date    float64     `json:"date" map:"date"`
	Size    []PhotoSize `json:"sizes" map:"sizes"`
	Width   float64     `json:"width" map:"width"`
	Height  float64     `json:"height" map:"height"`

	biggerImageUrl string
}

func (p *Photo) GetBiggerImageUrl() (currentUrl string) {
	if p.biggerImageUrl != "" {
		return p.biggerImageUrl
	}

	var maxSize float64 = 0
	for _, size := range p.Size {
		crntSize := size.Width * size.Height
		if crntSize > maxSize {
			currentUrl = size.Url
			maxSize = crntSize
		}
	}

	p.biggerImageUrl = currentUrl
	return
}
