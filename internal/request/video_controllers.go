package request

type VideoControllers struct {
	Pause    bool   `json:"pause"`
	Timeline string `json:"timeline"`
	Movie    string `json:"movie"`
}