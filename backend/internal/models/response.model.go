package models

// APIResponse adalah wrapper standar untuk semua response API.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func OK(data interface{}) APIResponse {
	return APIResponse{Success: true, Data: data}
}

func OKMessage(message string) APIResponse {
	return APIResponse{Success: true, Message: message}
}

func Fail(err string) APIResponse {
	return APIResponse{Success: false, Error: err}
}
