package workers

// WorkerMessage represents the message structure for the workers
type WorkerMessage struct {
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
	ID      string      `json:"id"`
}

// UserPayload represents the user data in the message
type UserPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
