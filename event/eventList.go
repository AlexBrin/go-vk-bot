package event

const (
	CommandEvent = "message_command"
	PayloadEvent = "message_payload"

	// messages
	// input message
	MessageNewEvent = "message_new"
	// output message
	MessageReplyEvent = "message_reply"
	// edit message
	MessageEditEvent = "message_edit"
	// subscribe on messages
	MessageAllowEvent = "message_allow"
	// unsubscribe from messages
	MessageDenyEvent = "message_deny"

	// photos
	//EventPhotoNew            = "photo_new"
	//EventPhotoCommentNew     = "photo_comment_new"
	//EventPhotoCommentEdit    = "photo_comment_edit"
	//EventPhotoCommentDelete  = "photo_comment_delete"
	//EventPhotoCommentRestore = "photo_comment_restore"

	// audios
	//EventAudioNew = "audio_new"

	// videos
	//EventVideoNew            = "video_new"
	//EventVideoCommentNew     = "video_comment_new"
	//EventVideoCommentEdit    = "video_comment_edit"
	//EventVideoCommentDelete  = "video_comment_delete"
	//EventVideoCommentRestore = "video_comment_restore"

	// walls
	//EventWallNew    = "wall_post_new"
	//EventWallRepost = "wall_repost"

	// wall replies
	//EventWallReplyNew     = "wall_reply_new"
	//EventWallReplyEdit    = "wall_reply_edit"
	//EventWallReplyDelete  = "wall_reply_delete"
	//EventWallReplyRestore = "wall_reply_restore"
)
