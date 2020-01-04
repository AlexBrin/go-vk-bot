package object

type Price struct {
	Amount float64 `json:"amount" map:"amount"`
}

type Market struct {
	ID          float64 `json:"id" map:"id"`
	OwnerID     float64 `json:"owner_id" map:"owner_id"`
	Title       string  `json:"title" map:"title"`
	Description string  `json:"description" map:"description"`
	//	Price *Price        `json:"price" map:"price"`
}
