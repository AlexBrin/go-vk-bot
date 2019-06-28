package object

type AudioMessage struct {
	Duration float64   `json:"duration" map:"duration"`
	Waveform []float64 `json:"waveform" map:"waveform"`
	LinkOGG  string    `json:"link_ogg" map:"link_ogg"`
	LinkMP3  string    `json:"link_mp3" map:"link_mp3"`
}
