package healthcheck

// Healthcheck struct
type Healthcheck struct {
	Status  string `json:"status,omitempty" example:"ok"`
	Message string `json:"message,omitempty" example:"Server is accepting connections"`
}
