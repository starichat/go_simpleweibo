package response

// 错误响应体的响应
type Error struct {
	Error string `json:"error"`
}