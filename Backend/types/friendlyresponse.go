package types

// FriendlyResponse : Friendly response sent back from backend to front-end for better error handling
type FriendlyResponse struct {
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
	Error   error       `json:"error"`
	Message string      `json:"Message"`
}

// New : Creates a new friendly message object
func NewFriendlyResponse(status int, result interface{}, err error, message string) FriendlyResponse {
	return FriendlyResponse{Status: status, Result: result, Error: err, Message: message}
}
