package object

const (
	MESSAGEACTION_CHAT_PHOTO_UPDATE        = "chat_photo_update"
	MESSAGEACTION_CHAT_PHOTO_REMOVE        = "chat_photo_REMOVE"
	MESSAGEACTION_CHAT_TITLE_UPDATE        = "chat_title_update"
	MESSAGEACTION_CHAT_INVITE_USER         = "chat_invite_user"
	MESSAGEACTION_CHAT_INVITE_USER_BY_LINK = "chat_invite_user_by_link"
	MESSAGEACTION_CHAT_KICK_USER           = "chat_kick_user"
	MESSAGEACTION_CHAT_PIN_MESSAGE         = "chat_pin_message"
	MESSAGEACTION_CHAT_UNPIN_MESSAGE       = "chat_unpin_message"
)

type MessageActionPhoto struct {
	Photo50  string `json:"photo_50"`
	Photo100 string `json:"photo_100"`
	Photo200 string `json:"photo_200"`
}

type MessageAction struct {
	Type     string             `json:"type"`
	MemberID float64            `json:"member_id"`
	Text     string             `json:"text"`
	Email    string             `json:"email"`
	Photo    MessageActionPhoto `json:"photo"`
}

type PrivateMessage struct {
	ID                float64           `json:"id" map:"id"`
	Date              float64           `json:"date" map:"date"`
	PeerID            float64           `json:"peer_id" map:"peer_id"`
	FromID            float64           `json:"from_id" map:"from_id"`
	UserID            float64           `json:"from_id" map:"from_id"`
	Text              string            `json:"text" map:"text"`
	RandomID          float64           `json:"random_id" map:"random_id"`
	Attachments       []*Attachment     `json:"attachments" map:"attachments"`
	Important         bool              `json:"important" map:"important"`
	Geo               *Geo              `json:"geo" map:"geo"`
	Payload           string            `json:"payload" map:"payload"`
	ForwardedMessages []*PrivateMessage `json:"fwd_messages" map:"fwd_messages"`
	ReplyMessage      *PrivateMessage   `json:"reply_message" map:"reply_message"`
	Action            *MessageAction    `json:"action" map:"action"`
}
