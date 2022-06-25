package vimeworld

// Error is a api error.
type Error struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	Comment   string `json:"comment"`
}

func (e Error) Error() string {
	return "go-vimeworld: " + e.ErrorMsg
}
