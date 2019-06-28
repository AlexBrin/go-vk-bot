package object

type Attachment struct {
	Type        string       `json:"type" map:"type"`
	Photo       *Photo       `json:"photo" map:"photo"`
	Video       *Video       `json:"video" map:"video"`
	Audio       *Audio       `json:"audio" map:"audio"`
	Doc         *Doc         `json:"doc" map:"doc"`
	Link        *Link        `json:"link" map:"link"`
	Market      *Market      `json:"market" map:"market"`
	MarketAlbum *MarketAlbum `json:"market_album" map:"market_album"`
	Wall        *Wall        `json:"wall" map:"wall"`
	WallReply   *WallReply   `json:"wall_reply" map:"wall_reply"`
	Sticker     *Sticker     `json:"sticker" map:"sticker"`
	Gift        *Gift        `json:"gift" map:"gift"`
}
