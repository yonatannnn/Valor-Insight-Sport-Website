package domain

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
