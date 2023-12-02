package main

type terminal struct {
	terminalID int64  `json:"id"`
	serverID   int64  `json:"id"`
	Lokasi     string `json:"location"`
	status     string `json:"status"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type user struct {
	userID   int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type server struct {
	serverID    int64  `json:"id"`
	status      string `json:"status"`
	lokasi      string `json:"lokasi"`
	server_type string `json:"serverType"`
}

type payment struct {
	paymentID     int64  `json:"id"`
	transactionID string `json:"id"`
	status        string `json:"status"`
	amount        string `json:"amount"`
}
