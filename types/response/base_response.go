package response

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Payload any    `json:"payload,omitempty"`
}
