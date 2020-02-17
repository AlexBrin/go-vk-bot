package object

import "fmt"

type Preview struct {
	Photo        *Photo        `json:"photo" map:"photo"`
	Graffiti     *Graffiti     `json:"graffiti" map:"graffiti"`
	AudioMessage *AudioMessage `json:"audio_message" map:"audio_message"`
}

type Doc struct {
	ID      float64  `json:"id" map:"id"`
	OwnerID float64  `json:"owner_id" map:"owner_id"`
	Title   string   `json:"title" map:"title"`
	Size    float64  `json:"size" map:"size"`
	Ext     string   `json:"ext" map:"ext"`
	Url     string   `json:"url" map:"url"`
	Date    float64  `json:"date" map:"date"`
	Type    float64  `json:"type" map:"type"`
	Preview *Preview `json:"preview" map:"preview"`
}

func (d *Doc) BuildAttachment() string {
	return fmt.Sprintf("doc%d_%d", int(d.OwnerID), int(d.ID))
}
